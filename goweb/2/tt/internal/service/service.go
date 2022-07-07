package service

import (
	"goweb/2/tt/internal/domain"
	"goweb/2/tt/internal/repository"
)

type Service interface {
	ListAll() ([]domain.Product, error)
	Store(name string, price float64) (domain.Product, error)
}

type service struct {
	r repository.Repository
}

func NewService(r repository.Repository) Service {
	return &service{r: r}
}
