package products

import (
	"Clase1_Parte1/productos/internal/domain"
)

type Service interface {
	GetAll() ([]domain.Product, error)
	Store(nombre, color string, precio, stock int, codigo string, publicado bool, fechaCreacion string) (domain.Product, error)
	GetByID(id int) (domain.Product, error)
	GetByName(nombre string) ([]domain.Product, error)
	Update(id int, nombre, color string, precio, stock int, codigo string, publicado bool, fechaCreacion string) (domain.Product, error)
	UpdateNamePrice(id int, nombre string, precio int) (domain.Product, error)
	Delete(id int) (domain.Product, error)
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
	return s.repository.GetAll()
}

func (s *service) GetByID(id int) (domain.Product, error) {
	return s.repository.GetByID(id)
}

func (s *service) GetByName(nombre string) ([]domain.Product, error) {
	return s.repository.GetByName(nombre)
}

func (s *service) Store(nombre, color string, precio, stock int, codigo string, publicado bool, fechaCreacion string) (domain.Product, error) {
	return s.repository.Store(nombre, color, precio, stock, codigo, publicado, fechaCreacion)
}

func (s *service) Update(id int, nombre, color string, precio, stock int, codigo string, publicado bool, fechaCreacion string) (domain.Product, error) {
	return s.repository.Update(id, nombre, color, precio, stock, codigo, publicado, fechaCreacion)
}

func (s *service) UpdateNamePrice(id int, nombre string, precio int) (domain.Product, error) {
	return s.repository.UpdateNamePrice(id, nombre, precio)
}

func (s *service) Delete(id int) (domain.Product, error) {
	return s.repository.Delete(id)
}
