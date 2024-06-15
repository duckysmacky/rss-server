package api

import (
	"time"

	"github.com/duckysmacky/rss-server/internal/db"
	"github.com/google/uuid"
)

type User struct {
	ID         uuid.UUID `json:"id"`
	Createtime time.Time `json:"createTime"`
	Updatetime time.Time `json:"updateTime"`
	Username   string `json:"username"`
}

func FormatUserJSON(u db.User) User {
	return User {
		ID: u.ID,
		Createtime: u.Createtime,
		Updatetime: u.Updatetime,
		Username: u.Username,
	}
}