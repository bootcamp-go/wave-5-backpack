package products

import (
	"context"
	"goweb/internal/domain"
)

type Service interface {
	GetByName(string) (domain.Product, error)
	Store(domain.Product) (domain.Product, error)
	GetAll(context.Context) ([]domain.Product, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetByName(name string) (domain.Product, error) {
	product, err := s.repository.GetByName(name)
	return product, err
}

func (s *service) Store(product domain.Product) (domain.Product, error) {
	product, err := s.repository.Store(product)
	return product, err
}

func (s *service) GetAll(ctx context.Context) ([]domain.Product, error) {
	return s.repository.GetAll(ctx)
}
