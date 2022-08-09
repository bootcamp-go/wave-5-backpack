package products

import (
	"context"
	"errors"
	"fmt"

	"github.com/bootcamp-go/wave-5-backpack/tree/Ramos_Andres/goweb/practica/internal/domain"
)

type Service interface {
	Store(ctx context.Context, product domain.Product) (domain.Product, error)
	GetAll(ctx context.Context) ([]domain.Product, error)
	GetById(ctx context.Context, id uint64) (domain.Product, error)
	UpdateTotal(ctx context.Context, product domain.Product) (domain.Product, error)
	UpdatePartial(ctx context.Context, product domain.Product) (domain.Product, error)
	Delete(ctx context.Context, id uint64) error
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
	products, err := s.repository.GetAll(ctx)
	if err != nil {
		return nil, errors.New("no se pudo obtener los productos")
	}
	return products, nil
}

func (s *service) Store(ctx context.Context, product domain.Product) (domain.Product, error) {
	producto, err := s.repository.Store(ctx, product)
	if err != nil {
		return domain.Product{}, errors.New("no se pudo guardar el producto")
	}

	return producto, nil
}

func (s *service) GetById(ctx context.Context, id uint64) (domain.Product, error) {
	producto, err := s.repository.GetById(ctx, id)
	if err != nil {
		return domain.Product{}, fmt.Errorf("no se pudo encontrar el producto con el id: %d", id)
	}
	return producto, nil
}

func (s *service) UpdateTotal(ctx context.Context, product domain.Product) (domain.Product, error) {
	producto, err := s.repository.Update(ctx, product)
	if err != nil {
		return domain.Product{}, err
	}
	return producto, nil
}

func (s *service) UpdatePartial(ctx context.Context, product domain.Product) (domain.Product, error) {
	producto, err := s.repository.Update(ctx, product)
	if err != nil {
		return domain.Product{}, err
	}
	return producto, nil
}

func (s *service) Delete(ctx context.Context, id uint64) error {
	err := s.repository.Delete(ctx, id)
	if err != nil {
		return fmt.Errorf("no se pudo encontrar el producto con el id: %d", id)
	}
	return nil
}
