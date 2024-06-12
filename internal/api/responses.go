package api

import (
	"encoding/json"
	"log"
	"net/http"
)

type statusResponse struct {
	Status string `json:"status"`
}

type errorResponse struct {
	Error string `json:"error"`
	Message string `json:"message"`
}

func RespondWithError(w http.ResponseWriter, code int, err string, message string) {
	var response = errorResponse {err, message}

	log.Printf("Responding an error %v: %v\n", code, message)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(response)
}

func RespondWithStatus(w http.ResponseWriter) {
	var response = statusResponse {Status: "Alive"}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	json.NewEncoder(w).Encode(response)
}