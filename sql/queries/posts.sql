-- name: CreatePost :one
INSERT INTO posts (id, create_time, update_time, publish_date, url, feed_id, title, description)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- name: GetUserPosts :many
SELECT posts.* from posts
JOIN follows ON posts.feed_id = follows.feed_id
WHERE follows.user_id = $1
ORDER BY posts.publish_date DESC
LIMIT $2;