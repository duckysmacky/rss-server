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
	Username   string `json:"username"`
	APIKey string `json:"apiKey"`
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