package products

import "storage/internal/domain"

type Service interface {
	Store(domain.Product) (int, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{repository: r}
}

func (s *service) Store(p domain.Product) (int, error) {
	return s.repository.Store(p)
}
