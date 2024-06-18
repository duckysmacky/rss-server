package handlers

import (
	"errors"
	"net/http"
	"strings"
)

func getAPIKey(h http.Header) (string, error) {
	var auth = h.Get("Authorization")
	if auth == "" {
		return "", errors.New("header: no authorization data found")
	}

	var data = strings.Split(auth, " ")
	if len(data) != 2 || data[0] != "ApiKey" {
		return "", errors.New("header: invalid authorization info")
	}

	return data[1], nil
}