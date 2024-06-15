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

func (database Database) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	var decoder = json.NewDecoder(r.Body)
	var userData = newUser {}

	if err := decoder.Decode(&userData); err != nil {
		api.RespondWithError(w, http.StatusBadRequest, "Bad Request", fmt.Sprintf("An error occured while trying to parse JSON: %v", err))
		return
	}

	var user, err = database.Queries.CreateUser(r.Context(), db.CreateUserParams{
		ID: uuid.New(),
		Createtime: time.Now(),
		Updatetime: time.Now(),
		Username: userData.Username,
	})
	if err != nil {
		api.RespondWithError(w, http.StatusInternalServerError, "Internal Server Error", fmt.Sprintf("An error occured while trying to create user: %v", err))
	}

	api.RespondWithJSON(w, http.StatusCreated, api.FormatUserJSON(user))
}