package products

import "github.com/bootcamp-go/storage/internal/domains"

type Service interface {
	Store(domains.Product) (int, error)
	GetByName(name string) (domains.Product, error)
	GetAll()([]domains.Product, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{repository: r}
}

func (s *service) Store(p domains.Product) (int, error) {
	return s.repository.Store(p)
}

func (s *service) GetByName(name string) (domains.Product, error) {
	return s.repository.GetByName(name)
}
  
func (s *service) GetAll()  ([]domains.Product, error ) {
	return s.repository.GetAll()

}

