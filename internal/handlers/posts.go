package handlers

import (
	"net/http"

	"github.com/duckysmacky/rss-server/internal/api"
	"github.com/duckysmacky/rss-server/internal/db"
)


func (d DatabaseConfig) HandleGetUserPosts(w http.ResponseWriter, r *http.Request, user db.User) {
	posts, err := d.Queries.GetUserPosts(r.Context(), db.GetUserPostsParams {
		UserID: user.ID,
		Limit: 10,
	})
	if err != nil {
		api.RespondWithError(w, http.StatusInternalServerError, "An error occured while trying to get user's followed feed", err)
		return
	}

	api.RespondWithJSON(w, http.StatusOK, api.FormatPostsJSON(posts))
}