package products

import (
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
		return domain.Products{}, err
	}

	lastId++

	producto, err := s.repository.Store(lastId, nombre, color, precio, stock, codigo, publicado, fecha)
	if err != nil {
		return domain.Products{}, err
	}

	return producto, nil
}

func (s *service) Update(id int, nombre string, color string, precio float64, stock int, codigo string, publicado bool, fecha string) (domain.Products, error) {
	return s.repository.Update(id, nombre, color, precio, stock, codigo, publicado, fecha)
}

func (s *service) UpdateName(id int, nombre string) (domain.Products, error) {
	return s.repository.UpdateName(id, nombre)
}

func (s *service) UpdatePrice(id int, precio float64) (domain.Products, error) {
	return s.repository.UpdatePrice(id, precio)
}

func (s *service) Delete(id int) error {
	return s.repository.Delete(id)
}
