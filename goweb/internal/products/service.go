package products

import (
	"fmt"
	"goweb/internal/domain"
)

type Service interface {
	GetAll() ([]domain.Products, error)
	CreateProduct(id int, nombre, color string, precio float64, stock int, código string, publicado bool, fecha_de_creación string) (domain.Products, error)
	Update(id int, nombre, color string, precio float64, stock int, código string, publicado bool, fecha_de_creación string) (domain.Products, error)
	Delete(id int) error
	UpdateOne(id int, nombre string, precio float64) (domain.Products, error)
}

type service struct {
	repository Repository
}

func InitService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetAll() ([]domain.Products, error) {

	ps, err := s.repository.GetAll()

	if err != nil {
		return nil, err
	}
	return ps, nil
}

func (s *service) CreateProduct(id int, nombre, color string, precio float64, stock int, código string, publicado bool, fecha_de_creación string) (domain.Products, error) {
	lastID, err := s.repository.LastID()
	if err != nil {
		return domain.Products{}, fmt.Errorf("Error obteniendo el último Id: %w", err)
	}
	lastID++
	producto, err := s.repository.CreateProduct(lastID, nombre, color, precio, stock, código, publicado, fecha_de_creación)
	if err != nil {
		return domain.Products{}, fmt.Errorf("Error creando un producto: %w", err)
	}

	return producto, nil
}

func (s *service) Update(id int, nombre, color string, precio float64, stock int, código string, publicado bool, fecha_de_creación string) (domain.Products, error) {
	return s.repository.Update(id, nombre, color, precio, stock, código, publicado, fecha_de_creación)
}

func (s *service) Delete(id int) error {
	return s.repository.Delete(id)
}

func (s *service) UpdateOne(id int, nombre string, precio float64) (domain.Products, error) {
	return s.repository.UpdateOne(id, nombre, precio)
}
