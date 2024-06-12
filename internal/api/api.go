package api

import (
	"encoding/json"
	"net/http"
)

type Error struct {
	Code int
	Message string
}

type Status struct {
	Code int
	Status string
}

func RespondError(w http.ResponseWriter, code int, message string) {
	var response = Error {code, message}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(response)
}

func RespondStatus(w http.ResponseWriter) {
	var code = 200
	var response = Status {
		Code: code,
		Status: "Alive",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(response)
}