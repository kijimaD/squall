package fetcher

import (
	"fmt"
	"squall/factories"
	"squall/models"
	"testing"

	"github.com/rs/xid"
	"github.com/stretchr/testify/assert"
)

func TestGetFeedReal(t *testing.T) {
	rfg := RealFeedGetter{}
	urls, err := rfg.GetFeed("https://go.dev/blog/feed.atom?format=xml")
	assert.True(t, 0 < len(urls))
	assert.NoError(t, err)
}

func TestGetFeedMock(t *testing.T) {
	mfg := MockFeedGetter{}
	urls, err := mfg.GetFeed("mock")
	assert.NoError(t, err)
	assert.Equal(t, 1, len(urls))
	assert.Equal(t, "https://google.com", urls[0])
}

func TestFetchFeeds_取得できる(t *testing.T) {
	// レコードが作成されている
	beforeID := models.GetMaxID(getDB(), models.Entry{})
	defer func() {
		assert.NotEqual(t, beforeID, models.GetMaxID(getDB(), models.Entry{}))
	}()

	mock := MockFeedGetter{}
	err := FetchFeeds(feedSources{feedSource{Desc: "desc", URL: "mock url"}}, &mock)
	assert.NoError(t, err)
}

func TestCreateEntry_作成できる(t *testing.T) {
	// レコードが作成されている
	beforeID := models.GetMaxID(getDB(), models.Entry{})
	defer func() {
		assert.NotEqual(t, beforeID, models.GetMaxID(getDB(), models.Entry{}))
	}()

	err := createEntry(fmt.Sprintf("url-%s", xid.New().String()))
	assert.NoError(t, err)
}

func TestCreateEntry_URLがすでに存在していると作成しない(t *testing.T) {
	id := xid.New().String()
	url := fmt.Sprintf("url-%s", id)
	var deps []factories.Dependency
	_, deps = factories.MakeEntry(factories.Fields{"URL": url}, deps)
	for _, m := range deps {
		assert.NoError(t, getDB().Create(m).Error)
	}

	// レコードが作成されてない
	beforeID := models.GetMaxID(getDB(), models.Entry{})
	defer func() {
		assert.Equal(t, beforeID, models.GetMaxID(getDB(), models.Entry{}))
	}()

	err := createEntry(url)
	assert.NoError(t, err)
}
