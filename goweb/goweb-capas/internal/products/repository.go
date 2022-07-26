package products

import (
	"fmt"
	"goweb-capas/pkg/store"
)

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"nombre"`
	Type  string  `json:"tipo"`
	Count int     `json:"cantidad"`
	Price float64 `json:"precio"`
}

var lastID int

const (
	ProductNotFound = "product %d not found"
	FailReading     = "cant read database"
	FailWriting     = "cant write database, error: %w"
)

type Repository interface {
	GetAll() ([]Product, error)
	Store(id int, nombre, tipo string, cantidad int, precio float64) (Product, error)
	LastID() (int, error)
	Update(id int, nombre, tipo string, cantidad int, precio float64) (Product, error)
	Patch(id int, nombre string, precio float64) (Product, error)
	Delete(id int) error
}

type repository struct {
	db store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll() ([]Product, error) {
	var Products []Product
	if err := r.db.Read(&Products); err != nil {
		return nil, fmt.Errorf(FailReading)
	}
	return Products, nil
}

func (r *repository) LastID() (int, error) {
	var Products []Product
	if err := r.db.Read(&Products); err != nil {
		return 0, fmt.Errorf(FailReading)
	}
	if len(Products) == 0 {
		return 0, nil
	}

	return Products[len(Products)-1].ID, nil
}

func (r *repository) Store(id int, nombre, tipo string, cantidad int, precio float64) (Product, error) {
	var Products []Product

	if err := r.db.Read(&Products); err != nil {
		return Product{}, fmt.Errorf(FailReading)
	}

	p := Product{id, nombre, tipo, cantidad, precio}
	Products = append(Products, p)

	if err := r.db.Write(Products); err != nil {
		return Product{}, fmt.Errorf(FailWriting, err)
	}

	return p, nil
}

func (r *repository) Update(id int, nombre, tipo string, cantidad int, precio float64) (Product, error) {
	var Products []Product

	if err := r.db.Read(&Products); err != nil {
		return Product{}, fmt.Errorf(FailReading)
	}

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
		return Product{}, fmt.Errorf(ProductNotFound, id)
	}

	if err := r.db.Write(Products); err != nil {
		return Product{}, fmt.Errorf(FailWriting, err)
	}

	return p, nil
}

func (r *repository) Patch(id int, nombre string, precio float64) (Product, error) {
	var Products []Product

	if err := r.db.Read(&Products); err != nil {
		return Product{}, fmt.Errorf(FailReading)
	}

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
		return Product{}, fmt.Errorf(ProductNotFound, id)
	}

	if err := r.db.Write(Products); err != nil {
		return Product{}, fmt.Errorf(FailWriting, err)
	}

	return p, nil
}

func (r *repository) Delete(id int) error {
	var Products []Product

	if err := r.db.Read(&Products); err != nil {
		return fmt.Errorf(FailReading)
	}

	deleted := false
	var index int
	for i := range Products {
		if Products[i].ID == id {
			index = i
			deleted = true
		}
	}

	if !deleted {
		return fmt.Errorf(ProductNotFound, id)
	}

	Products = append(Products[:index], Products[index+1:]...)
	if err := r.db.Write(Products); err != nil {
		return fmt.Errorf(FailWriting, err)
	}
	return nil
}
