package productos

import (
	"fmt"
)

type Productos struct {
	Id            int     `json:"id"`
	Nombre        string  `json:"nombre"`
	Color         string  `json:"color"`
	Precio        float64 `json:"precio"`
	Stock         int     `json:"stock"`
	Codigo        string  `json:"codigo"`
	Publicado     bool    `json:"publicado"`
	FechaCreación string  `json:"fecha_creacion"`
}

var products []Productos
var lastID int

type Repository interface {
	LastID() (int, error)
	Store(id int, nombre string, color string, precio float64, stock int, codigo string, publicado bool, fechaCreacion string) (Productos, error)
	GetAll() ([]Productos, error)
	Update(id int, nombre string, color string, precio float64, stock int, codigo string, publicado bool, fechaCreacion string) (Productos, error)
	UpdatePrecio(id int, precion float64) (Productos, error)
	Delete(id int) error
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) Store(id int, nombre string, color string, precio float64, stock int, codigo string, publicado bool, fechaCreacion string) (Productos, error) {
	p := Productos{id, nombre, color, precio, stock, codigo, publicado, fechaCreacion}
	lastID = p.Id
	products = append(products, p)
	return p, nil
}

func (r *repository) Update(id int, nombre string, color string, precio float64, stock int, codigo string, publicado bool, fechaCreacion string) (Productos, error) {
	p := Productos{Nombre: nombre, Color: color, Precio: precio, Stock: stock, Codigo: codigo, Publicado: publicado, FechaCreación: fechaCreacion}
	updated := false
	for i := range products {
		if products[i].Id == id {
			p.Id = id
			products[i] = p
			updated = true
		}
	}
	if !updated {
		return Productos{}, fmt.Errorf("producto %d no encontrado", id)
	}
	return p, nil
}

func (r *repository) GetAll() ([]Productos, error) {
	if len(products) == 0 {
		return nil, fmt.Errorf("no hay productos registrados")
	}
	return products, nil
}

func (r *repository) UpdatePrecio(id int, precio float64) (Productos, error) {
	updated := false
	var p Productos
	for i := range products {
		if products[i].Id == id {
			products[i].Precio = precio
			p = products[i]
			updated = true
		}
	}
	if !updated {
		return Productos{}, fmt.Errorf("producto %d no encontrado", id)
	}
	return p, nil
}

func (r *repository) Delete(id int) error {
	deleted := false
	var index int
	for i := range products {
		if products[i].Id == id {
			index = i
			deleted = true
		}
	}
	if !deleted {
		return fmt.Errorf("producto %d no encontrado", id)
	}
	products = append(products[:index], products[index+1:]...)
	return nil
}

func (r *repository) LastID() (int, error) {
	return lastID, nil
}
