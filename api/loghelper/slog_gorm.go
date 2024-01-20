package loghelper

import (
	"context"
	"errors"
	"log/slog"
	"time"

	gormLogger "gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
)

type GormSlogLogger struct {
	LogLevel gormLogger.LogLevel
}

//nolint:ireturn
func (l *GormSlogLogger) LogMode(level gormLogger.LogLevel) gormLogger.Interface {
	newLogger := *l
	newLogger.LogLevel = level

	return &newLogger
}

func (l GormSlogLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= gormLogger.Info {
		slog.LogAttrs(ctx, slog.LevelInfo, msg, slog.Any("data", data))
	}
}

func (l GormSlogLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= gormLogger.Warn {
		slog.LogAttrs(ctx, slog.LevelWarn, msg, slog.Any("data", data))
	}
}

func (l GormSlogLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= gormLogger.Error {
		slog.LogAttrs(ctx, slog.LevelError, msg, slog.Any("data", data))
	}
}

func (l GormSlogLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	if l.LogLevel <= gormLogger.Silent {
		return
	}

	elapsed := time.Since(begin)
	switch {
	case err != nil && l.LogLevel >= gormLogger.Error && !errors.Is(err, gormLogger.ErrRecordNotFound):
		sql, rows := fc()
		slog.LogAttrs(
			ctx,
			slog.LevelError,
			"GORM-Trace",
			slog.String("logtype", "sqllog"),
			slog.String("source", utils.FileWithLineNum()),
			slog.Any("error", err),
			slog.Duration("elapsed", elapsed),
			slog.Int64("rows", rows),
			slog.String("sql", sql),
		)
	case l.LogLevel == gormLogger.Info:
		sql, rows := fc()
		slog.LogAttrs(
			ctx,
			slog.LevelInfo,
			"GORM-Trace",
			slog.String("logtype", "sqllog"),
			slog.String("source", utils.FileWithLineNum()),
			slog.Duration("elapsed", elapsed),
			slog.Int64("rows", rows),
			slog.String("sql", sql),
		)
	}
}
