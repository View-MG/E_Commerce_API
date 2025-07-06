package service

import "github.com/View-MG/be-project/internal/repository"

type service struct {
	repo repository.AllRepo
}

type AllOfService struct {
	User    UserService
	Product ProductService
	Cart    CartService
	Order   OrderService
}

func NewService(repo repository.AllRepo) AllOfService {
	srv := &service{
		repo: repo,
	}
	return AllOfService{
		User:    srv,
		Product: srv,
		Cart:    srv,
		Order:   srv,
	}
}
