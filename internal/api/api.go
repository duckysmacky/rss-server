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

type codeResponse struct {
	Message string `json:"message"`
}

func respondWithStatus(w http.ResponseWriter) {
	var response = statusResponse {Status: "Alive"}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	json.NewEncoder(w).Encode(response)
}

func RespondWithError(w http.ResponseWriter, code int, err string, message string) {
	var response = errorResponse {err, message}
	
	log.Printf("Responding an error %v: %v\n", code, message)
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	
	json.NewEncoder(w).Encode(response)
}

func RespondWithCode(w http.ResponseWriter, code int, message string) {
	var response = codeResponse {message}

	log.Printf("Responding with a code %v: %v\n", code, message)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(response)
}

func RespondWithJSON(w http.ResponseWriter, code int, response interface{}) {
	log.Printf("Responding with a code %v\n", code)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(response)
}