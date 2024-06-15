package main

import (
	"log"

	"github.com/duckysmacky/rss-server/internal/handlers"
)

func main() {
	var port = 8080
	var server = handlers.NewServer("localhost", port)

	log.Printf("Server listening on port %v", port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}