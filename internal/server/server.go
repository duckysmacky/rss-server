package server

import (
	"fmt"
	"net/http"

	"github.com/duckysmacky/rss-server/internal/routers"
)

func NewServer[T string | int](addr string, port T) *http.Server {
	var server = http.Server {
		Handler: routers.NewRouter(),
		Addr: fmt.Sprintf("%v:%v", addr, port),
	}

	return &server
}