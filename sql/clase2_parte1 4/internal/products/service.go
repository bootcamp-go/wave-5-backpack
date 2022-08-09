package products

import (
	"clase2_parte1/internal/domain"
	"context"
)

type Service interface {
	Store(p domain.Product) (domain.Product, error)
	GetOne(id int) (domain.Product, error)
	Update(product domain.Product) (domain.Product, error)
	GetAll() ([]domain.Product, error)
	GetFullData(id int) ([]domain.ProductAndWarehouse, error)
	Delete(id int) error
	GetOneWithcontext(ctx context.Context, id int) (domain.Product, error)
}

type service struct {
	product Repository
}

func NewService(product Repository) Service {
	return &service{
		product: product,
	}
}

func (s *service) Store(p domain.Product) (domain.Product, error) {
	product, err := s.product.Store(p)
	if err != nil {
		return domain.Product{}, err
	}

	p.Id = product.Id
	return p, nil
}

func (s *service) GetOne(id int) (domain.Product, error) {
	product, err := s.product.GetOne(id)
	if err != nil {
		return domain.Product{}, err
	}

	return product, err
}

func (s *service) Update(product domain.Product) (domain.Product, error) {
	product, err := s.product.Update(product)
	if err != nil {
		return domain.Product{}, err
	}
	return product, err
}

func (s *service) GetAll() ([]domain.Product, error) {
	product, err := s.product.GetAll()
	if err != nil {
		return []domain.Product{}, err
	}

	return product, err
}

func (s *service) Delete(id int) error {
	err := s.product.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) GetFullData(id int) ([]domain.ProductAndWarehouse, error) {
	productsAndWarehouses, err := s.product.GetFullData(id)
	if err != nil {
		return []domain.ProductAndWarehouse{}, err
	}

	return productsAndWarehouses, err
}

func (s *service) GetOneWithcontext(ctx context.Context, id int) (domain.Product, error) {
	product, err := s.product.GetOneWithcontext(ctx, id)
	if err != nil {
		return domain.Product{}, err
	}

	return product, err
}