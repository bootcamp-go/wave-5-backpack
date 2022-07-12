package products

import (
	"fmt"
	"goweb/internal/domain"
)

type Service interface {
	GetAll() ([]domain.Product, error)
	Store(name, color string, price, stock int, code string, published bool, date string) (domain.Product, error)
	Update(id int, name, color string, price, stock int, code string, published bool, date string) (domain.Product, error)
	UpdateName(id int, name string) (domain.Product, error)
	Delete(id int) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetAll() ([]domain.Product, error) {
	ps, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}

	return ps, nil
}

func (s *service) Store(name, color string, price, stock int, code string, published bool, date string) (domain.Product, error) {
	lastID, err := s.repository.LastID()
	if err != nil {
		return domain.Product{}, fmt.Errorf("error getting product last id: %w", err)
	}
	lastID++
	product, err := s.repository.Store(lastID, name, color, price, stock, code, published, date)
	if err != nil {
		return domain.Product{}, fmt.Errorf("error creating product: %w", err)
	}
	return product, nil
}

func (s *service) Update(id int, name, color string, price, stock int, code string, published bool, date string) (domain.Product, error) {
	p, err := s.repository.Update(id, name, color, price, stock, code, published, date)
	if err != nil {
		return domain.Product{}, fmt.Errorf("error updating product %w", err)
	}
	return p, nil
}

func (s *service) UpdateName(id int, name string) (domain.Product, error) {
	product, err := s.repository.UpdateName(id, name)
	if err != nil {
		return domain.Product{}, fmt.Errorf("error updating product %w", err)
	}
	return product, nil
}

func (s *service) Delete(id int) error {
	if err := s.repository.Delete(id); err != nil {
		return fmt.Errorf("error deleting product %w", err)
	}
	return nil
}
