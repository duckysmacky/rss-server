package handlers

import (
	"encoding/json"
	"fmt"
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
		api.RespondWithError(w, http.StatusBadRequest, "Bad Request", fmt.Sprintf("An error occured while trying to parse JSON: %v", err))
		return
	}

	var user, err = d.Queries.CreateUser(r.Context(), db.CreateUserParams{
		ID: uuid.New(),
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
		Username: userData.Username,
	})
	if err != nil {
		api.RespondWithError(w, http.StatusInternalServerError, "Internal Server Error", fmt.Sprintf("An error occured while trying to create user: %v", err))
	}

	api.RespondWithJSON(w, http.StatusCreated, api.FormatUserJSON(user))
}

func (d DatabaseConfig) handleGetUserByAPIKey(w http.ResponseWriter, r *http.Request) {
	key, err := getAPIKey(r.Header)
	if err != nil {
		api.RespondWithError(w, http.StatusForbidden, "Forbidden", fmt.Sprintf("An error occured while trying to authenticate: %v", err))
		return
	}

	user, err := d.Queries.GetUserByAPIKey(r.Context(), key)
	if err != nil {
		api.RespondWithError(w, http.StatusNotFound, "Not found", fmt.Sprintf("An error occured while trying to get user by the api key: %v", err))
		return
	}

	api.RespondWithJSON(w, http.StatusOK, api.FormatUserJSON(user))
}