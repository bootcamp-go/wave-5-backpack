package service

import (
	"testing/3/tm/internal/domain"
	"testing/3/tm/internal/repository"
)

type Service interface {
	Read(id int) (domain.Product, error)
	ReadAll() ([]domain.Product, error)
	Create(name string, price float64, quantity int) (domain.Product, error)
	Update(id int, name string, price float64, quantity int) (domain.Product, error)
	UpdateNamePrice(id int, name string, price float64) (domain.Product, error)
	Delete(id int) error
}

type service struct {
	r repository.Repository
}

func NewService(r repository.Repository) Service {
	return &service{r: r}
}
