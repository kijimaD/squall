package routers

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
	"squall/config"
	"squall/generated"
	"squall/helper"
	"squall/loghelper"
	"time"

	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/getkin/kin-openapi/routers/gorillamux"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	sloggin "github.com/samber/slog-gin"
)

func NewRouter() (*gin.Engine, error) {
	gin.DefaultWriter = io.Discard

	r := gin.Default()
	r.Use(CORS())
	if validator, err := MakeValidateMiddleware(); err == nil {
		r.Use(validator)
	} else {
		return nil, err
	}
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

	generated.RegisterHandlersWithOptions(
		r,
		&BaseHandler{},
		generated.GinServerOptions{
			ErrorHandler: func(c *gin.Context, err error, statusCode int) {
				c.JSON(statusCode, gin.H{"message": err.Error()})
			},
		},
	)

	return r, nil
}

func CORS() gin.HandlerFunc {
	config := cors.Config{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders: []string{"Content-Type", "Authorization"},
		MaxAge:       12 * time.Hour,
	}
	return cors.New(config)
}

func MakeValidateMiddleware() (gin.HandlerFunc, error) {
	doc, err := generated.GetSwagger()
	if err != nil {
		return nil, fmt.Errorf("OpenAPIドキュメントの読み込みに失敗した: %w", err)
	}
	doc.Servers = nil

	if err := doc.Validate(context.Background()); err != nil {
		return nil, fmt.Errorf("OpenAPIドキュメントが適切でない: %w", err)
	}

	router, err := gorillamux.NewRouter(doc)
	if err != nil {
		return nil, fmt.Errorf("OpenAPIドキュメントルータの作成に失敗した: %w", err)
	}

	return func(c *gin.Context) {
		route, pathParams, err := router.FindRoute(c.Request)
		if err != nil {
			helper.ErrorResponse(c, err)
			c.Abort()

			return
		}

		requestValidationInput := &openapi3filter.RequestValidationInput{
			Request:    c.Request,
			PathParams: pathParams,
			Route:      route,
			Options: &openapi3filter.Options{
				MultiError:          true,
				SkipSettingDefaults: true,
			},
		}
		// リクエストがOpenAPIに準拠しているかを確認して、準拠してないと処理を中断する
		if err = openapi3filter.ValidateRequest(c.Request.Context(), requestValidationInput); err != nil {
			helper.ErrorResponse(c, err)

			return
		}

		w := &responseBodyWriter{body: &bytes.Buffer{}, ResponseWriter: c.Writer}
		c.Writer = w
		c.Next()

		// レスポンスがOpenAPIに準拠しているかを確認して、準拠してないとログに警告を出す
		if err = openapi3filter.ValidateResponse(c.Request.Context(), &openapi3filter.ResponseValidationInput{
			RequestValidationInput: requestValidationInput,
			Status:                 c.Writer.Status(),
			Header:                 c.Writer.Header(),
			Body:                   io.NopCloser(w.body),
			Options:                &openapi3filter.Options{},
		}); err != nil {
			fmt.Println(err)
			c.Abort()

			return
		}
	}, nil
}

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)

	i, err := r.ResponseWriter.Write(b)
	if err != nil {
		return 0, fmt.Errorf("書き込みに失敗した: %w", err)
	}

	return i, nil
}
