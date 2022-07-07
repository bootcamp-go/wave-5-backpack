package products

import (
	"errors"
	"goweb/internal/domain"
)

var lastId int
var products []domain.Product

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetAll() ([]domain.Product, error) {
	return products, nil
}

func (r *repository) Store(id int, nombre string, color string, precio float64, stock int, codigo string, publicado bool, fechaCreacion string) (domain.Product, error) {
	producto := domain.Product{
		Id:            id,
		Nombre:        nombre,
		Color:         color,
		Precio:        precio,
		Stock:         stock,
		Codigo:        codigo,
		Publicado:     publicado,
		FechaCreacion: fechaCreacion,
	}

	products = append(products, producto)
	lastId = id

	return producto, nil
}

func (r *repository) GetById(id int) (domain.Product, error) {
	for _, product := range products {
		if product.Id == id {
			return product, nil
		}
	}
	return domain.Product{}, errors.New("no se encontró el producto")
}

func (r *repository) LastId() (int, error) {
	return lastId, nil
}
