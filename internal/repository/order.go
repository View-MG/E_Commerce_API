package repository

import "github.com/View-MG/be-project/internal/entity"

type OrderRepository interface {
	Save(order entity.Order) error
}

func (r repository) Save(order entity.Order) error {
	return nil
}
