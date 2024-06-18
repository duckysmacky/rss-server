-- +goose Up

CREATE TABLE users (
    id UUID PRIMARY KEY,
    create_time TIMESTAMP NOT NULL,
    update_time TIMESTAMP NOT NULL,
    username TEXT NOT NULL
);

-- +goose Down

DROP TABLE users;