package testhelper

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func ConvertToBody(t *testing.T, data interface{}) *strings.Reader {
	t.Helper()
	b, err := json.Marshal(data)
	assert.NoError(t, err)

	return strings.NewReader(string(b))
}

func MakeRequest(t *testing.T, method string, path string, body interface{}) (*http.Request, *httptest.ResponseRecorder) {
	t.Helper()
	headers := map[string]string{}

	var br io.Reader
	if body != nil {
		switch b := body.(type) {
		case string:
			br = strings.NewReader(b)
		default:
			br = ConvertToBody(t, b)
		}
		headers["Content-Type"] = gin.MIMEJSON
	}
	req := httptest.NewRequest(method, path, br)
	for key, value := range headers {
		req.Header.Set(key, value)
	}
	rec := httptest.NewRecorder()

	return req, rec
}
