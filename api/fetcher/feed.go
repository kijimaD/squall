package fetcher

import (
	"fmt"
	"log"

	"github.com/mmcdole/gofeed"
)

// type FeedGetter interface {
// 	GetFeed(url string) ([]byte, error)
// }

// 登録されているフィードソースから取得する
// TODO: ここでインターフェースを渡すようにする
func FetchFeeds(feedSources feedSources) error {
	for _, f := range feedSources {
		err := getFeedURL(f.URL)
		if err != nil {
			log.Println(err)
		}
	}
	return nil
}

// フィードソースから各フィードエントリを取得する
func getFeedURL(url string) error {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(url)
	if err != nil {
		return err
	}

	for _, f := range feed.Items {
		fmt.Printf("%#v\n", f.Link)
	}
	return nil
}
