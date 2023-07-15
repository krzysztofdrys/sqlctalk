-- +goose Up
-- +goose StatementBegin
ALTER TABLE users ADD COLUMN updated_at TIMESTAMP NOT NULL DEFAULT now();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users DROP COLUMN IF EXISTS updated_at;
-- +goose StatementEnd
