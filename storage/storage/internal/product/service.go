package product

import "storage/internal/domain"

type Service interface {
	GetByName(name string) (domain.Product, error)
	Store(product domain.Product) (domain.Product, error)
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
	return s.repository.GetByName(name)
}

func (s *service) Store(product domain.Product) (domain.Product, error) {
	return s.repository.Store(product)
}
