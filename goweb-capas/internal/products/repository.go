package products

import "fmt"

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"nombre"`
	Type  string  `json:"tipo"`
	Count int     `json:"cantidad"`
	Price float64 `json:"precio"`
}

var Products []Product
var lastID int

type Repository interface {
	GetAll() ([]Product, error)
	Store(id int, nombre, tipo string, cantidad int, precio float64) (Product, error)
	LastID() (int, error)
	Update(id int, nombre, tipo string, cantidad int, precio float64) (Product, error)
	Patch(id int, nombre string, precio float64) (Product, error)
	Delete(id int) error
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetAll() ([]Product, error) {
	return Products, nil
}

func (r *repository) LastID() (int, error) {
	return lastID, nil
}

func (r *repository) Store(id int, nombre, tipo string, cantidad int, precio float64) (Product, error) {
	p := Product{id, nombre, tipo, cantidad, precio}
	Products = append(Products, p)
	lastID = p.ID
	return p, nil
}

func (r *repository) Update(id int, nombre, tipo string, cantidad int, precio float64) (Product, error) {
	p := Product{Name: nombre, Type: tipo, Count: cantidad, Price: precio}
	updated := false
	for i := range Products {
		if Products[i].ID == id {
			p.ID = id
			Products[i] = p
			updated = true
		}
	}

	if !updated {
		return Product{}, fmt.Errorf("producto %d no encontrado", id)
	}

	return p, nil
}

func (r *repository) Patch(id int, nombre string, precio float64) (Product, error) {
	updated := false
	var p Product
	for i := range Products {
		if Products[i].ID == id {
			Products[i].Name = nombre
			Products[i].Price = precio
			p = Products[i]
			updated = true
		}
	}

	if !updated {
		return Product{}, fmt.Errorf("producto %d no encontrado", id)
	}

	return p, nil
}

func (r *repository) Delete(id int) error {
	deleted := false
	var index int
	for i := range Products {
		if Products[i].ID == id {
			index = i
			deleted = true
		}
	}

	if !deleted {
		return fmt.Errorf("producto %d no encontrado", id)
	}

	Products = append(Products[:index], Products[index+1:]...)
	return nil
}
