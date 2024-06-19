package handlers

import (
	"database/sql"

	"github.com/duckysmacky/rss-server/internal/db"
)

type DatabaseConfig struct {
	Queries *db.Queries
}

var Database DatabaseConfig

func ConnectDatabase(address string) (*DatabaseConfig, error) {
	var conn, err = sql.Open("postgres", address)
	if err != nil {
		return nil, err
	}

	var queries = db.New(conn)

	Database = DatabaseConfig {
		Queries: queries,
	}

	return &Database, nil
}