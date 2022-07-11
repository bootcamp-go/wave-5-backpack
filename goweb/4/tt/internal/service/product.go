package service

import (
	"errors"
	"goweb/4/tt/internal/domain"
)

func (s *service) Read(id int) (domain.Product, error) {
	product, err := s.r.Read(id)
	if err != nil {
		return domain.Product{}, err
	}

	return product, nil
}

func (s *service) ReadAll() ([]domain.Product, error) {
	products, err := s.r.ReadAll()
	if err != nil {
		return []domain.Product{}, err
	}
	if len(products) == 0 {
		return []domain.Product{}, errors.New("no se han encontrado productos en el listado")
	}
	return products, nil
}

func (s *service) Create(name string, price float64, quantity int) (domain.Product, error) {
	products, err := s.r.ReadAll()
	if err != nil {
		return domain.Product{}, err
	}

	if len(products) != 0 {
		for _, p := range products {
			if p.Name == name && p.Price == price && p.Quantity == quantity {
				return domain.Product{}, errors.New("el producto ingresado ya existe")
			}
		}
	}

	product, err := s.r.Create(name, price, quantity)
	if err != nil {
		return domain.Product{}, err
	}

	return product, nil
}

func (s *service) Update(id int, name string, price float64, quantity int) (domain.Product, error) {
	updatedProduct, err := s.r.Update(id, name, price, quantity)
	if err != nil {
		return domain.Product{}, err
	}

	return updatedProduct, nil
}

func (s *service) UpdateNamePrice(id int, name string, price float64) (domain.Product, error) {
	updatedProduct, err := s.r.UpdateNamePrice(id, name, price)
	if err != nil {
		return domain.Product{}, err
	}

	return updatedProduct, nil
}

func (s *service) Delete(id int) error {
	err := s.r.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
