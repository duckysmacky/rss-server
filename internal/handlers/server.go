package handlers

import (
	"fmt"
	"net/http"
)

func NewServer[T string | int](addr string, port T) *http.Server {
	return &http.Server{
		Handler: newRouter(),
		Addr:    fmt.Sprintf("%v:%v", addr, port),
	}
}
