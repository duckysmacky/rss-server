// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: feeds.sql

package db

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createFeed = `-- name: CreateFeed :one
INSERT INTO feeds (id, create_time, update_time, url, name, user_id) 
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id, create_time, update_time, url, name, user_id
`

type CreateFeedParams struct {
	ID         uuid.UUID
	CreateTime time.Time
	UpdateTime time.Time
	Url        string
	Name       string
	UserID     uuid.UUID
}

func (q *Queries) CreateFeed(ctx context.Context, arg CreateFeedParams) (Feed, error) {
	row := q.db.QueryRowContext(ctx, createFeed,
		arg.ID,
		arg.CreateTime,
		arg.UpdateTime,
		arg.Url,
		arg.Name,
		arg.UserID,
	)
	var i Feed
	err := row.Scan(
		&i.ID,
		&i.CreateTime,
		&i.UpdateTime,
		&i.Url,
		&i.Name,
		&i.UserID,
	)
	return i, err
}

const getFeeds = `-- name: GetFeeds :many
SELECT id, create_time, update_time, url, name, user_id FROM feeds
`

func (q *Queries) GetFeeds(ctx context.Context) ([]Feed, error) {
	rows, err := q.db.QueryContext(ctx, getFeeds)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Feed
	for rows.Next() {
		var i Feed
		if err := rows.Scan(
			&i.ID,
			&i.CreateTime,
			&i.UpdateTime,
			&i.Url,
			&i.Name,
			&i.UserID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
