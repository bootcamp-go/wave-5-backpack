package products

import (
	"fmt"
	"proyecto_meli/internal/domain"
)

type Service interface {
	GetAll() ([]domain.Product, error)
	GetById(id int) (domain.Product, error)
	FilterList(id int, name, color string, price float64, stock int, codigo string, publicado bool, fecha string) ([]domain.Product, error)
	Store(name, color string, price float64, stock int, codigo string, publicado bool, fecha string) (domain.Product, error)
	Update(id int, name, color string, price float64, stock int, codigo string, publicado bool, fecha string) (domain.Product, error)
	Delete(id int) error
	Update_Name_Price(id int, name string, price float64) (domain.Product, error)
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
	ps, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return ps, nil
}

func (s *service) GetById(id int) (domain.Product, error) {
	p, err := s.repository.GetById(id)
	if err != nil {
		return domain.Product{}, err
	}
	return p, nil
}

func (s *service) FilterList(id int, name, color string, price float64, stock int, codigo string, publicado bool, fecha string) ([]domain.Product, error) {
	ps, err := s.repository.FilterList(id, name, color, price, stock, codigo, publicado, fecha)
	if err != nil {
		return nil, err
	}
	return ps, nil
}

func (s *service) Store(name, color string, price float64, stock int, codigo string, publicado bool, fecha string) (domain.Product, error) {
	lastID, err := s.repository.LastID()
	if err != nil {
		return domain.Product{}, fmt.Errorf("error getting product last id: %w", err)
	}
	lastID++
	product, err := s.repository.Store(lastID, name, color, price, stock, codigo, publicado, fecha)
	if err != nil {
		return domain.Product{}, fmt.Errorf("error creating product: %w", err)
	}
	return product, nil
}

func (s *service) Update(id int, name, color string, price float64, stock int, codigo string, publicado bool, fecha string) (domain.Product, error) {
	p, err := s.repository.Update(id, name, color, price, stock, codigo, publicado, fecha)
	if err != nil {
		return domain.Product{}, fmt.Errorf("error updating product %w", err)
	}
	return p, nil
}

func (s *service) Delete(id int) error {
	if err := s.repository.Delete(id); err != nil {
		return fmt.Errorf("error deleting product %w", err)
	}
	return nil
}

func (s *service) Update_Name_Price(id int, name string, price float64) (domain.Product, error) {
	product, err := s.repository.Update_Name_Price(id, name, price)
	if err != nil {
		return domain.Product{}, fmt.Errorf("error updating product %w", err)
	}
	return product, nil
}
