// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: follows.sql

package db

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createFollow = `-- name: CreateFollow :one
INSERT INTO "follows" (id, create_time, update_time, user_id, feed_id)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, create_time, update_time, user_id, feed_id
`

type CreateFollowParams struct {
	ID         uuid.UUID
	CreateTime time.Time
	UpdateTime time.Time
	UserID     uuid.UUID
	FeedID     uuid.UUID
}

func (q *Queries) CreateFollow(ctx context.Context, arg CreateFollowParams) (Follow, error) {
	row := q.db.QueryRowContext(ctx, createFollow,
		arg.ID,
		arg.CreateTime,
		arg.UpdateTime,
		arg.UserID,
		arg.FeedID,
	)
	var i Follow
	err := row.Scan(
		&i.ID,
		&i.CreateTime,
		&i.UpdateTime,
		&i.UserID,
		&i.FeedID,
	)
	return i, err
}

const getUserFollows = `-- name: GetUserFollows :many
SELECT id, create_time, update_time, user_id, feed_id FROM "follows" WHERE user_id = $1
`

func (q *Queries) GetUserFollows(ctx context.Context, userID uuid.UUID) ([]Follow, error) {
	rows, err := q.db.QueryContext(ctx, getUserFollows, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Follow
	for rows.Next() {
		var i Follow
		if err := rows.Scan(
			&i.ID,
			&i.CreateTime,
			&i.UpdateTime,
			&i.UserID,
			&i.FeedID,
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
