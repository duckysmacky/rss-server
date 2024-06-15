-- name: CreateUser :one
INSERT INTO users (id, createTime, updateTime, username) 
VALUES ($1, $2, $3, $4)
RETURNING *;