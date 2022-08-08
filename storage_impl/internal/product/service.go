package product

import (
	"context"
	"errors"
	"storage/impl/internal/domain"
)

var (
	ErrDB       = errors.New("dataBase error")
	ErrNotFound = errors.New("product not found")
)

type Service interface {
	GetAll(ctx context.Context) ([]domain.Product, error)
	Get(ctx context.Context, id int) (domain.Product, error)
	Save(ctx context.Context, nombre, color string, precio, stock int, codigo string, publicado bool, fechaCreacion string) (domain.Product, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{repository: r}
}

func (s *service) GetAll(ctx context.Context) ([]domain.Product, error) {
	products, err := s.repository.GetAll(ctx)
	if err != nil {
		return []domain.Product{}, ErrDB
	}
	return products, nil
}

func (s *service) Get(ctx context.Context, id int) (domain.Product, error) {
	product, err := s.repository.Get(ctx, id)
	if err != nil {
		return domain.Product{}, err
	}
	return product, nil
}

func (s *service) Save(ctx context.Context, nombre, color string, precio, stock int, codigo string, publicado bool, fechaCreacion string) (domain.Product, error) {
	product := domain.Product{
		Nombre:        nombre,
		Color:         color,
		Precio:        precio,
		Stock:         stock,
		Codigo:        codigo,
		Publicado:     publicado,
		FechaCreacion: fechaCreacion,
	}

	id, err := s.repository.Save(ctx, product)
	if err != nil {
		return domain.Product{}, ErrDB
	}
	productDB, err := s.Get(ctx, id)
	if err != nil {
		return domain.Product{}, ErrDB
	}
	return productDB, nil
}
