package sqldb

import (
	"context"
	"database/sql"
	"fmt"
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
	DB *sql.DB
}

func (r Repository) InsertUser(ctx context.Context, u user.User) (user.User, error) {
	row := r.DB.QueryRowContext(ctx, insert, u.FirstName, u.LastName)
	result := user.User{}
	if err := row.Scan(&result.ID, &result.FirstName, &result.LastName, &result.CreatedAt); err != nil {
		return result, fmt.Errorf("failed to insert: %w", err)
	}
	return result, nil
}

func (r Repository) GetUser(ctx context.Context, id string) (user.User, error) {
	row := r.DB.QueryRowContext(ctx, get, id)
	result := user.User{}
	if err := row.Scan(&result.ID, &result.FirstName, &result.LastName, &result.CreatedAt); err != nil {
		return result, fmt.Errorf("failed to get: %w", err)
	}
	return result, nil
}
