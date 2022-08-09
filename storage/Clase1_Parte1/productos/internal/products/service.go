package products

import (
	"Clase1_Parte1/productos/internal/domain"
	"context"
)

type Service interface {
	GetAll(ctx context.Context) ([]domain.Product, error)
	Store(ctx context.Context, nombre, color string, precio, stock int, codigo string, publicado bool, fechaCreacion string) (domain.Product, error)
	GetByID(ctx context.Context, id int) (domain.Product, error)
	GetByName(ctx context.Context, nombre string) ([]domain.Product, error)
	Update(ctx context.Context, id int, nombre, color string, precio, stock int, codigo string, publicado bool, fechaCreacion string) (domain.Product, error)
	UpdateNamePrice(ctx context.Context, id int, nombre string, precio int) (domain.Product, error)
	Delete(ctx context.Context, id int) (domain.Product, error)
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

func (s *service) GetByID(ctx context.Context, id int) (domain.Product, error) {
	return s.repository.GetByID(ctx, id)
}

func (s *service) GetByName(ctx context.Context, nombre string) ([]domain.Product, error) {
	return s.repository.GetByName(ctx, nombre)
}

func (s *service) Store(ctx context.Context, nombre, color string, precio, stock int, codigo string, publicado bool, fechaCreacion string) (domain.Product, error) {
	return s.repository.Store(ctx, nombre, color, precio, stock, codigo, publicado, fechaCreacion)
}

func (s *service) Update(ctx context.Context, id int, nombre, color string, precio, stock int, codigo string, publicado bool, fechaCreacion string) (domain.Product, error) {
	return s.repository.Update(ctx, id, nombre, color, precio, stock, codigo, publicado, fechaCreacion)
}

func (s *service) UpdateNamePrice(ctx context.Context, id int, nombre string, precio int) (domain.Product, error) {
	return s.repository.UpdateNamePrice(ctx, id, nombre, precio)
}

func (s *service) Delete(ctx context.Context, id int) (domain.Product, error) {
	return s.repository.Delete(ctx, id)
}
