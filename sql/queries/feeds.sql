-- name: CreateFeed :one
INSERT INTO feeds (id, create_time, update_time, url, name, user_id) 
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetFeeds :many
SELECT * FROM feeds;