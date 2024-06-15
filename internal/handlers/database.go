package handlers

import (
	"database/sql"
	"log"

	"github.com/duckysmacky/rss-server/internal/db"
)

type Database struct {
	Queries *db.Queries
}

func NewDatabase(address string) *Database {
	var conn, err = sql.Open("postgres", address)
	if err != nil {
		log.Fatal(err)
	}

	var queries = db.New(conn)

	return &Database {
		Queries: queries,
	}
}