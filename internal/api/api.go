package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type statusResponse struct {
	Status string `json:"status"`
}

type errorResponse struct {
	Message string `json:"message"`
	Error string `json:"error"`
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

func RespondWithError(w http.ResponseWriter, code int, message string, err error) {
	var response = errorResponse {message, fmt.Sprint(err)}
	
	log.Printf("Responding with an error (%v): %v\n", code, message)
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	
	json.NewEncoder(w).Encode(response)
}

func RespondWithCode(w http.ResponseWriter, code int, message string) {
	var response = codeResponse {message}

	log.Printf("Responding with a code (%v): %v\n", code, message)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(response)
}

func RespondWithJSON(w http.ResponseWriter, code int, response interface{}) {
	log.Printf("Responding with a JSON (%v)\n", code)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(response)
}