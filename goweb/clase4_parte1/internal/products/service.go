package products

import (
	"fmt"

	"clase4_parte1/internal/domain"
)

type Service interface {
	GetAll() ([]domain.Product, error)
	Store(name, productType string, count int, price float64) (domain.Product, error)
	Update(id int, name, productType string, count int, price float64) (domain.Product, error)
	UpdateName(id int, name string) (domain.Product, error)
	Delete(id int) error
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

func (s *service) Update(id int, name, productType string, count int, price float64) (domain.Product, error) {
	p, err := s.repo.Update(id, name, productType, count, price)
	if err != nil {
		return domain.Product{}, fmt.Errorf("error updating product %w", err)
	}
	return p, nil
}

func (s *service) UpdateName(id int, name string) (domain.Product, error) {
	product, err := s.repo.UpdateName(id, name)
	if err != nil {
		return domain.Product{}, fmt.Errorf("error updating product %w", err)
	}
	return product, nil
}

func (s *service) Delete(id int) error {
	if err := s.repo.Delete(id); err != nil {
		return fmt.Errorf("error deleting product %w", err)
	}
	return nil
}
