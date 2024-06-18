package handlers

import (
	"database/sql"
	"log"

	"github.com/duckysmacky/rss-server/internal/db"
)

type DatabaseConfig struct {
	Queries *db.Queries
}

var database DatabaseConfig

func ConnectDatabase(address string) {
	var conn, err = sql.Open("postgres", address)
	if err != nil {
		log.Fatal(err)
	}

	var queries = db.New(conn)

	database = DatabaseConfig {
		Queries: queries,
	}
}