package routers

import (
	"io"
	"log/slog"
	"os"
	"squall/config"
	"squall/loghelper"

	"github.com/gin-gonic/gin"
	sloggin "github.com/samber/slog-gin"
)

func NewRouter() (*gin.Engine, error) {
	gin.DefaultWriter = io.Discard

	r := gin.Default()
	l := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: loghelper.GetSlogLogLevel(config.Config.LogLevel),
	}))
	slog.SetDefault(l)
	config := sloggin.Config{
		DefaultLevel:       slog.LevelInfo,
		ClientErrorLevel:   slog.LevelWarn,
		ServerErrorLevel:   slog.LevelError,
		WithUserAgent:      true,
		WithRequestID:      true,
		WithRequestBody:    true,
		WithRequestHeader:  false,
		WithResponseBody:   true,
		WithResponseHeader: false,
		WithSpanID:         false,
		WithTraceID:        false,
	}
	l = l.With("logtype", "resplog")
	r.Use(sloggin.NewWithConfig(l, config))

	r.GET("/", Roots)
	r.GET("/entries", GetEntries)

	return r, nil
}
