package main

import (
	"log"
	"os"

	"github.com/duckysmacky/rss-server/internal/server"
	"github.com/joho/godotenv"
)

func main() {
	// load .env
	if err := godotenv.Load(); err != nil {
	  log.Fatal("Error loading .env file")
	}

	var port = os.Getenv("PORT")
	var server = server.NewServer("localhost", port)

	log.Printf("Server starting on port %v", port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}