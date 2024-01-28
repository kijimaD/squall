package routers

import (
	"encoding/json"
	"fmt"
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
	assert.Equal(t, 200, rec.Code)

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
	assert.Equal(t, 200, rec.Code)

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
	assert.Equal(t, 200, rec.Code)

	var es []models.Entry
	err := json.Unmarshal(rec.Body.Bytes(), &es)
	assert.NoError(t, err)
	assert.True(t, len(es) <= 3)
}

func TestGetEntries_IDパラメータ指定で排除できる(t *testing.T) {
	var deps []factories.Dependency
	paramEntry, deps := factories.MakeEntry(factories.Fields{}, deps)
	_, deps = factories.MakeEntry(factories.Fields{}, deps)
	for _, m := range deps {
		assert.NoError(t, getDB().Create(m).Error)
	}

	r, _ := NewRouter()
	req, rec := testhelper.MakeRequest(t,
		http.MethodGet,
		fmt.Sprintf("/entries?ignore_ids=%d", *paramEntry.ID),
		nil,
	)
	r.ServeHTTP(rec, req)
	assert.Equal(t, 200, rec.Code)

	var es []models.Entry
	err := json.Unmarshal(rec.Body.Bytes(), &es)
	assert.NoError(t, err)
	for _, e := range es {
		assert.NotEqual(t, *paramEntry.ID, *e.ID)
	}
}

// ================

func TestPostDoneEntry_既読にできる(t *testing.T) {
	var deps []factories.Dependency
	entry, deps := factories.MakeEntry(factories.Fields{}, deps)
	for _, m := range deps {
		assert.NoError(t, getDB().Create(m).Error)
	}

	r, _ := NewRouter()
	req, rec := testhelper.MakeRequest(t,
		http.MethodPost,
		fmt.Sprintf("/entries/%d/done", *entry.ID),
		nil,
	)
	r.ServeHTTP(rec, req)
	assert.Equal(t, 200, rec.Code)

	var result models.Entry
	err := json.Unmarshal(rec.Body.Bytes(), &result)
	assert.NoError(t, err)
	assert.Equal(t, entry.ID, result.ID)
	assert.Equal(t, entry.URL, result.URL)
	assert.Equal(t, true, result.IsDone)
}

func TestPostDoneEntry_存在しないIDを指定すると404を返す(t *testing.T) {
	r, _ := NewRouter()
	req, rec := testhelper.MakeRequest(t,
		http.MethodPost,
		"/entries/9999999/done",
		nil,
	)
	r.ServeHTTP(rec, req)
	assert.Equal(t, 404, rec.Code)
}
