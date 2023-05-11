-- +goose Up
-- +goose StatementBegin
ALTER TABLE users ADD COLUMN tags TEXT[];
ALTER TABLE users ADD COLUMN metadata jsonb NOT NULL DEFAULT '{}';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users DROP COLUMN IF EXISTS tags;
ALTER TABLE users DROP COLUMN IF EXISTS metadata;
-- +goose StatementEnd
