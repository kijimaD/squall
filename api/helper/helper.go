package helper

import (
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/getkin/kin-openapi/routers"
	"github.com/gin-gonic/gin"
	gosqlite "github.com/glebarez/go-sqlite"
	"gorm.io/gorm"
	sqlite3 "modernc.org/sqlite/lib"
)

func GetPtr[T any](x T) *T {
	return &x
}

func OKResponse(c *gin.Context) {
	c.PureJSON(http.StatusOK, StandardResponse{Code: http.StatusOK, Message: "OK"})
}

var (
	errOpenapiRequest  *openapi3filter.RequestError
	errOpenapiSecurity *openapi3filter.SecurityRequirementsError
	errOpenapiRoute    *routers.RouteError
	errUnmarshallJSON  *json.UnmarshalTypeError
	errSQLite          *gosqlite.Error
)

var (
	// ErrRecordNotFound はアプリケーションのロジックによるnot found.
	ErrRecordNotFound = errors.New("関連レコードがなかった")
)

type BadRequestResponse struct {
	Code    int      `json:"code"`
	Message string   `json:"message"`
	Errors  []string `json:"errors"`
}

type StandardResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func ErrorResponse(c *gin.Context, err error) {
	var status int
	var message string
	var errs []string

	fallback := func() {
		slog.Error(err.Error())
		status = http.StatusInternalServerError
		message = "internal server error"
	}

	switch {
	case err.Error() == "EOF":
		status = http.StatusBadRequest
		message = "bad request"
		errs = []string{err.Error()}
	case errors.As(err, &errOpenapiRequest), errors.Is(err, openapi3filter.ErrInvalidEmptyValue), errors.Is(err, openapi3filter.ErrInvalidRequired):
		status = http.StatusBadRequest
		message = "bad request / oapi"
		errs = []string{err.Error()}
	case errors.As(err, &errUnmarshallJSON):
		status = http.StatusBadRequest
		message = "bad request / invalid JSON format"
		errs = []string{err.Error()}
	case errors.As(err, &errSQLite):
		switch errSQLite.Code() {
		case sqlite3.SQLITE_CONSTRAINT_CHECK:
			status = http.StatusBadRequest
			message = "bad request / check constraint"
			errs = []string{err.Error()}
		case sqlite3.SQLITE_CONSTRAINT_UNIQUE:
			status = http.StatusBadRequest
			message = "bad request / unique constraint"
			errs = []string{err.Error()}
		default:
			fallback()
		}
	case errors.As(err, &errOpenapiSecurity):
		status = http.StatusForbidden
		message = "forbidden / oapi"
	case errors.As(err, &errOpenapiRoute):
		status = http.StatusNotFound
		message = fmt.Sprintf("not found / %s", err.Error())
	case errors.Is(err, gorm.ErrRecordNotFound):
		status = http.StatusNotFound
		message = "not found / by ORM"
	case errors.Is(err, ErrRecordNotFound):
		status = http.StatusNotFound
		message = "not found / by logic"
	default:
		fallback()
	}

	if status == http.StatusBadRequest {
		c.PureJSON(status, BadRequestResponse{Code: status, Message: message, Errors: errs})
	} else {
		c.PureJSON(status, StandardResponse{Code: status, Message: message})
	}
}
