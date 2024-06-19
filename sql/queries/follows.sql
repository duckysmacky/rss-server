-- name: CreateFollow :one
INSERT INTO follows (id, create_time, update_time, user_id, feed_id)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetFollows :many
SELECT * FROM follows
WHERE user_id = $1;

-- name: DeleteFollow :exec
DELETE FROM follows
WHERE user_id = $1 AND id = $2;