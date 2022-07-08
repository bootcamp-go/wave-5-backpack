package products

import (
	"errors"
	"fmt"
	"goweb/internal/domain"
)

type Service interface {
	GetAll() ([]domain.Product, error)
	Store(Nombre string, Color string, Precio float64, Stock int, Codigo string, Publicado bool, FechaCreacion string) (domain.Product, error)
	GetById(id int) (domain.Product, error)
	Update(id int, Nombre string, Color string, Precio float64, Stock int, Codigo string, Publicado bool, FechaCreacion string) (domain.Product, error)
	Delete(id int) error
	UpdateNombre(id int, Nombre string) (domain.Product, error)
	UpdatePrecio(id int, Precio float64) (domain.Product, error)
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

func (s *service) Store(Nombre string, Color string, Precio float64, Stock int, Codigo string, Publicado bool, FechaCreacion string) (domain.Product, error) {
	id, err := s.repository.LastId()
	if err != nil {
		return domain.Product{}, errors.New("no se pudo cargar el último id de los productos")
	}

	id++

	producto, err := s.repository.Store(id, Nombre, Color, Precio, Stock, Codigo, Publicado, FechaCreacion)
	if err != nil {
		return domain.Product{}, errors.New("no se pudo guardar el producto")
	}

	return producto, nil
}

func (s *service) GetById(id int) (domain.Product, error) {
	producto, err := s.repository.GetById(id)
	if err != nil {
		return domain.Product{}, fmt.Errorf("no se pudo encontrar el producto con el id: %d", id)
	}
	return producto, nil
}

func (s *service) Update(id int, Nombre string, Color string, Precio float64, Stock int, Codigo string, Publicado bool, FechaCreacion string) (domain.Product, error) {
	producto, err := s.repository.Update(id, Nombre, Color, Precio, Stock, Codigo, Publicado, FechaCreacion)
	if err != nil {
		return domain.Product{}, fmt.Errorf("no se pudo actualizar el producto con el id: %d", id)
	}
	return producto, nil
}

func (s *service) Delete(id int) error {
	err := s.repository.Delete(id)
	if err != nil {
		return fmt.Errorf("no se pudo eliminar el producto con el id: %d", id)
	}
	return nil
}

func (s *service) UpdateNombre(id int, Nombre string) (domain.Product, error) {
	producto, err := s.repository.UpdateNombre(id, Nombre)
	if err != nil {
		return domain.Product{}, fmt.Errorf("no se pudo actualizar el nombre del producto con el id: %d", id)
	}
	return producto, nil
}

func (s *service) UpdatePrecio(id int, Precio float64) (domain.Product, error) {
	producto, err := s.repository.UpdatePrecio(id, Precio)
	if err != nil {
		return domain.Product{}, fmt.Errorf("no se pudo actualizar el precio del producto con el id: %d", id)
	}
	return producto, nil
}
