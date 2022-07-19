package products

import "github.com/bootcamp-go/wave-5-backpack/goweb/internal/domain"

type Service interface {
	Store(nombre string, color string, precio float64, stock int, codigo string, publicado bool, fechaCreacion string) (domain.Products, error)
	Update(id int, nombre string, color string, precio float64, stock int, codigo string, publicado bool, fechaCreacion string) (domain.Products, error)
	UpdatePrecioStock(id int, precio float64, stock int) (domain.Products, error)
	GetAll() ([]domain.Products, error)
	GetByID(id int) (domain.Products, error)
	Delete(id int) (int, error)
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

type service struct {
	repository Repository
}

func (s service) Store(nombre string, color string, precio float64, stock int, codigo string, publicado bool, fechaCreacion string) (domain.Products, error) {
	return s.repository.Store(nombre, color, precio, stock, codigo, publicado, fechaCreacion)
}

func (s service) Update(id int, nombre string, color string, precio float64, stock int, codigo string, publicado bool, fechaCreacion string) (domain.Products, error) {
	return s.repository.Update(id, nombre, color, precio, stock, codigo, publicado, fechaCreacion)
}

func (s service) UpdatePrecioStock(id int, precio float64, stock int) (domain.Products, error) {
	return s.repository.UpdatePrecioStock(id, precio, stock)
}

func (s service) GetAll() ([]domain.Products, error) {
	products, err := s.repository.GetAll()

	if err != nil {
		return nil, err
	}
	return products, nil
}
func (s service) GetByID(id int) (domain.Products, error) {
	return s.repository.GetByID(id)
}

func (s service) Delete(id int) (int, error) {
	return s.repository.Delete(id)
}
