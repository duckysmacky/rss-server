package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/duckysmacky/rss-server/internal/api"
	"github.com/duckysmacky/rss-server/internal/db"
	"github.com/google/uuid"
)

type newUser struct {
	Username string `json:"username"`
}

func (d DatabaseConfig) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	var decoder = json.NewDecoder(r.Body)
	var userData = newUser {}

	if err := decoder.Decode(&userData); err != nil {
		api.RespondWithError(w, http.StatusBadRequest, "An error occured while trying to parse JSON", err)
		return
	}

	user, err := d.Queries.CreateUser(r.Context(), db.CreateUserParams{
		ID: uuid.New(),
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
		Username: userData.Username,
	})
	if err != nil {
		api.RespondWithError(w, http.StatusInternalServerError, "An error occured while trying to create user", err)
	}

	api.RespondWithJSON(w, http.StatusCreated, api.FormatUserJSON(user))
}

func (d DatabaseConfig) handleGetUserByAPIKey(w http.ResponseWriter, r *http.Request) {
	key, err := getAPIKey(r.Header)
	if err != nil {
		api.RespondWithError(w, http.StatusForbidden, "An error occured while trying to authenticate", err)
		return
	}

	user, err := d.Queries.GetUserByAPIKey(r.Context(), key)
	if err != nil {
		api.RespondWithError(w, http.StatusNotFound, "An error occured while trying to get user", err)
		return
	}

	api.RespondWithJSON(w, http.StatusOK, api.FormatUserJSON(user))
}