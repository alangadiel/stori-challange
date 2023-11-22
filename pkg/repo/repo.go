package repo

import (
	"context"

	"github.com/jackc/pgx/v5"
)

const (
	dbURL = "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
)

type Repository struct {
	dbConn *pgx.Conn
}

func CreateRepository(ctx context.Context) (Repository, error) {
	conn, err := pgx.Connect(ctx, dbURL)
	if err != nil {
		return Repository{}, err
	}

	return Repository{
		dbConn: conn,
	}, nil
}

func (r *Repository) Close(ctx context.Context) {
	r.dbConn.Close(ctx)
}
