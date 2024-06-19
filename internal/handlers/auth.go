package handlers

import (
	"errors"
	"net/http"
	"strings"

	"github.com/duckysmacky/rss-server/internal/api"
	"github.com/duckysmacky/rss-server/internal/db"
)

type authHandler func(http.ResponseWriter, *http.Request, db.User)

func (d DatabaseConfig) AuthUser(handler authHandler) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request)  {
		key, err := getAPIKey(r.Header)
		if err != nil {
			api.RespondWithError(w, http.StatusForbidden, "An error occured while trying to authenticate", err)
			return
		}
	
		user, err := d.Queries.GetUser(r.Context(), key)
		if err != nil {
			api.RespondWithError(w, http.StatusNotFound, "An error occured while trying to get user", err)
			return
		}
	
		handler(w, r, user)
	}
}

func getAPIKey(h http.Header) (string, error) {
	var auth = h.Get("Authorization")
	if auth == "" {
		return "", errors.New("header: no authorization data found")
	}

	var data = strings.Split(auth, " ")
	if len(data) != 2 || data[0] != "ApiKey" {
		return "", errors.New("header: invalid authorization info")
	}

	return data[1], nil
}