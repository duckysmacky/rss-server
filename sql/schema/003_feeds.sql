-- +goose Up

CREATE TABLE feeds (
    id UUID PRIMARY KEY,
    create_time TIMESTAMP NOT NULL,
    update_time TIMESTAMP NOT NULL,
    url TEXT UNIQUE NOT NULL,
    name TEXT NOT NULL,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE
);

-- +goose Down

DROP TABLE feeds;