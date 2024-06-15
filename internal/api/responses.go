package api

import (
	"net/http"
)

func ResponseBadRequest(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, http.StatusBadRequest, "Bad request", "Something went wrong!")
}

func ConfirmStatus(w http.ResponseWriter, r *http.Request) {
	respondWithStatus(w)
}