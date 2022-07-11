package products

import (
	"errors"
	"fmt"
	"goweb/internal/domain"
)

var lastId int
var products []domain.Product

type repositoryMemoria struct{}

//func NewRepositoryMemoria() Repository {
//	return &repositoryMemoria{}
//}

func (r *repositoryMemoria) GetAll() ([]domain.Product, error) {
	return products, nil
}

func (r *repositoryMemoria) Store(id int, nombre string, color string, precio float64, stock int, codigo string, publicado bool, fechaCreacion string) (domain.Product, error) {
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

func (r *repositoryMemoria) GetById(id int) (domain.Product, error) {
	for _, product := range products {
		if product.Id == id {
			return product, nil
		}
	}
	return domain.Product{}, errors.New("no se encontró el producto")
}

func (r *repositoryMemoria) LastId() (int, error) {
	return lastId, nil
}

func (r *repositoryMemoria) Update(id int, nombre string, color string, precio float64, stock int, codigo string, publicado bool, fechaCreacion string) (domain.Product, error) {
	producto := domain.Product{
		Nombre:        nombre,
		Color:         color,
		Precio:        precio,
		Stock:         stock,
		Codigo:        codigo,
		Publicado:     publicado,
		FechaCreacion: fechaCreacion,
	}

	updated := false
	for i := range products {
		if products[i].Id == id {
			producto.Id = id
			products[i] = producto
			updated = true
		}
	}

	if !updated {
		return domain.Product{}, fmt.Errorf("no se encontró el producto con el id: %d", id)
	}

	return producto, nil
}

func (r *repositoryMemoria) Delete(id int) error {
	deleted := false
	var index int
	for i := range products {
		if products[i].Id == id {
			index = i
			deleted = true
		}
	}

	if !deleted {
		return fmt.Errorf("no se encontró el producto con el id: %d", id)
	}

	products = append(products[:index], products[index+1:]...)

	return nil
}

func (r *repositoryMemoria) UpdateNombre(id int, nombre string) (domain.Product, error) {
	updated := false
	var index int
	for i := range products {
		if products[i].Id == id {
			products[i].Nombre = nombre
			index = i
			updated = true
		}
	}

	if !updated {
		return domain.Product{}, fmt.Errorf("no se encontró el producto con el id: %d", id)
	}

	return products[index], nil
}

func (r *repositoryMemoria) UpdatePrecio(id int, precio float64) (domain.Product, error) {
	updated := false
	var index int
	for i := range products {
		if products[i].Id == id {
			products[i].Precio = precio
			index = i
			updated = true
		}
	}

	if !updated {
		return domain.Product{}, fmt.Errorf("no se encontró el producto con el id: %d", id)
	}

	return products[index], nil
}
