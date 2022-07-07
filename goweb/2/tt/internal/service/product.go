package service

import (
	"errors"
	"goweb/2/tt/internal/domain"
)

func (s *service) ListAll() ([]domain.Product, error) {
	products, err := s.r.SelectAll()
	if err != nil {
		return []domain.Product{}, err
	}
	if products == nil {
		return []domain.Product{}, errors.New("no se han encontrado productos en el listado")
	}
	return products, nil
}

func (s *service) Store(name string, price float64) (domain.Product, error) {
	products, err := s.r.SelectAll()
	if err != nil {
		return domain.Product{}, err
	}
	for _, p := range products {
		if p.Name == name && p.Price == price {
			return domain.Product{}, errors.New("el producto ingresado ya existe")
		}
	}

	product, err := s.r.Insert(name, price)
	if err != nil {
		return domain.Product{}, err
	}

	return product, nil
}
