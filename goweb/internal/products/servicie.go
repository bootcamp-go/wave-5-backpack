package products

import (
	"github.com/bootcamp-go/wave-5-backpack/goweb/internal/domain"
)

type Service interface {
	GetAll() ([]domain.Product, error)
	GetProduct(id int) (domain.Product, error)
	Store(nombre string, color string, precio float64, stock int, codigo string, publicado bool) (domain.Product, error)
	UpdateAll(id int, nombre string, color string, precio float64, stock int, codigo string, publicado bool) (domain.Product, error)
	Delete(id int) error
	Update(id int, nombre string, precio float64) (domain.Product, error)
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

func (s service) GetProduct(id int) (domain.Product, error) {
	product, err := s.repository.GetProduct(id)

	if err != nil {
		return domain.Product{}, err
	}

	return product, nil
}

func (s service) Store(nombre string, color string, precio float64, stock int, codigo string, publicado bool) (domain.Product, error) {
	newID, err := s.repository.LastID()
	if err != nil {
		return domain.Product{}, err
	}

	newProduct, err := s.repository.Store(newID, nombre, color, precio, stock, codigo, publicado)
	if err != nil {
		return domain.Product{}, err
	}

	return newProduct, nil
}

func (s service) UpdateAll(id int, nombre string, color string, precio float64, stock int, codigo string, publicado bool) (domain.Product, error) {
	product, err := s.repository.UpdateAll(id, nombre, color, precio, stock, codigo, publicado)

	if err != nil {
		return product, err
	}

	return product, nil

}

func (s service) Delete(id int) error {
	if err := s.repository.Delete(id); err != nil {
		return err
	}

	return nil
}

func (s service) Update(id int, nombre string, precio float64) (domain.Product, error) {
	product, err := s.repository.Update(id, nombre, precio)

	if err != nil {
		return product, err
	}

	return product, nil
}
