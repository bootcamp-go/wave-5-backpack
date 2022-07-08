package products

import (
	"fmt"
	"goweb/internal/domain"
)

var ps []domain.Products
var lastID int

type Repository interface {
	GetAll() ([]domain.Products, error)
	CreateProduct(id int, nombre, color string, precio float64, stock int, código string, publicado bool, fecha_de_creación string) (domain.Products, error)
	LastID() (int, error)
	Update(id int, nombre, color string, precio float64, stock int, código string, publicado bool, fecha_de_creación string) (domain.Products, error)
	Delete(id int) error
	UpdateOne(id int, nombre string, precio float64) (domain.Products, error)
}

type repository struct{}

func InitRepository() Repository {
	return &repository{}
}

func (r *repository) GetAll() ([]domain.Products, error) {
	return ps, nil
}

func (r *repository) LastID() (int, error) {
	return lastID, nil
}

func (r *repository) CreateProduct(id int, nombre, color string, precio float64, stock int, código string, publicado bool, fecha_de_creación string) (domain.Products, error) {
	p := domain.Products{
		Id:            id,
		Nombre:        nombre,
		Color:         color,
		Precio:        precio,
		Stock:         stock,
		Codigo:        código,
		Publicado:     publicado,
		FechaCreacion: fecha_de_creación,
	}
	ps = append(ps, p)
	lastID = p.Id
	return p, nil
}

func (r *repository) Update(id int, nombre, color string, precio float64, stock int, código string, publicado bool, fecha_de_creación string) (domain.Products, error) {
	p := domain.Products{
		Id:            id,
		Nombre:        nombre,
		Color:         color,
		Precio:        precio,
		Stock:         stock,
		Codigo:        código,
		Publicado:     publicado,
		FechaCreacion: fecha_de_creación,
	}
	updated := false
	for i := range ps {
		if ps[i].Id == id {
			p.Id = id
			ps[i] = p
			updated = true
		}
	}
	if !updated {
		return domain.Products{}, fmt.Errorf("Producto %d no encontrado", id)
	}
	return p, nil
}

func (r *repository) Delete(id int) error {
	deleted := false
	var index int
	for i := range ps {
		if ps[i].Id == id {
			index = i
			deleted = true
		}
	}
	if !deleted {
		return fmt.Errorf("Producto %d no encontrado", id)
	}
	ps = append(ps[:index], ps[index+1:]...)
	return nil

}

func (r *repository) UpdateOne(id int, nombre string, precio float64) (domain.Products, error) {
	var p domain.Products
	updated := false
	for i := range ps {
		if ps[i].Id == id {
			ps[i].Nombre = nombre
			ps[i].Precio = precio
			updated = true
			p = ps[i]
		}
	}
	if !updated {
		return domain.Products{}, fmt.Errorf("Producto %d no encontrado", id)
	}
	return p, nil
}
