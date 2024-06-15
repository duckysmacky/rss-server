package handlers

import (
	"fmt"
	"net/http"
)

func NewServer[T string | int](addr string, port T, database Database) *http.Server {
	return &http.Server{
		Handler: newRouter(database),
		Addr:    fmt.Sprintf("%v:%v", addr, port),
	}
}
