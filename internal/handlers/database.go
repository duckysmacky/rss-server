package handlers

import (
	"database/sql"
	"log"

	"github.com/duckysmacky/rss-server/internal/db"
)

type DatabaseConfig struct {
	Queries *db.Queries
}

var Database DatabaseConfig

func ConnectDatabase(address string) {
	var conn, err = sql.Open("postgres", address)
	if err != nil {
		log.Fatal(err)
	}

	var queries = db.New(conn)

	Database = DatabaseConfig {
		Queries: queries,
	}
}