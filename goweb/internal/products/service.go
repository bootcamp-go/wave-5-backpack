package products

import (
	"fmt"
	"web-server/internal/domain"
)

type Service interface {
	GetAll() ([]domain.Products, error)
	Store(nombre string, color string, precio float64, stock int, codigo string, publicado bool, fecha string) (domain.Products, error)
	Update(id int, nombre string, color string, precio float64, stock int, codigo string, publicado bool, fecha string) (domain.Products, error)
	UpdateName(id int, nombre string) (domain.Products, error)
	UpdatePrice(id int, precio float64) (domain.Products, error)
	Delete(id int) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetAll() ([]domain.Products, error) {
	productsSlide, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}

	return productsSlide, nil
}

func (s *service) Store(nombre string, color string, precio float64, stock int, codigo string, publicado bool, fecha string) (domain.Products, error) {
	lastId, err := s.repository.LastID()
	if err != nil {
		return domain.Products{}, fmt.Errorf("error getting product last id: %w", err)
	}

	lastId++

	producto, err := s.repository.Store(lastId, nombre, color, precio, stock, codigo, publicado, fecha)
	if err != nil {
		return domain.Products{}, fmt.Errorf("error creating product: %w", err)

	}

	return producto, nil
}

func (s *service) Update(id int, nombre string, color string, precio float64, stock int, codigo string, publicado bool, fecha string) (domain.Products, error) {
	p, err := s.repository.Update(id, nombre, color, precio, stock, codigo, publicado, fecha)
	if err != nil {
		return domain.Products{}, fmt.Errorf("error updating product %w", err)
	}
	return p, nil
}

func (s *service) UpdateName(id int, nombre string) (domain.Products, error) {
	p, err := s.repository.UpdateName(id, nombre)
	if err != nil {
		return domain.Products{}, fmt.Errorf("Error updating name product %w", err)
	}
	return p, nil
}

func (s *service) UpdatePrice(id int, precio float64) (domain.Products, error) {
	p, err := s.repository.UpdatePrice(id, precio)
	if err != nil {
		return domain.Products{}, fmt.Errorf("Error updating price product %w", err)
	}
	return p, nil
}

func (s *service) Delete(id int) error {
	if err := s.repository.Delete(id); err != nil {
		return fmt.Errorf("error deleting product %w", err)
	}

	return nil
}
