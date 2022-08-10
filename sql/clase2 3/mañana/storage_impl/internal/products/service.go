package products

import (
	"context"
	"fmt"

	"github.com/bootcamp-go/storage/internal/domains"
)

type Service interface {
	Get(context.Context, int) (domains.Product, error)
	GetAll(context.Context, int) ([]domains.Products, error)
	Store(context.Context, domains.Product) (int, error)
	GetByName(context.Context, string) (domains.Product, error)
	Update(context.Context, domains.Product) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{repository: r}
}

func (s *service) Get(ctx context.Context, id int) (domains.Product, error) {
	return s.repository.Get(ctx, id)
}

func (s *service) GetAll(ctx context.Context, id int) ([]domains.Products, error) {
	return s.repository.GetAll(ctx, id)
}

func (s *service) Store(ctx context.Context, p domains.Product) (int, error) {
	return s.repository.Store(ctx, p)
}

func (s *service) GetByName(ctx context.Context, name string) (domains.Product, error) {
	return s.repository.GetByName(ctx, name)
}

func (s *service) Update(ctx context.Context, p domains.Product) error {
	if !s.repository.Exists(ctx, p.ID) {
		return fmt.Errorf("not exists product id %v", p.ID)
	}

	product, err := s.repository.Get(ctx, p.ID)
	if err != nil {
		return err
	}

	if p.Name == "" {
		p.Name = product.Name
	}

	if p.Count == 0 {
		p.Count = product.Count
	}

	if p.Type == "" {
		p.Type = product.Type
	}

	if p.Price == 0 {
		p.Price = product.Price
	}

	if p.WarehouseId == 0 {
		p.WarehouseId = product.WarehouseId
	}

	return s.repository.Update(ctx, p)
}
