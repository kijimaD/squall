package fetcher

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	r := strings.NewReader(`
- desc: Go blog
  url: https://go.dev/blog/feed.atom?format=xml
`)
	err := Run(r)
	assert.NoError(t, err)
}
