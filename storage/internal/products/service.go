package products

import (
	"github.com/bootcamp-go/wave-5-backpack/storage/internal/domain"
)

type Service interface {
	GetAll() ([]domain.Product, error)
	GetProductByName(name string) (domain.Product, error)
	Store(product domain.Product) (int, error)
	UpdateAll(product domain.Product) (domain.Product, error)
	Delete(id int) error
	Update(id int, nombre string, precio float64) (domain.Product, error)
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
	return nil, nil
}

func (s *service) GetProductByName(name string) (domain.Product, error) {
	product, err := s.repository.GetProductByName(name)

	if err != nil {
		return domain.Product{}, err
	}

	return product, nil
}

func (s *service) Store(product domain.Product) (int, error) {
	id, err := s.repository.Store(product)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *service) UpdateAll(product domain.Product) (domain.Product, error) {
	return domain.Product{}, nil
}

func (s *service) Delete(id int) error {
	return nil
}

func (s *service) Update(id int, nombre string, precio float64) (domain.Product, error) {
	return domain.Product{}, nil
}
