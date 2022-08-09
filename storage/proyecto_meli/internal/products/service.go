package products

import (
	"context"
	"fmt"
	"proyecto_meli/internal/domain"
)

type Service interface {
	GetAll(ctx context.Context) ([]domain.Product, error)
	GetById(ctx context.Context, id int) (domain.Product, error)
	FilterList(ctx context.Context, id int, name, color string, price float64, stock int, codigo string, publicado bool, fecha string) ([]domain.Product, error)
	Store(ctx context.Context, name, color string, price float64, stock int, codigo string, publicado bool, fecha string) (domain.Product, error)
	Update(ctx context.Context, id int, name, color string, price float64, stock int, codigo string, publicado bool, fecha string) (domain.Product, error)
	Delete(ctx context.Context, id int) error
	Update_Name_Price(ctx context.Context, id int, name string, price float64) (domain.Product, error)
	GetProductByName(ctx context.Context, name string) ([]domain.Product, error)
}
type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetAll(ctx context.Context) ([]domain.Product, error) {
	ps, err := s.repository.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return ps, nil
}

func (s *service) GetById(ctx context.Context, id int) (domain.Product, error) {
	p, err := s.repository.GetById(ctx, id)
	if err != nil {
		return domain.Product{}, err
	}
	return p, nil
}

func (s *service) FilterList(ctx context.Context, id int, name, color string, price float64, stock int, codigo string, publicado bool, fecha string) ([]domain.Product, error) {
	ps, err := s.repository.FilterList(ctx, id, name, color, price, stock, codigo, publicado, fecha)
	if err != nil {
		return nil, err
	}
	return ps, nil
}

func (s *service) Store(ctx context.Context, name, color string, price float64, stock int, codigo string, publicado bool, fecha string) (domain.Product, error) {
	product := domain.Product{
		Nombre:        name,
		Color:         color,
		Precio:        price,
		Stock:         stock,
		Codigo:        codigo,
		Publicado:     publicado,
		FechaCreacion: fecha,
	}
	product, err := s.repository.Store(ctx, product)
	if err != nil {
		return domain.Product{}, fmt.Errorf("error creating product: %w", err)
	}
	return product, nil
}

func (s *service) GetProductByName(ctx context.Context, name string) ([]domain.Product, error) {
	products, err := s.repository.GetProductByName(ctx, name)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (s *service) Update(ctx context.Context, id int, name, color string, price float64, stock int, codigo string, publicado bool, fecha string) (domain.Product, error) {
	p := domain.Product{
		Id:            id,
		Nombre:        name,
		Color:         color,
		Precio:        price,
		Stock:         stock,
		Codigo:        codigo,
		Publicado:     publicado,
		FechaCreacion: fecha,
	}
	p, err := s.repository.Update(ctx, p)
	if err != nil {
		return domain.Product{}, fmt.Errorf("error updating product %w", err)
	}
	return p, nil
}

func (s *service) Delete(ctx context.Context, id int) error {
	if err := s.repository.Delete(ctx, id); err != nil {
		return fmt.Errorf("error deleting product %w", err)
	}
	return nil
}

func (s *service) Update_Name_Price(ctx context.Context, id int, name string, price float64) (domain.Product, error) {
	product, err := s.repository.Update_Name_Price(ctx, id, name, price)
	if err != nil {
		return domain.Product{}, fmt.Errorf("error updating product %w", err)
	}
	return product, nil
}
