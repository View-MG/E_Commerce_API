package repository

import (
	"github.com/jmoiron/sqlx"
)

type repository struct {
	db *sqlx.DB
}

type AllRepo struct {
	Order OrderRepository
	User UserRepository
}

func NewRepository(db *sqlx.DB) AllRepo {
	repo := &repository{db: db}
	return AllRepo{
		Order: repo,
		User: repo,
	}

}
