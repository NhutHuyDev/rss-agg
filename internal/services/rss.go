package services

import (
	"encoding/xml"
	"io"
	"net/http"
	"time"

	"github.com/NhutHuyDev/rss-agg/api"
)

func FetchRSSFeedFromURL(url string) (api.RSSFeed, error) {
	httpClient := http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := httpClient.Get(url)
	if err != nil {
		return api.RSSFeed{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return api.RSSFeed{}, err
	}

	rssFeed := api.RSSFeed{}
	err = xml.Unmarshal(data, &rssFeed)
	if err != nil {
		return api.RSSFeed{}, err
	}

	return rssFeed, nil
}
