-- name: CreateFeed :one
INSERT INTO feeds (id, create_time, update_time, url, name, user_id) 
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetFeeds :many
SELECT * FROM feeds;

-- name: GetFeedsToFetch :many
SELECT * FROM feeds
ORDER BY fetch_time ASC NULLS FIRST
LIMIT $1;

-- name: UpdateFetchTime :one
UPDATE feeds
SET fetch_time = now(), update_time = now()
WHERE id = $1
RETURNING *;