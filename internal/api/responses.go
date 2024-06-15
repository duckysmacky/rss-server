package api

import (
	"net/http"
)

func ResponseStatus(w http.ResponseWriter, r *http.Request) {
	respondWithStatus(w)
}

func ResponseBadRequest(w http.ResponseWriter, r *http.Request) {
	RespondWithError(w, http.StatusBadRequest, "Bad request", "Something went wrong!")
}

func ResponseCreated(w http.ResponseWriter, r *http.Request) {
	RespondWithCode(w, http.StatusCreated, "Success")
}