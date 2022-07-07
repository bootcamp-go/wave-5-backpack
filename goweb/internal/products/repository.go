package products

import (
	"errors"
	"goweb/internal/domain"
)

type Repository interface {
	GetAll() ([]domain.Product, error)
	Store(Id int, Nombre string, Color string, Precio float64, Stock int, Codigo string, Publicado bool, FechaCreacion string) (domain.Product, error)
	GetById(id int) (domain.Product, error)
	LastId() (int, error)
}

var lastId int
var products []domain.Product

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetAll() ([]domain.Product, error) {
	return products, nil
}

func (r *repository) Store(Id int, Nombre string, Color string, Precio float64, Stock int, Codigo string, Publicado bool, FechaCreacion string) (domain.Product, error) {
	producto := domain.Product{Id, Nombre, Color, Precio, Stock, Codigo, Publicado, FechaCreacion}

	products = append(products, producto)
	lastId = Id

	return producto, nil
}

func (r *repository) GetById(id int) (domain.Product, error) {
	for _, product := range products {
		if product.Id == id {
			return product, nil
		}
	}
	return domain.Product{}, errors.New("No se encontró el producto")
}

func (r *repository) LastId() (int, error) {
	return lastId, nil
}
