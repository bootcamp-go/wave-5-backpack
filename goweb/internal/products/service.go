package products

import (
	"errors"
	"fmt"

	"github.com/bootcamp-go/wave-5-backpack/tree/Ramos_Andres/goweb/practica/internal/domain"
)

type Service interface {
	GetAll() ([]domain.Product, error)
	Store(name string, color string, price float64, stock uint64, code string, published bool, createdAt string) (domain.Product, error)
	GetById(id uint64) (domain.Product, error)
	UpdateTotal(id uint64, name string, color string, price float64, stock uint64, code string, published bool, createdAt string) (domain.Product, error)
	UpdatePartial(id uint64, nombre string, color string, precio float64, stock uint64, codigo string, publicado bool, fechaCreacion string) (domain.Product, error)
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
	return products, errors.New("no se pudo obtener los productos")
}

func (s *service) Store(name string, color string, price float64, stock uint64, code string, published bool, createdAt string) (domain.Product, error) {
	id, err := s.repository.LastId()
	if err != nil {
		return domain.Product{}, errors.New("no se pudo cargar el último id de los productos")
	}

	id++

	producto, err := s.repository.Store(id, name, color, price, stock, code, published, createdAt)
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

func (s *service) UpdateTotal(id uint64, name string, color string, price float64, stock uint64, code string, published bool, createdAt string) (domain.Product, error) {
	producto, err := s.repository.UpdateTotal(id, name, color, price, stock, code, published, createdAt)
	if err != nil {
		return domain.Product{}, err
	}
	return producto, nil
}

func (s *service) UpdatePartial(id uint64, name string, color string, price float64, stock uint64, code string, published bool, createdAt string) (domain.Product, error) {
	producto, err := s.repository.UpdatePartial(id, name, color, price, stock, code, published, createdAt)
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
