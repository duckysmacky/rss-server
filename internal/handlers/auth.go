package handlers

import (
	"errors"
	"net/http"
	"strings"
)

func getAPIKey(h http.Header) (string, error) {
	var auth = h.Get("Authorization")
	if auth == "" {
		return "", errors.New("no authorization info found")
	}

	var data = strings.Split(auth, " ")
	if len(data) != 2 || data[0] != "ApiKey" {
		return "", errors.New("invalid authorization header data")
	}

	return data[1], nil
}