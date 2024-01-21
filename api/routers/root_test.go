package routers

import (
	"encoding/json"
	"net/http"
	"squall/config"
	"squall/consts"
	"squall/generated"
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
	assert.Equal(t, http.StatusOK, rec.Code)

	var status generated.GetRootResp
	err := json.Unmarshal(rec.Body.Bytes(), &status)
	assert.Nil(t, err)
	assert.Equal(t, "live", status.Status)
	assert.Equal(t, string(config.AppEnvTesting), status.Env)
	assert.Equal(t, consts.AppVersion, status.Version)
}
