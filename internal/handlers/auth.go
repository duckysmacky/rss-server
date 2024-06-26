package handlers

import (
	"errors"
	"net/http"
	"strings"

	"github.com/duckysmacky/rss-server/internal/api"
	"github.com/duckysmacky/rss-server/internal/db"
)

// callback function signature
type authHandler func(http.ResponseWriter, *http.Request, db.User)

// Accepts a callback function to call after authorizing user via Authorization header with an API key
func (d DatabaseConfig) AuthUser(handler authHandler) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request)  {
		// getting the key
		key, err := getAPIKey(r.Header)
		if err != nil {
			api.RespondWithError(w, http.StatusForbidden, "An error occured while trying to authenticate", err)
			return
		}
	
		// getting user with api key
		user, err := d.Queries.GetUserByApiKey(r.Context(), key)
		if err != nil {
			api.RespondWithError(w, http.StatusNotFound, "An error occured while trying to get user", err)
			return
		}
	
		// calling the callback function with found user
		handler(w, r, user)
	}
}

// Fetches an API key from request's header (Authorization: ApiKey <key>)
func getAPIKey(h http.Header) (string, error) {
	// find the authorization header data
	var auth = h.Get("Authorization")
	if auth == "" {
		return "", errors.New("header: no authorization data found")
	}

	// parse data
	var data = strings.Split(auth, " ")
	if len(data) != 2 || data[0] != "ApiKey" {
		return "", errors.New("header: invalid authorization info")
	}

	return data[1], nil
}