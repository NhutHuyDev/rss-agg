package main

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/NhutHuyDev/rss-agg/internal/db"
)

func StartScraping(
	queries *db.Queries,
	concurrencyNum int,
	timeBetweenReq time.Duration,
) {
	log.Printf("Scraping on %v goroutines every %s duration", concurrencyNum, timeBetweenReq)
	ticker := time.NewTicker(timeBetweenReq)

	for ; ; <-ticker.C {
		feeds, err := queries.GetNextFeedsToFetch(
			context.Background(),
			int32(concurrencyNum),
		)

		if err != nil {
			log.Printf("error fetching next feeds for scraping: %v", err)
			continue
		}

		wg := &sync.WaitGroup{}
		for _, feed := range feeds {
			wg.Add(1)

			go scrapeFeed(queries, wg, feed)

			wg.Wait()
		}

	}
}

func scrapeFeed(queries *db.Queries, wg *sync.WaitGroup, feed db.Feed) {
	defer wg.Done()

	_, err := queries.MarkFeedAsFetched(context.Background(), feed.ID)
	if err != nil {
		log.Printf("error marking feed as fetched: %v", err)
		return
	}

	rssFeed, err := FetchRSSFeedFromURL(feed.Url)
	if err != nil {
		log.Printf("error fetching feed: %v", err)
		return
	}

	// for _, item := range rssFeed.Channel.Item {

	// }

	log.Printf("feed %s collected, %v posts found", feed.Name, len(rssFeed.Channel.Item))
}
