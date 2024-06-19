-- +goose Up

CREATE TABLE posts (
    id UUID PRIMARY KEY,
    create_time TIMESTAMP NOT NULL,
    update_time TIMESTAMP NOT NULL,
    publish_date TIMESTAMP NOT NUlL,
    url TEXT NOT NULL UNIQUE,
    feed_id UUID NOT NULL REFERENCES feeds(id) ON DELETE CASCADE,
    title TEXT NOT NULL,
    description TEXT
);

-- +goose Down

DROP TABLE posts;