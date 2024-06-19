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

type Post struct {
	ID          uuid.UUID `json:"id"`
	CreateTime  time.Time `json:"createTime"`
	UpdateTime  time.Time `json:"updateTime"`
	PublishDate time.Time `json:"publishDate"`
	Url         string	  `json:"url"`
	FeedID      uuid.UUID `json:"feedId"`
	Title       string 	  `json:"title"`
	Description *string   `json:"description"`
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

func FormatPostJSON(p db.Post) Post {
	var description *string
	if p.Description.Valid {
		description = &p.Description.String
	}

	return Post {
		ID: p.ID,
		CreateTime: p.CreateTime,
		UpdateTime: p.UpdateTime,
		PublishDate: p.PublishDate,
		Url: p.Url,
		FeedID: p.FeedID,
		Title: p.Title,
		Description: description,
	}
}

func FormatPostsJSON(p []db.Post) []Post {
	var posts = []Post {}
	for _, post := range p {
		posts = append(posts, FormatPostJSON(post))
	}

	return posts
}