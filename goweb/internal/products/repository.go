package products

import (
	"fmt"
	"web-server/internal/domain"
	"web-server/pkg/store"
)

const (
	ProductNotFound = "product %d not found"
	FailReading     = "cant read database"
	FailWriting     = "cant write database, error: %w"
)

type Repository interface {
	GetAll() ([]domain.Products, error)
	LastID() (int, error)
	Store(id int, nombre string, color string, precio float64, stock int, codigo string, publicado bool, fecha string) (domain.Products, error)
	Update(id int, nombre string, color string, precio float64, stock int, codigo string, publicado bool, fecha string) (domain.Products, error)
	UpdateName(id int, nombre string) (domain.Products, error)
	UpdatePrice(id int, precio float64) (domain.Products, error)
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

func (r *repository) GetAll() ([]domain.Products, error) {
	var productsSlide []domain.Products
	if err := r.db.Read(&productsSlide); err != nil {
		return nil, fmt.Errorf(FailReading)
	}
	return productsSlide, nil
}

func (r *repository) LastID() (int, error) {
	var productsSlide []domain.Products
	if err := r.db.Read(&productsSlide); err != nil {
		return 0, fmt.Errorf(FailReading)
	}

	if len(productsSlide) == 0 {
		return 0, nil
	}

	return productsSlide[len(productsSlide)-1].Id, nil
}

func (r *repository) Store(id int, nombre string, color string, precio float64, stock int, codigo string, publicado bool, fecha string) (domain.Products, error) {
	var productsSlide []domain.Products
	if err := r.db.Read(&productsSlide); err != nil {
		return domain.Products{}, fmt.Errorf(FailReading)
	}

	p := domain.Products{Id: id, Nombre: nombre, Color: color, Precio: precio, Stock: stock, Codigo: codigo, Publicado: publicado, Fecha: fecha}
	productsSlide = append(productsSlide, p)

	if err := r.db.Write(productsSlide); err != nil {
		return domain.Products{}, fmt.Errorf(FailWriting, err)
	}

	return p, nil
}

func (r *repository) Update(id int, nombre string, color string, precio float64, stock int, codigo string, publicado bool, fecha string) (domain.Products, error) {
	var productsSlide []domain.Products

	if err := r.db.Read(&productsSlide); err != nil {
		return domain.Products{}, fmt.Errorf(FailReading)
	}

	p := domain.Products{Nombre: nombre, Color: color, Precio: precio, Stock: stock, Codigo: codigo, Publicado: publicado, Fecha: fecha}
	updated := false
	for i := range productsSlide {
		if productsSlide[i].Id == id {
			p.Id = id
			productsSlide[i] = p
			updated = true
		}
	}

	if !updated {
		return domain.Products{}, fmt.Errorf("producto %d no encontrado", id)
	}

	return p, nil
}

func (r *repository) UpdateName(id int, nombre string) (domain.Products, error) {
	var productsSlide []domain.Products

	updated := false
	var p domain.Products
	for i := range productsSlide {
		if productsSlide[i].Id == id {
			productsSlide[i].Nombre = nombre
			p = productsSlide[i]
			updated = true
		}
	}

	if !updated {
		return domain.Products{}, fmt.Errorf("producto %d no encontrado", id)
	}

	return p, nil
}

func (r *repository) UpdatePrice(id int, precio float64) (domain.Products, error) {
	var productsSlide []domain.Products

	updated := false
	var p domain.Products
	for i := range productsSlide {
		if productsSlide[i].Id == id {
			productsSlide[i].Precio = precio
			p = productsSlide[i]
			updated = true
		}
	}

	if !updated {
		return domain.Products{}, fmt.Errorf("producto %d no encontrado", id)
	}

	return p, nil
}

func (r *repository) Delete(id int) error {
	var productsSlide []domain.Products

	if err := r.db.Read(&productsSlide); err != nil {
		return fmt.Errorf(FailReading)
	}

	deleted := false
	var index int
	for i := range productsSlide {
		if productsSlide[i].Id == id {
			index = i
			deleted = true
		}
	}

	if !deleted {
		return fmt.Errorf(ProductNotFound, id)
	}

	productsSlide = append(productsSlide[:index], productsSlide[index+1:]...)

	if err := r.db.Write(productsSlide); err != nil {
		return fmt.Errorf(FailWriting, err)
	}

	return nil
}
