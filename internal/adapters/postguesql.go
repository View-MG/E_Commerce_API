package adapters

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type PostgresAdapter struct {
	DB *sqlx.DB
}

func NewPostgresAdapter(dsn string) (*PostgresAdapter, error) {
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect: %w", err)
	}

	return &PostgresAdapter{DB: db}, nil
}
