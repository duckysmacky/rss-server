package api

import (
	"time"

	"github.com/duckysmacky/rss-server/internal/db"
	"github.com/google/uuid"
)

type User struct {
	ID         uuid.UUID `json:"id"`
	CreateTime time.Time `json:"createTime"`
	UpdateTime time.Time `json:"updateTime"`
	Username   string    `json:"username"`
	APIKey     string    `json:"apiKey"`
}

type Feed struct {
	ID         uuid.UUID `json:"id"`
	CreateTime time.Time `json:"createTime"`
	UpdateTime time.Time `json:"updateTime"`
	Url        string    `json:"url"`
	Name       string    `json:"name"`
	UserID     uuid.UUID `json:"userId"`
}

type Follow struct {
	ID         uuid.UUID `json:"id"`
	CreateTime time.Time `json:"createTime"`
	UpdateTime time.Time `json:"updateTime"`
	UserID     uuid.UUID `json:"userId"`
	FeedID     uuid.UUID `json:"feedId"`
}

func FormatUserJSON(u db.User) User {
	return User {
		ID: u.ID,
		CreateTime: u.CreateTime,
		UpdateTime: u.UpdateTime,
		Username: u.Username,
		APIKey: u.ApiKey,
	}
}

func FormatFeedJSON(f db.Feed) Feed {
	return Feed {
		ID: f.ID,
		CreateTime: f.CreateTime,
		UpdateTime: f.UpdateTime,
		Url: f.Url,
		Name: f.Name,
		UserID: f.UserID,
	}
}

func FormatFeedsJSON(f []db.Feed) []Feed {
	var feeds = []Feed {}
	for _, feed := range f {
		feeds = append(feeds, FormatFeedJSON(feed))
	}

	return feeds
}

func FormatFollowJSON(f db.Follow) Follow {
	return Follow {
		ID: f.ID,
		CreateTime: f.CreateTime,
		UpdateTime: f.UpdateTime,
		UserID: f.UserID,
		FeedID: f.FeedID,
	}
}

func FormatFollowsJSON(f []db.Follow) []Follow {
	var follows = []Follow {}
	for _, follow := range f {
		follows = append(follows, FormatFollowJSON(follow))
	}

	return follows
}