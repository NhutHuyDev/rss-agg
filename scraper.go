package main

// import (
// 	"context"
// 	"database/sql"
// 	"log"
// 	"strings"
// 	"sync"
// 	"time"

// 	"github.com/NhutHuyDev/rss-agg/internal/db"
// 	"github.com/google/uuid"
// )

// func StartScraping(
// 	queries *db.Queries,
// 	concurrencyNum int,
// 	timeBetweenReq time.Duration,
// ) {
// 	log.Printf("Scraping on %v goroutines every %s duration", concurrencyNum, timeBetweenReq)
// 	ticker := time.NewTicker(timeBetweenReq)

// 	for ; ; <-ticker.C {
// 		feeds, err := queries.GetNextFeedsToFetch(
// 			context.Background(),
// 			int32(concurrencyNum),
// 		)

// 		if err != nil {
// 			log.Printf("error fetching next feeds for scraping: %v", err)
// 			continue
// 		}

// 		wg := &sync.WaitGroup{}
// 		for _, feed := range feeds {
// 			wg.Add(1)

// 			go scrapeFeed(queries, wg, feed)

// 			wg.Wait()
// 		}

// 	}
// }

// func scrapeFeed(queries *db.Queries, wg *sync.WaitGroup, feed db.Feed) {
// 	defer wg.Done()

// 	_, err := queries.MarkFeedAsFetched(context.Background(), feed.ID)
// 	if err != nil {
// 		log.Printf("error marking feed as fetched: %v", err)
// 		return
// 	}

// 	rssFeed, err := FetchRSSFeedFromURL(feed.Url)
// 	if err != nil {
// 		log.Printf("error fetching feed: %v", err)
// 		return
// 	}

// 	for index, item := range rssFeed.Channel.Item {
// 		description := sql.NullString{}
// 		if item.Description != "" {
// 			description.String = item.Description
// 			description.Valid = true
// 		}

// 		pubAt, err := time.Parse(time.RFC1123, item.PubDate)
// 		if err != nil {
// 			log.Printf("could not parse date %v with err %v", item.PubDate, err)
// 			continue
// 		}

// 		_, err = queries.CreatePost(context.Background(), db.CreatePostParams{
// 			ID:          uuid.New(),
// 			CreatedAt:   time.Now().UTC(),
// 			UpdatedAt:   time.Now().UTC(),
// 			Title:       item.Title,
// 			Description: description,
// 			Url:         item.Link,
// 			PublishedAt: pubAt,
// 			FeedID:      feed.ID,
// 		})
// 		if err != nil {
// 			if strings.Contains(err.Error(), "duplicate key") {
// 				continue
// 			}

// 			log.Printf("failed to create post: %v", err)
// 		}

// 		log.Printf("feed %s collected. %v posts found", feed.Name, index+1)
// 	}
// }
