package routers

import (
	"encoding/json"
	"net/http"
	"squall/config"
	"squall/testhelper"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoot(t *testing.T) {
	r, _ := NewRouter()
	req, rec := testhelper.MakeRequest(t,
		http.MethodGet,
		"/",
		nil,
	)
	r.ServeHTTP(rec, req)

	var status statusResp
	err := json.Unmarshal(rec.Body.Bytes(), &status)
	assert.Nil(t, err)
	assert.Equal(t, "live", status.Status)
	assert.Equal(t, config.AppEnvTesting, status.Env)
	assert.Equal(t, http.StatusOK, rec.Code)
}
