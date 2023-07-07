-- +goose Up
-- +goose StatementBegin
ALTER TABLE users
  ADD COLUMN apikey varchar(64) UNIQUE NOT NULL DEFAULT(
  encode(sha256(random()::text::bytea), 'hex')
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
