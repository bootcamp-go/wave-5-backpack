package products

import "github.com/abelardolugo/go-web/internal/domain"

type Service interface {
	GetAll() ([]domain.Product, error)
	Store(nombre string, cantidad int, precio float64) (domain.Product, error)
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

func (s *service) Store(nombre string, cantidad int, precio float64) (domain.Product, error) {
	lastID, err := s.repository.LastID()
	if err != nil {
		return domain.Product{}, err
	}

	lastID++

	producto, err := s.repository.Store(lastID, nombre, cantidad, precio)
	if err != nil {
		return domain.Product{}, err
	}

	return producto, nil
}
