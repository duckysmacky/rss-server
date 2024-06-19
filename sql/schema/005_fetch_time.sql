-- +goose Up

ALTER TABLE feeds
ADD COLUMN fetch_time TIMESTAMP;

-- +goose Down

ALTER TABLE feeds
DROP COLUMN fetch_time;