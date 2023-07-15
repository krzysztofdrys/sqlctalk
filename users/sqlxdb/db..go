package sqlxdb

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/krzysztofdrys/talks/users/user"
)

const (
	insert = `
	INSERT INTO users(id, first_name, last_name, created_at)
	VALUES (gen_random_uuid(), $1, $2, now())
	RETURNING id, first_name, last_name, created_at
`

	get = `
		SELECT id, first_name, last_name, created_at
		FROM users
		WHERE id = $1
`
)

type Repository struct {
	DB *sqlx.DB
}

func (r Repository) InsertUser(ctx context.Context, u user.User) (user.User, error) {
	result := user.User{}
	if err := r.DB.GetContext(ctx, &result, insert, u.FirstName, u.LastName); err != nil {
		return result, fmt.Errorf("failed to insert: %w", err)
	}
	return result, nil
}

func (r Repository) GetUser(ctx context.Context, id string) (user.User, error) {
	result := user.User{}
	if err := r.DB.GetContext(ctx, &result, get, id); err != nil {
		return result, fmt.Errorf("failed to get: %w", err)
	}
	return result, nil
}
