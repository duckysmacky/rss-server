package rss

import (
	"context"
	"database/sql"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/duckysmacky/rss-server/internal/db"
	"github.com/google/uuid"
)

func StartScraper(queries db.Queries, feeds int, interval time.Duration) {
	log.Println("RSS Scraper has started")
	log.Printf("Fetching %v feeds every %v seconds", feeds, time.Duration.Seconds(interval))
	// set up a new ticker
	var ticker = time.NewTicker(interval)

	// upon recieving a signal from the ticker (interval passed)
	for ;; <- ticker.C {
		// fetch a list of feeds
		feeds, err := queries.GetFeedsToFetch(context.Background(), int32(feeds))
		if err != nil {
			log.Println("Error fetching feeds: ", err)
		}

		var wg = sync.WaitGroup {}
		for _, feed := range feeds {
			wg.Add(1)
			go fetchFeed(&wg, queries, feed)
		}
		wg.Wait()
	}
}

func fetchFeed(wg *sync.WaitGroup, queries db.Queries, feed db.Feed) {
	defer wg.Done()

	feed, err := queries.UpdateFetchTime(context.Background(), feed.ID)
	if err != nil {
		log.Println("Error updating feed fetch time: ", err)
		return
	}

	rssFeed, err := parseRSS(feed.Url)
	if err != nil {
		log.Println("Error parsing a RSS feed: ", err)
	}

	for _, item := range rssFeed.Channel.Items {
		publishDate, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err != nil {
			log.Printf("Error parsing date %v to Timestamp: %v", item.PubDate, err)
		}

		var description = sql.NullString {
			String: item.Description,
			Valid: item.Description != "",
		}

		_, err = queries.CreatePost(context.Background(), db.CreatePostParams {
			ID: uuid.New(),
			CreateTime: time.Now(),
			UpdateTime: time.Now(),
			PublishDate: publishDate,
			Url: item.Link,
			FeedID: feed.ID,
			Title: item.Title,
			Description: description,
		})
		if err != nil {
			// TODO - check via db request for already existing
			if strings.Contains(err.Error(), "duplicate key") {
				continue
			}
			log.Println("Error creating a new post: ", err)
		}
	}
	log.Printf("Found %v posts on feed \"%s\"", len(rssFeed.Channel.Items), feed.Name)
}