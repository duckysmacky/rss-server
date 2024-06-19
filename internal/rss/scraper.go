package rss

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/duckysmacky/rss-server/internal/db"
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

	// for _, item := range rssFeed.Channel.Item {
	// 	log.Println("Found a post: ", item.Title)
	// }
	log.Printf("Found %v posts on feed %s", len(rssFeed.Channel.Item), feed.Name)
}