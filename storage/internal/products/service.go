package products

import (
	"context"
	"fmt"
	"storage/internal/domain"
)

type Service interface {
	GetAll() ([]domain.Products, error)
	GetById(id int) (domain.Products, error)
	CreateProduct(product domain.Products) (domain.Products, error)
	Update(ctx context.Context, p domain.Products, id int) (domain.Products, error)
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
	products, err := s.repository.GetAll()
	fmt.Println(products)
	if err != nil {
		return []domain.Products{}, fmt.Errorf("no hay datos")
	}
	return products, nil
}

func (s *service) GetById(id int) (domain.Products, error) {
	getProduct, err := s.repository.GetById(id)

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

func (s *service) Update(ctx context.Context, p domain.Products, id int) (domain.Products, error) {
	if id <= 0 {
		return domain.Products{}, fmt.Errorf("el id no puede ser 0")
	}
	productData, err := s.GetById(id)
	if err != nil {
		return domain.Products{}, err
	}
	if p.Nombre != "" {
		productData.Nombre = p.Nombre
	}
	if p.Color != "" {
		productData.Color = p.Color
	}
	if p.Precio > 0 {
		productData.Precio = p.Precio
	}
	if p.Stock >= 0 {
		productData.Stock = p.Stock
	}

	if err := s.repository.Update(ctx, productData); err != nil {
		return domain.Products{}, err
	}

	return productData, nil
}

func (s *service) Delete(id int) error {
	return s.repository.Delete(id)
}

func (s *service) UpdateOne(id int, nombre string, precio float64) (domain.Products, error) {
	return s.repository.UpdateOne(id, nombre, precio)
}
