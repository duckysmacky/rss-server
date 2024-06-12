package handlers

import (
	"net/http"

	"github.com/duckysmacky/rss-server/internal/api"
)

func responseStatus(w http.ResponseWriter, r *http.Request) {
	api.RespondWithStatus(w)
}

func responseBadRequest(w http.ResponseWriter, r *http.Request) {
	api.RespondWithError(w, http.StatusBadRequest, "Bad request", "Something went wrong!")
}