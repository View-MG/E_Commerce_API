package adapters

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type PostgresAdapter struct {
	DB *sqlx.DB
}

func NewPostgresAdapter(username, password, host, port, dbname string) (*PostgresAdapter, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", username, password, host, port, dbname)
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect: %w", err)
	}

	return &PostgresAdapter{DB: db}, nil
}
