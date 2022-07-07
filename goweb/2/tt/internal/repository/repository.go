package repository

import "goweb/2/tt/internal/domain"

type Repository interface {
	SelectAll() ([]domain.Product, error)
	Insert(name string, price float64) (domain.Product, error)
	LastId() int
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}
