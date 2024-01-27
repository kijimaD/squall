package fetcher

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetFeedEvent(t *testing.T) {
	err := getFeedURL("https://go.dev/blog/feed.atom?format=xml")
	assert.NoError(t, err)
}
