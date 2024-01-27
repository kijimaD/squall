package fetcher

import "io"

func Run(r io.Reader) error {
	feedSources, err := loadFeedSources(r)
	if err != nil {
		return err
	}

	rfg := RealFeedGetter{}
	err = FetchFeeds(feedSources, &rfg)
	if err != nil {
		return err
	}

	return nil
}
