package products

import "proyecto_meli/internal/domain"

type Service interface {
	GetAll() ([]domain.Product, error)
	//Store(nombre, tipo string, cantidad int, precio float64) (domain.Product, error)
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
