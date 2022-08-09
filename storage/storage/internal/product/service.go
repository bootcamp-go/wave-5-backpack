package product

import (
	"context"
	"storage/internal/domain"
)

type Service interface {
	GetAll(ctx context.Context) ([]domain.Product, error)
	GetByName(ctx context.Context, name string) (domain.Product, error)
	Store(ctx context.Context, product domain.Product) (domain.Product, error)
	Update(ctx context.Context, product domain.Product) (domain.Product, error)
	Delete(ctx context.Context, id int) error
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

func (s *service) GetByName(ctx context.Context, name string) (domain.Product, error) {
	return s.repository.GetByName(ctx, name)
}

func (s *service) Store(ctx context.Context, product domain.Product) (domain.Product, error) {
	return s.repository.Store(ctx, product)
}

func (s *service) Update(ctx context.Context, product domain.Product) (domain.Product, error) {
	return s.repository.Update(ctx, product)
}

func (s *service) Delete(ctx context.Context, id int) error {
	return s.repository.Delete(ctx, id)
}
