package products

import (
	"github.com/bootcamp-go/wave-5-backpack/goweb/internal/domain"
)

type Service interface {
	GetAll() ([]domain.Product, error)
	Store(nombre string, color string, precio float64, stock int, codigo string, publicado bool, FechaCreacion string) (domain.Product, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s service) GetAll() ([]domain.Product, error) {
	products, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (s service) Store(nombre string, color string, precio float64, stock int, codigo string, publicado bool, FechaCreacion string) (domain.Product, error) {
	newID, err := s.repository.LastID()
	if err != nil {
		return domain.Product{}, err
	}

	newProduct, err := s.repository.Store(newID, nombre, color, precio, stock, codigo, publicado, FechaCreacion)
	if err != nil {
		return domain.Product{}, err
	}

	return newProduct, nil
}
