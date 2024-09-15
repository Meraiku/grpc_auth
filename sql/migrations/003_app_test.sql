-- +goose Up
INSERT INTO app (id, name, secret) VALUES (1, 'test', 'test_secret');

-- +goose Down
DELETE FROM app WHERE id = 1;