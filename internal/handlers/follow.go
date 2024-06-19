package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/duckysmacky/rss-server/internal/api"
	"github.com/duckysmacky/rss-server/internal/db"
	"github.com/google/uuid"
)

type newFollow struct {
	FeedID uuid.UUID `json:"feedId"`
}

func (d DatabaseConfig) HandleFollowFeed(w http.ResponseWriter, r *http.Request, user db.User) {
	var decoder = json.NewDecoder(r.Body)
	var followData = newFollow {}

	if err := decoder.Decode(&followData); err != nil {
		api.RespondWithError(w, http.StatusBadRequest, "An error occured while trying to parse JSON", err)
		return
	}

	follow, err := d.Queries.CreateFollow(r.Context(), db.CreateFollowParams {
		ID: uuid.New(),
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
		UserID: user.ID,
		FeedID: followData.FeedID,
	})
	if err != nil {
		api.RespondWithError(w, http.StatusInternalServerError, "An error occured while trying to follow a new feed", err)
		return
	}

	api.RespondWithJSON(w, http.StatusCreated, api.FormatFollowJSON(follow))
}

func (d DatabaseConfig) HandleGetUserFollows(w http.ResponseWriter, r *http.Request, user db.User) {
	follows, err := d.Queries.GetUserFollows(r.Context(), user.ID)
	if err != nil {
		api.RespondWithError(w, http.StatusNotFound, "An error occured while trying to get user's follows", err)
	}

	api.RespondWithJSON(w, http.StatusCreated, api.FormatFollowsJSON(follows))
}