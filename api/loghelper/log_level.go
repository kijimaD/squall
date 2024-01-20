package loghelper

import (
	"log/slog"

	gormLogger "gorm.io/gorm/logger"
)

type LogLevel string

var (
	LogLevelError LogLevel = "error"
	LogLevelWarn  LogLevel = "warn"
	LogLevelInfo  LogLevel = "info"
	LogLevelDebug LogLevel = "debug"
)

func GetSlogLogLevel(logLevel LogLevel) slog.Level {
	switch logLevel {
	case LogLevelError:
		return slog.LevelError
	case LogLevelWarn:
		return slog.LevelWarn
	case LogLevelInfo:
		return slog.LevelInfo
	case LogLevelDebug:
		return slog.LevelDebug
	default:
		return slog.LevelInfo
	}
}

func GetGormLogLevel(logLevel LogLevel) gormLogger.LogLevel {
	switch logLevel {
	case LogLevelError:
		return gormLogger.Error
	case LogLevelWarn:
		return gormLogger.Warn
	case LogLevelInfo:
		return gormLogger.Info
	default:
		return gormLogger.Info
	}
}
