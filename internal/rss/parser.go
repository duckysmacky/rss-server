package rss

import (
	"encoding/xml"
	"io"
	"net/http"
	"time"
)

type Item struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	PubDate     string `xml:"pubDate"`
	GUID        string `xml:"guid"`
	Description string `xml:"description"`
}

type RSSFeed struct {
	Channel struct {
		Title         string `xml:"title"`
		Link          string `xml:"link"`
		Description   string `xml:"description"`
		Generator     string `xml:"generator"`
		Language      string `xml:"language"`
		LastBuildDate string `xml:"lastBuildDate"`
		Items         []Item `xml:"item"`
	} `xml:"channel"`
}

func parseRSS(url string) (*RSSFeed, error) {
	var rssFeed = RSSFeed {}
	// new http client to access the data
	var client = http.Client {
		Timeout: 10 * time.Second,
	}

	// accessing data
	res, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// reading data
	rssData, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	// parsing raw xml into rss feed struct
	err = xml.Unmarshal(rssData, &rssFeed)
	if err != nil {
		return nil, err
	}

	return &rssFeed, nil
}