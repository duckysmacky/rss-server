-- +goose Up

CREATE TABLE users (
    id UUID PRIMARY KEY,
    createTime TIMESTAMP NOT NULL,
    updateTime TIMESTAMP NOT NULL,
    username TEXT NOT NULL
);

-- +goose Down

DROP TABLE users;