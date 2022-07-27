package products

import "testing-goweb/internal/domain"

type Service interface {
	GetAll() ([]domain.Product, error)
	Store(nombre, color string, precio, stock int, codigo string, publicdo bool, fechaCreacion string) (domain.Product, error)
	Update(id int, nombre, color string, precio, stock int, codigo string, publicdo bool, fechaCreacion string) (domain.Product, error)
	UpdatePrecioStock(id, precio, stock int) (domain.Product, error)
	Delete(id int) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) GetAll() ([]domain.Product, error) {
	list, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (s *service) Store(nombre, color string, precio, stock int, codigo string, publicado bool, fechaCreacion string) (domain.Product, error) {
	lastID, err := s.repository.LastID()
	if err != nil {
		return domain.Product{}, err
	}
	lastID++
	product, err := s.Store(nombre, color, precio, stock, codigo, publicado, fechaCreacion)
	if err != nil {
		return domain.Product{}, err
	}
	return product, nil
}
func (s *service) Update(id int, nombre, color string, precio, stock int, codigo string, publicado bool, fechaCreacion string) (domain.Product, error) {
	return s.repository.Update(id, nombre, color, precio, stock, codigo, publicado, fechaCreacion)
}
func (s *service) UpdatePrecioStock(id, precio, stock int) (domain.Product, error) {
	return s.repository.UpdatePrecioStock(id, precio, stock)
}
func (s *service) Delete(id int) error {
	return s.repository.Delete(id)
}
