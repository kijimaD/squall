package fetcher

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadFeedSources(t *testing.T) {
	r := strings.NewReader(`
- desc: Go blog
  url: https://go.dev/blog/feed.atom?format=xml
`)
	feeds, err := loadFeedSources(r)
	assert.NoError(t, err)
	expect := feedSources{
		feedSource{Desc: "Go blog", URL: "https://go.dev/blog/feed.atom?format=xml"},
	}
	assert.Equal(t, expect, feeds)
}
