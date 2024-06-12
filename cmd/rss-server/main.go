package main

import (
	"log"
	"net/http"
	"os"

	"github.com/duckysmacky/rss-server/internal/routers"
	"github.com/joho/godotenv"
)

func main() {
	// load .env
	if err := godotenv.Load(); err != nil {
	  log.Fatal("Error loading .env file")
	}

	var port = os.Getenv("PORT")

	var server = http.Server {
		Handler: routers.NewRouter(),
		Addr: "localhost:" + port,
	}

	log.Printf("Server starting on port %v", port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}