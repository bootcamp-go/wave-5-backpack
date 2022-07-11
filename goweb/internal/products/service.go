package products

import "github.com/bootcamp-go/wave-5-backpack/goweb/internal/models"

type Service interface {
	Store(nombre, color, precio, stock, codigo, publicado, fechaCreacion) (models.Products, error)
	Update(id, nombre, color, precio, stock, codigo, publicado, fechaCreacion) (models.Products, error)
	UpdatePrecioStock(id, precio, stock) (models.Products, error)
	GetAll() ([]models.Products, error)
	GetByID(id int) (models.Products, error)
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

func (s service) Store(nombre, color, precio, stock, codigo, publicado, fechaCreacion) (models.Products, error) {
	return s.repository.Store(nombre, color, precio, stock, codigo, publicado, fechaCreacion)
}

func (s service) Update(id, nombre, color, precio, stock, codigo, publicado, fechaCreacion) (models.Products, error) {
	return s.repository.Store(id, nombre, color, precio, stock, codigo, publicado, fechaCreacion)
}

func (s service) UpdatePrecioStock(id, precio, stock) (models.Products, error) {
	return s.repository.Store(id, precio, stock)
}

func (s service) GetAll() ([]models.Products, error) {
	products, err := s.repository.GetAll()

	if err != nil {
		return nil, err
	}
	return products, nil
}
func (s service) GetByID(id int) (models.Products, error) {
	return s.repository.GetByID(id)
}

func (s service) Delete(id int) (int, error) {
	return s.repository.Delete(id)
}
