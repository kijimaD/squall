package fetcher

import (
	"log"
	"squall/models"

	"github.com/mmcdole/gofeed"
)

type GetFeeder interface {
	GetFeed(url string) ([]string, error)
}

type RealFeedGetter struct{}

// フィードソースから各フィードエントリを取得する
func (r *RealFeedGetter) GetFeed(sourceURL string) ([]string, error) {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(sourceURL)
	if err != nil {
		return nil, err
	}

	var urls []string
	for _, f := range feed.Items {
		urls = append(urls, f.Link)
	}
	return urls, nil
}

type MockFeedGetter struct{}

// フィードソースから各フィードエントリを取得する
func (r *MockFeedGetter) GetFeed(sourceURL string) ([]string, error) {
	return []string{"https://google.com"}, nil
}

// 登録されているフィードソースから取得する
func FetchFeeds(feedSources feedSources, gf GetFeeder) error {
	for _, f := range feedSources {
		urls, err := gf.GetFeed(f.URL)
		if err != nil {
			log.Println(err)
		}
		for _, u := range urls {
			err = createEntry(u)
			return err
		}
	}
	return nil
}

func createEntry(url string) error {
	entry := models.Entry{URL: url}
	var es []models.Entry
	err := getDB().Where(&entry).Find(&es).Error
	if err != nil {
		return err
	}
	if len(es) == 0 {
		// URLがまだ存在しないときだけ作成する
		err = getDB().Create(&entry).Error
		return nil
	}
	return nil
}
