package fetcher

import (
	"squall/models"
	"testing"

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

func TestFetchFeeds(t *testing.T) {
	// レコードが作成されている
	beforeID := models.GetMaxID(getDB(), models.Entry{})
	defer func() {
		assert.NotEqual(t, beforeID, models.GetMaxID(getDB(), models.Entry{}))
	}()

	mock := MockFeedGetter{}
	err := FetchFeeds(feedSources{feedSource{Desc: "desc", URL: "mock url"}}, &mock)
	assert.NoError(t, err)
}
