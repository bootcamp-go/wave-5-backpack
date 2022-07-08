package repository

import "goweb/3/tm/internal/domain"

type Repository interface {
	Read(id int) (domain.Product, error)
	ReadAll() ([]domain.Product, error)
	Create(name string, price float64, quantity int) (domain.Product, error)
	Update(id int, name string, price float64, quantity int) (domain.Product, error)
	UpdateNamePrice(id int, name string, price float64) (domain.Product, error)
	Delete(id int) error
	LastId() int
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}
