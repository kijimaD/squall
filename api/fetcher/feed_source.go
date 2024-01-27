package fetcher

import (
	"io"

	"gopkg.in/yaml.v2"
)

type feedSource struct {
	Desc string
	URL  string
}

type feedSources []feedSource

func loadFeedSources(r io.Reader) (feedSources, error) {
	feeds := feedSources{}
	err := yaml.NewDecoder(r).Decode(&feeds)
	if err != nil {
		return nil, err
	}

	return feeds, nil
}
