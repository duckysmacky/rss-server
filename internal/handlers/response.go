package handlers

import (
	"net/http"

	"github.com/duckysmacky/rss-server/internal/api"
)

func ResponseStatus(w http.ResponseWriter, r *http.Request) {
	api.RespondStatus(w)
}