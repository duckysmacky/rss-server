package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/duckysmacky/rss-server/internal/api"
	"github.com/duckysmacky/rss-server/internal/db"
	"github.com/google/uuid"
)

type newFeed struct {
	Url string  `json:"url"`
	Name string `json:"name"`
}

func (d DatabaseConfig) HandleCreateFeed(w http.ResponseWriter, r *http.Request, user db.User) {
	var decoder = json.NewDecoder(r.Body)
	var feedData = newFeed {}

	if err := decoder.Decode(&feedData); err != nil {
		api.RespondWithError(w, http.StatusBadRequest, "An error occured while trying to parse JSON", err)
		return
	}

	feed, err := d.Queries.CreateFeed(r.Context(), db.CreateFeedParams {
		ID: uuid.New(),
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
		Url: feedData.Url,
		Name: feedData.Name,
		UserID: user.ID,
	})
	if err != nil {
		api.RespondWithError(w, http.StatusInternalServerError, "An error occured while trying to create new feed", err)
		return
	}

	api.RespondWithJSON(w, http.StatusCreated, api.FormatFeedJSON(feed))
}

func (d DatabaseConfig) HandleGetFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := d.Queries.GetFeeds(r.Context())
	if err != nil {
		api.RespondWithError(w, http.StatusInternalServerError, "An error occured while trying to get feeds", err)
	}

	api.RespondWithJSON(w, http.StatusOK, api.FormatFeedsJSON(feeds))
}