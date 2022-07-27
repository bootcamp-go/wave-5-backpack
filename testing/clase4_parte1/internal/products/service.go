package products

import "clase4_parte1/internal/domain"

type Service interface {
	GetAll() ([]domain.Product, error)
	Store(nombre, tipo string, cantidad int, precio float64) (domain.Product, error)
	Update(id int, name, productType string, count int, price float64) (domain.Product, error)
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

func (s *service) Store(nombre, tipo string, cantidad int, precio float64) (domain.Product, error) {
	lastID, err := s.repository.LastID()
	if err != nil {
		return domain.Product{}, err
	}

	lastID++

	producto, err := s.repository.Store(lastID, nombre, tipo, cantidad, precio)
	if err != nil {
		return domain.Product{}, err
	}

	return producto, nil
}

func (s *service) Update(id int, name, productType string, count int, price float64) (domain.Product, error) {
	return s.repository.Update(id, name, productType, count, price)
}

func (s *service) UpdateName(id int, name string) (domain.Product, error) {
	return s.repository.UpdateName(id, name)
}

func (s *service) Delete(id int) error {
	return s.repository.Delete(id)
}
