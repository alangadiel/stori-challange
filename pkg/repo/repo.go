package repo

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

type Repository struct {
	dbConn *pgx.Conn
}

func CreateRepository(ctx context.Context) (Repository, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:5432/%s?sslmode=disable",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_DB"))

	conn, err := pgx.Connect(ctx, connStr)
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
