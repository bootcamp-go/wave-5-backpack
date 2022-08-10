package products

import (
	"github.com/nictes1/storage-implementation/internal/domain"
)

type Service interface {
	Store(p domain.Product) (domain.Product, error)
	GetOne(id int) (domain.Product, error)
	Update(id int, name, productType string, count int, price float64) (domain.Product, error)
	GetAll() ([]domain.Product, error)
	Delete(id int) error
}

type service struct {
	product Repository
}

func NewService(product Repository) Service {
	return &service{
		product: product,
	}
}

func (s *service) Store(p domain.Product) (domain.Product, error) {
	product, err := s.product.Store(p)
	if err != nil {
		return domain.Product{}, err
	}

	p.ID = product.ID
	return p, nil
}

func (s *service) GetOne(id int) (domain.Product, error) {
	product, err := s.product.GetOne(id)
	if err != nil {
		return domain.Product{}, err
	}

	return product, err
}

func (s *service) GetAll() ([]domain.Product, error) {
	product, err := s.product.GetAll()
	if err != nil {
		return []domain.Product{}, err
	}

	return product, err
}

func (s *service) Update(id int, name, productType string, count int, price float64) (domain.Product, error) {
	return s.product.Update(id, name, productType, count, price)
}

func (s *service) Delete(id int) error {
	err := s.product.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
