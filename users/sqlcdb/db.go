package sqlcdb

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"github.com/krzysztofdrys/talks/users/sqlcdb/db"
	"github.com/krzysztofdrys/talks/users/user"
)

func New(sqldb *sql.DB) Repository {
	return Repository{DB: db.New(sqldb)}
}

type Repository struct {
	DB *db.Queries
}

func (r Repository) InsertUser(ctx context.Context, u user.User) (user.User, error) {
	res, err := r.DB.InsertUser(ctx, db.InsertUserParams{
		FirstName: u.FirstName,
		LastName:  u.LastName,
	})
	if err != nil {
		return user.User{}, fmt.Errorf("failed to insert: %w", err)
	}
	return user.User{
		ID:        res.ID.String(),
		FirstName: res.FirstName,
		LastName:  res.LastName,
		CreatedAt: res.CreatedAt,
	}, nil

}

func (r Repository) GetUser(ctx context.Context, id string) (user.User, error) {
	uuidID, err := uuid.Parse(id)
	if err != nil {
		return user.User{}, err
	}
	res, err := r.DB.GetUser(ctx, uuidID)
	if err != nil {
		return user.User{}, err
	}
	return user.User{
		ID:        res.ID.String(),
		FirstName: res.FirstName,
		LastName:  res.LastName,
		CreatedAt: res.CreatedAt,
	}, nil
}
