package users

import (
	"context"
	"database/sql"
	"github.com/google/go-cmp/cmp"
	"github.com/jmoiron/sqlx"
	"github.com/krzysztofdrys/talks/users/sqlcdb"
	"github.com/krzysztofdrys/talks/users/sqldb"
	"github.com/krzysztofdrys/talks/users/sqlxdb"
	"github.com/krzysztofdrys/talks/users/user"
	"testing"

	_ "github.com/jackc/pgx/stdlib"
)

type repository interface {
	InsertUser(ctx context.Context, user user.User) (user.User, error)
	GetUser(ctx context.Context, id string) (user.User, error)
}

func TestSQL(t *testing.T) {
	db, err := sql.Open("pgx", "user=postgres password=postgres host=localhost port=5432 database=users sslmode=disable")
	if err != nil {
		t.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name string
		repo repository
	}{
		{
			name: "vanilla sql",
			repo: sqldb.Repository{DB: db},
		},
		{
			name: "sqlx",
			repo: sqlxdb.Repository{DB: sqlx.NewDb(db, "pgx")},
		},
		{
			name: "sqlc",
			repo: sqlcdb.New(db),
		},
	}

	ctx := context.Background()

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			u := user.User{FirstName: "Jan", LastName: "Kowalski"}

			insert, err := tc.repo.InsertUser(ctx, u)
			if err != nil {
				t.Fatal(err)
			}

			if insert.FirstName != u.FirstName || insert.LastName != u.LastName {
				t.Fatal(insert)
			}

			get, err := tc.repo.GetUser(ctx, insert.ID)
			if err != nil {
				t.Fatal(err)
			}

			if diff := cmp.Diff(get, insert); diff != "" {
				t.Fatal(diff)
			}
		})
	}

}
