package productos

import (
	"fmt"

	"github.com/bootcamp-go/wave-5-backpack/tree/ospina_christian/goweb/clase3_parte2/internal/domain"
)

type Service interface {
	Store(nombre string, color string, precio float64, stock int, codigo string, publicado bool, fechaCreacion string) (domain.Productos, error)
	GetAll() ([]domain.Productos, error)
	Update(id int, nombre string, color string, precio float64, stock int, codigo string, publicado bool, fechaCreacion string) (domain.Productos, error)
	UpdatePrecio(id int, precio float64) (domain.Productos, error)
	Delete(id int) error
	GetForId(id int) (domain.Productos, error)
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{repo: r}
}

func (s *service) Update(id int, nombre string, color string, precio float64, stock int, codigo string, publicado bool, fechaCreacion string) (domain.Productos, error) {
	p, err := s.repo.Update(id, nombre, color, precio, stock, codigo, publicado, fechaCreacion)
	if err != nil {
		return domain.Productos{}, fmt.Errorf("error al actualizar el producto: %w", err)
	}
	return p, nil
}

func (s *service) Store(nombre string, color string, precio float64, stock int, codigo string, publicado bool, fechaCreacion string) (domain.Productos, error) {
	lastID, err := s.repo.LastID()
	if err != nil {
		return domain.Productos{}, fmt.Errorf("error obteniendo al obtener la ultima id: %w", err)
	}
	lastID++
	product, err := s.repo.Store(lastID, nombre, color, precio, stock, codigo, publicado, fechaCreacion)
	if err != nil {
		return domain.Productos{}, fmt.Errorf("error al crear el producto: %w", err)
	}
	return product, nil

}

func (s *service) GetAll() ([]domain.Productos, error) {
	products, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (s *service) GetForId(id int) (domain.Productos, error) {
	product, err := s.repo.GetForId(id)
	if err != nil {
		return domain.Productos{}, fmt.Errorf("error al obtener el producto %w", err)
	}
	return product, nil
}

func (s *service) UpdatePrecio(id int, precio float64) (domain.Productos, error) {
	product, err := s.repo.UpdatePrecio(id, precio)
	if err != nil {
		return domain.Productos{}, fmt.Errorf("error updating product %w", err)
	}
	return product, nil
}

func (s *service) Delete(id int) error {
	if err := s.repo.Delete(id); err != nil {
		return fmt.Errorf("error deleting el producto %w", err)
	}
	return nil
}
