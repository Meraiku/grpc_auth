-- +goose Up
CREATE TABLE users(
    id UUID PRIMARY KEY,
    email TEXT UNIQUE NOT NULL,
    password BYTEA NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);
CREATE INDEX IF NOT EXISTS idx_email ON users (email);

-- +goose Down
DROP TABLE users;