package service

import "github.com/View-MG/be-project/internal/repository"

type service struct {
	repo repository.AllRepo
}

type AllOfService struct {
	Order OrderService
	User UserService
}

func NewService(repo repository.AllRepo) AllOfService {
	srv := &service{
		repo: repo,
	}
	return AllOfService{
		Order: srv,
		User: srv,
	}
}
