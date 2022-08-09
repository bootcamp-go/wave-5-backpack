package products

import (
	"context"
	"products_project/internal/domain"
)

type Service interface {
	GetAll(ctx context.Context) ([]domain.Product, error)
	Store(ctx context.Context, nombre, color string, precio, stock int, codigo string, publicado bool, fecha string) (domain.Product, error)
	Update(ctx context.Context, id int, nombre, color string, precio, stock int, codigo string, publicado bool, fecha string) (domain.Product, error)
	Delete(ctx context.Context, id int) (domain.Product, error)
	UpdateFields(ctx context.Context, id int, nombre string, precio int) (domain.Product, error)
	GetById(ctx context.Context, id int) (domain.Product, error)
	GetByName(ctx context.Context, nombre string) ([]domain.Product, error)
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
	return s.repository.GetAll(ctx)
}

func (s *service) Store(ctx context.Context, nombre, color string, precio, stock int, codigo string, publicado bool, fecha string) (domain.Product, error) {
	return s.repository.Store(ctx, nombre, color, precio, stock, codigo, publicado, fecha)
}

func (s *service) Update(ctx context.Context, id int, nombre, color string, precio, stock int, codigo string, publicado bool, fecha string) (domain.Product, error) {
	return s.repository.Update(ctx, id, nombre, color, precio, stock, codigo, publicado, fecha)
}

func (s *service) Delete(ctx context.Context, id int) (domain.Product, error) {
	return s.repository.Delete(ctx, id)
}

func (s *service) UpdateFields(ctx context.Context, id int, nombre string, precio int) (domain.Product, error) {
	return s.repository.UpdateFields(ctx, id, nombre, precio)
}

func (s *service) GetById(ctx context.Context, id int) (domain.Product, error) {
	return s.repository.GetById(ctx, id)
}

func (s *service) GetByName(ctx context.Context, nombre string) ([]domain.Product, error) {
	return s.repository.GetByName(ctx, nombre)
}
