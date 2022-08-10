package products

import (
	"context"
	"goweb/internal/domain"
)

type Service interface {
	GetByName(string) (domain.Product, error)
	Store(domain.Product) (domain.Product, error)
	GetAll(context.Context) ([]domain.Product, error)
	GetOne(context.Context, int) (domain.Product, error)
	GetOneFullData(context.Context, int) (domain.ProductAndWarehouse, error)
	Update(context.Context, domain.Product) (domain.Product, error)
	Delete(context.Context, int) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetByName(name string) (domain.Product, error) {
	product, err := s.repository.GetByName(name)
	return product, err
}

func (s *service) Store(product domain.Product) (domain.Product, error) {
	product, err := s.repository.Store(product)
	return product, err
}

func (s *service) GetAll(ctx context.Context) ([]domain.Product, error) {
	return s.repository.GetAll(ctx)
}

func (s *service) GetOne(ctx context.Context, id int) (domain.Product, error) {
	return s.repository.GetOne(ctx, id)
}

func (s *service) GetOneFullData(ctx context.Context, id int) (domain.ProductAndWarehouse, error) {
	return s.repository.GetOneFullData(ctx, id)
}

func (s *service) Update(ctx context.Context, product domain.Product) (domain.Product, error) {
	return s.repository.Update(ctx, product)
}

func (s *service) Delete(ctx context.Context, id int) error {
	return s.repository.Delete(ctx, id)
}
