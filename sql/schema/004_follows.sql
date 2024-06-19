-- +goose Up

CREATE TABLE "follows" (
    id UUID PRIMARY KEY,
    create_time TIMESTAMP NOT NULL,
    update_time TIMESTAMP NOT NULL,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    feed_id UUID NOT NULL REFERENCES feeds(id) ON DELETE CASCADE,
    UNIQUE(user_id, feed_id)
);

-- +goose Down

DROP TABLE "follows";