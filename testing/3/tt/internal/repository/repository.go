package repository

import (
	"testing/3/tt/internal/domain"
	"testing/3/tt/pkg/store"
)

type Repository interface {
	Read(id int) (domain.Product, error)
	ReadAll() ([]domain.Product, error)
	Create(name string, price float64, quantity int) (domain.Product, error)
	Update(id int, name string, price float64, quantity int) (domain.Product, error)
	UpdateNamePrice(id int, name string, price float64) (domain.Product, error)
	Delete(id int) error
	LastId() (int, error)
}

type repository struct {
	db store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{db: db}
}