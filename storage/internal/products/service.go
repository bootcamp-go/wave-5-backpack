package products

import (
	"fmt"
	"storage/internal/domain"
)

type Service interface {
	GetById(id int) (domain.Products, error)
	CreateProduct(product domain.Products) (domain.Products, error)
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

func (s *service) GetById(id int) (domain.Products, error) {
	getProduct, err := s.repository.GetById(id)
	fmt.Println(getProduct)
	fmt.Println(err)

	if err != nil {
		return domain.Products{}, fmt.Errorf("no se ha encontrado el producto")
	}
	return getProduct, nil
}

func (s *service) CreateProduct(product domain.Products) (domain.Products, error) {
	producto, err := s.repository.CreateProduct(product)
	if err != nil {
		return domain.Products{}, fmt.Errorf("error creando un producto: %w", err)
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
