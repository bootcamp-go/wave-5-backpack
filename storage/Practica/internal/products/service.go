package products

import (
	"errors"
	"fmt"

	"github.com/bootcamp-go/wave-5-backpack/tree/Ramos_Andres/goweb/practica/internal/domain"
)

type Service interface {
	GetAll() ([]domain.Product, error)
	Store(domain.Product) (domain.Product, error)
	GetById(id uint64) (domain.Product, error)
	UpdateTotal(product domain.Product) (domain.Product, error)
	UpdatePartial(product domain.Product) (domain.Product, error)
	Delete(id uint64) (domain.Product, error)
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
	products, err := s.repository.GetAll()
	if err != nil {
		return nil, errors.New("no se pudo obtener los productos")
	}
	return products, nil
}

func (s *service) Store(product domain.Product) (domain.Product, error) {
	producto, err := s.repository.Store(product)
	if err != nil {
		return domain.Product{}, errors.New("no se pudo guardar el producto")
	}

	return producto, nil
}

func (s *service) GetById(id uint64) (domain.Product, error) {
	producto, err := s.repository.GetById(id)
	if err != nil {
		return domain.Product{}, fmt.Errorf("no se pudo encontrar el producto con el id: %d", id)
	}
	return producto, nil
}

func (s *service) UpdateTotal(product domain.Product) (domain.Product, error) {
	producto, err := s.repository.Update(product)
	if err != nil {
		return domain.Product{}, err
	}
	return producto, nil
}

func (s *service) UpdatePartial(product domain.Product) (domain.Product, error) {
	producto, err := s.repository.Update(product)
	if err != nil {
		return domain.Product{}, err
	}
	return producto, nil
}

func (s *service) Delete(id uint64) (domain.Product, error) {
	producto, err := s.repository.Delete(id)
	if err != nil {
		return domain.Product{}, fmt.Errorf("no se pudo encontrar el producto con el id: %d", id)
	}
	return producto, nil
}
