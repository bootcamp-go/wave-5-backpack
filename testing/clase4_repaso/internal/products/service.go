package products

import (
	"clase4_repaso/internal/domain"
	"fmt"
)

type Service interface {
	GetAll() ([]domain.Product, error)
	Store(name, productType string, count int, price float64) (domain.Product, error)
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{repo: r}
}

func (s *service) GetAll() ([]domain.Product, error) {
	products, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (s *service) Store(name, productType string, count int, price float64) (domain.Product, error) {
	lastID, err := s.repo.LastID()
	if err != nil {
		return domain.Product{}, fmt.Errorf("error getting product last id: %w", err)
	}
	lastID++
	product, err := s.repo.Store(lastID, name, productType, count, price)
	if err != nil {
		return domain.Product{}, fmt.Errorf("error creating product: %w", err)
	}
	return product, nil
}
