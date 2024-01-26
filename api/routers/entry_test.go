package routers

import (
	"encoding/json"
	"net/http"
	"squall/factories"
	"squall/models"
	"squall/testhelper"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEntries_取得できる(t *testing.T) {
	var deps []factories.Dependency
	_, deps = factories.MakeEntry(factories.Fields{}, deps)
	for _, m := range deps {
		assert.NoError(t, getDB().Create(m).Error)
	}

	r, _ := NewRouter()
	req, rec := testhelper.MakeRequest(t,
		http.MethodGet,
		"/entries",
		nil,
	)
	r.ServeHTTP(rec, req)

	var es []models.Entry
	err := json.Unmarshal(rec.Body.Bytes(), &es)
	assert.NoError(t, err)
	assert.True(t, len(es) > 0)
}

func TestGetEntries_sizeを指定できる(t *testing.T) {
	var deps []factories.Dependency
	_, deps = factories.MakeEntry(factories.Fields{}, deps)
	_, deps = factories.MakeEntry(factories.Fields{}, deps)
	for _, m := range deps {
		assert.NoError(t, getDB().Create(m).Error)
	}

	r, _ := NewRouter()
	req, rec := testhelper.MakeRequest(t,
		http.MethodGet,
		"/entries?size=1",
		nil,
	)
	r.ServeHTTP(rec, req)

	var es []models.Entry
	err := json.Unmarshal(rec.Body.Bytes(), &es)
	assert.NoError(t, err)
	assert.True(t, len(es) == 1)
}

func TestGetEntries_sizeを超えない(t *testing.T) {
	var deps []factories.Dependency
	_, deps = factories.MakeEntry(factories.Fields{}, deps)
	_, deps = factories.MakeEntry(factories.Fields{}, deps)
	for _, m := range deps {
		assert.NoError(t, getDB().Create(m).Error)
	}

	r, _ := NewRouter()
	req, rec := testhelper.MakeRequest(t,
		http.MethodGet,
		"/entries?size=3",
		nil,
	)
	r.ServeHTTP(rec, req)

	var es []models.Entry
	err := json.Unmarshal(rec.Body.Bytes(), &es)
	assert.NoError(t, err)
	assert.True(t, len(es) <= 3)
}
