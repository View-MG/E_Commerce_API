package service

import (
	"errors"

	"github.com/View-MG/be-project/internal/entity"
)

type OrderService interface {
	CreateOrder(order entity.Order) error
}

func (s *service) CreateOrder(order entity.Order) error {
	if order.Total <= 0 {
		return errors.New("total must be positive")
	}
	// Business logic...
	if err := s.repo.Order.Save(order); err != nil {
		return err
	}
	return nil
}
