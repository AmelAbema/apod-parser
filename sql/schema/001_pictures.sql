-- +goose Up
CREATE TABLE IF NOT EXISTS apod_data
(
    id          SERIAL PRIMARY KEY,
    date        DATE UNIQUE NOT NULL,
    title       TEXT        NOT NULL,
    description TEXT,
    url         TEXT        NOT NULL,
    media_type  TEXT        NOT NULL
);

-- +goose Down
DROP TABLE apod_data;
