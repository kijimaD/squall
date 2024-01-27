package fetcher

import (
	"bytes"
	"io"
	"os"

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

func loadFeedSourcesFromFile() (feedSources, error) {
	b, err := os.ReadFile("./feeds.yml")
	if err != nil {
		return nil, err
	}
	f, err := loadFeedSources(bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	return f, nil
}
