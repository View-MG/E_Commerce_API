package repository

import (
	"github.com/jmoiron/sqlx"
)

type repository struct {
	db *sqlx.DB
}

type AllRepo struct {
	User    UserRepository
	Product ProductRepository
	Cart    CartRepository
	Order   OrderRepository
}

func NewRepository(db *sqlx.DB) AllRepo {
	repo := &repository{db: db}
	return AllRepo{
		User:    repo,
		Product: repo,
		Cart:    repo,
		Order:   repo,
	}

}
