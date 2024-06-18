package main

import (
	"log"
	"os"

	"github.com/duckysmacky/rss-server/internal/handlers"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

func main() {
	var port, dbAddr = getEnv()
	handlers.ConnectDatabase(dbAddr)

	var server = handlers.NewServer("localhost", port)
	log.Printf("Server listening on port %v", port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

}

func getEnv() (string, string) {
	godotenv.Load()

	var port = os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT not found in .env!")
	}

	var databaseAddress = os.Getenv("DATABASE_ADDRESS")
	if databaseAddress == "" {
		log.Fatal("DATABASE_ADDRESS not found in .env!")
	}

	return port, databaseAddress
}
