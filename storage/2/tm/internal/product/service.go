package product

import (
	"context"
	"errors"
	"storage/2/tm/internal/domain"
)

type Service interface {
	Read(id int) (domain.Product, error)
	ReadByName(name string) ([]domain.Product, error)
	ReadAll() ([]domain.Product, error)
	Create(name string, pType string, price float64, count int) (domain.Product, error)
	Update(ctx context.Context, id int, name string, pType string, price float64, count int) (domain.Product, error)
	Delete(id int) error
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r: r}
}

func (s *service) Read(id int) (domain.Product, error) {
	product, err := s.r.GetOne(id)
	if err != nil {
		return domain.Product{}, err
	}

	if product.ID == 0 {
		return domain.Product{}, errors.New("no existe un producto con ese id")
	}

	return product, nil
}

func (s *service) ReadByName(name string) ([]domain.Product, error) {
	products, err := s.r.GetByName(name)
	if err != nil {
		return []domain.Product{}, err
	}

	if len(products) == 0 {
		return []domain.Product{}, errors.New("no se han encontrado productos en el listado")
	}

	return products, nil
}

func (s *service) ReadAll() ([]domain.Product, error) {
	products, err := s.r.GetAll()
	if err != nil {
		return []domain.Product{}, err
	}

	if len(products) == 0 {
		return []domain.Product{}, errors.New("no se han encontrado productos en el listado")
	}

	return products, nil
}

func (s *service) Create(name string, pType string, price float64, count int) (domain.Product, error) {
	products, err := s.r.GetAll()
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

	product, err := s.r.Store(nProduct)
	if err != nil {
		return domain.Product{}, err
	}

	return product, nil
}

func (s *service) Update(ctx context.Context, id int, name string, pType string, price float64, count int) (domain.Product, error) {
	uProduct := domain.NewProduct(id, name, pType, price, count)
	updatedProduct, err := s.r.Update(ctx, uProduct)
	if err != nil {
		return domain.Product{}, err
	}

	return updatedProduct, nil
}

func (s *service) Delete(id int) error {
	err := s.r.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
