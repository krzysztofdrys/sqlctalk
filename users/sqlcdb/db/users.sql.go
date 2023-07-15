// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: users.sql

package db

import (
	"context"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

const getUser = `-- name: GetUser :one
SELECT id, first_name, last_name, created_at, updated_at, tags, metadata
FROM users
WHERE id = $1
`

func (q *Queries) GetUser(ctx context.Context, id uuid.UUID) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.CreatedAt,
		&i.UpdatedAt,
		pq.Array(&i.Tags),
		&i.Metadata,
	)
	return i, err
}

const getUserV2 = `-- name: GetUserV2 :many
SELECT id, id
FROM users
`

type GetUserV2Row struct {
	ID   uuid.UUID
	ID_2 uuid.UUID
}

func (q *Queries) GetUserV2(ctx context.Context) ([]GetUserV2Row, error) {
	rows, err := q.db.QueryContext(ctx, getUserV2)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetUserV2Row
	for rows.Next() {
		var i GetUserV2Row
		if err := rows.Scan(&i.ID, &i.ID_2); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const insertUser = `-- name: InsertUser :one
INSERT INTO users(id, first_name, last_name, created_at)
VALUES (gen_random_uuid(), $1, $2, now())
RETURNING id, first_name, last_name, created_at, updated_at, tags, metadata
`

type InsertUserParams struct {
	FirstName string
	LastName  string
}

func (q *Queries) InsertUser(ctx context.Context, arg InsertUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, insertUser, arg.FirstName, arg.LastName)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.CreatedAt,
		&i.UpdatedAt,
		pq.Array(&i.Tags),
		&i.Metadata,
	)
	return i, err
}