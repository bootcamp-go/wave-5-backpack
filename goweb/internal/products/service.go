package products

import "goweb/internal/domain"

type Service interface {
	GetAll() ([]domain.Product, error)
	Create(name string, color string, price float64, stock int, code string, publisher bool) (domain.Product, error)
	Update(id int, name string, color string, price float64, stock int, code string, publisher bool) (domain.Product, error)
	Delete(id int) error
	ParcialUpdate(id int, name string, price float64) (domain.Product, error)
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

// ParcialUpdate implements Service
func (s *service) ParcialUpdate(id int, name string, price float64) (domain.Product, error) {
	return s.repository.ParcialUpdate(id, name, price)
}

// Delete implements Service
func (s *service) Delete(id int) error {
	return s.repository.Delete(id)
}

// Update implements Service
func (s *service) Update(id int, name string, color string, price float64, stock int, code string, publisher bool) (domain.Product, error) {
	return s.repository.Update(id, name, color, price, stock, code, publisher)
}

// Create implements Service
func (s *service) Create(name string, color string, price float64, stock int, code string, publisher bool) (domain.Product, error) {
	return s.repository.Create(name, color, price, stock, code, publisher)
}


// GetAll implements Service
func (s *service) GetAll() ([]domain.Product, error) {
	ps, err := s.repository.GetAll()
	if err !=  nil {
		return nil, err
	}
	return ps, nil
}
