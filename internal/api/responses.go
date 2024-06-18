package api

import (
	"errors"
	"net/http"
)

func ResponseStatus(w http.ResponseWriter, r *http.Request) {
	respondWithStatus(w)
}

func ResponseBadRequest(w http.ResponseWriter, r *http.Request) {
	RespondWithError(w, http.StatusBadRequest, "Bad request", errors.New("something went wrong"))
}

func ResponseCreated(w http.ResponseWriter, r *http.Request) {
	RespondWithCode(w, http.StatusCreated, "Success")
}