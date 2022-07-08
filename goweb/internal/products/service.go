package products

import (
	"goweb/internal/domain"
)

type Service interface {
	GetAll() ([]domain.Product, error)
	Store(name, color string, price, stock int, code string, published bool, date string) (domain.Product, error)
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
		return domain.Product{}, err
	}

	lastID++

	producto, err := s.repository.Store(lastID, name, color, price, stock, code, published, date)
	if err != nil {
		return domain.Product{}, err
	}

	return producto, nil
}
