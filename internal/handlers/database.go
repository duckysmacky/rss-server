package handlers

import (
	"database/sql"
	"log"

	"github.com/duckysmacky/rss-server/internal/db"
)

type Database struct {
	Queries *db.Queries
}

var database Database

func ConnectDatabase(address string) {
	var conn, err = sql.Open("postgres", address)
	if err != nil {
		log.Fatal(err)
	}

	var queries = db.New(conn)

	database = Database {
		Queries: queries,
	}
}