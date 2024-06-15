package handlers

import (
	"fmt"
	"net/http"
)

func NewServer[T string | int](addr string, port T) *http.Server {
	var server = http.Server {
		Handler: newRouter(),
		Addr: fmt.Sprintf("%v:%v", addr, port),
	}

	return &server
}