package product

import (
	"context"
	"errors"
	"storage/3/tm/internal/domain"
)

type Service interface {
	Read(ctx context.Context, id string) (domain.Product, error)
	ReadAll(ctx context.Context) ([]domain.Product, error)
	Create(ctx context.Context, name string, pType string, price float64, count int) (domain.Product, error)
	Update(ctx context.Context, id string, name string, pType string, price float64, count int) (domain.Product, error)
	Delete(ctx context.Context, id string) error
}

type service struct {
	r RepositoryDynamo
}

func NewService(r RepositoryDynamo) Service {
	return &service{r: r}
}

func (s *service) Read(ctx context.Context, id string) (domain.Product, error) {
	product, err := s.r.GetOne(ctx, id)
	if err != nil {
		return domain.Product{}, err
	}

	if product.ID == "" {
		return domain.Product{}, errors.New("no existe un producto con ese id")
	}

	return product, nil
}

func (s *service) ReadAll(ctx context.Context) ([]domain.Product, error) {
	products, err := s.r.GetAll(ctx)
	if err != nil {
		return []domain.Product{}, err
	}

	if len(products) == 0 {
		return []domain.Product{}, errors.New("no se han encontrado productos en el listado")
	}

	return products, nil
}

func (s *service) Create(ctx context.Context, name string, pType string, price float64, count int) (domain.Product, error) {
	products, err := s.r.GetAll(ctx)
	if err != nil {
		return domain.Product{}, err
	}

	if len(products) != 0 {
		for _, p := range products {
			if p.Name == name && p.Type == pType && p.Price == price && p.Count == count {
				return domain.Product{}, errors.New("el producto ingresado ya existe")
			}
		}
	}

	nProduct := domain.Product{
		Name:  name,
		Type:  pType,
		Count: count,
		Price: price,
	}

	product, err := s.r.Store(ctx, nProduct)
	if err != nil {
		return domain.Product{}, err
	}

	return product, nil
}

func (s *service) Update(ctx context.Context, id string, name string, pType string, price float64, count int) (domain.Product, error) {
	uProduct := domain.NewProduct(id, name, pType, price, count)
	updatedProduct, err := s.r.Update(ctx, uProduct)
	if err != nil {
		return domain.Product{}, err
	}

	return updatedProduct, nil
}

func (s *service) Delete(ctx context.Context, id string) error {
	err := s.r.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
