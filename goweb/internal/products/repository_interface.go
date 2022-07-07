package products

import "goweb/internal/domain"

type Repository interface {
	GetAll() ([]domain.Product, error)
	Store(id int, nombre string, color string, precio float64, stock int, codigo string, publicado bool, fechaCreacion string) (domain.Product, error)
	GetById(id int) (domain.Product, error)
	LastId() (int, error)
}
