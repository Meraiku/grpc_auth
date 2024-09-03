-- +goose Up
CREATE TABLE app(
    id INTEGER PRIMARY KEY,
    name TEXT UNIQUE NOT NULL,
    secret TEXT UNIQUE NOT NULL
);

-- +goose Down
DROP TABLE app;