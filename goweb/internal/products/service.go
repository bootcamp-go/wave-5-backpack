package products

import "goweb/internal/domain"

type Service interface {
	GetAll() ([]domain.Product, error)
	Create(name string, color string, price float64, stock int, code string, publisher bool) (domain.Product, error)
}

type service struct {
	repository Repository
}

// Create implements Service
func (s *service) Create(name string, color string, price float64, stock int, code string, publisher bool) (domain.Product, error) {
	return s.repository.Create(name, color, price, stock, code, publisher)
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
