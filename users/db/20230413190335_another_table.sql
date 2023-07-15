-- +goose Up
-- +goose StatementBegin
create table books (
    id uuid primary key,
    owner_user_id uuid references users(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP table if exists books;
-- +goose StatementEnd
