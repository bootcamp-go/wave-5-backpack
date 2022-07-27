package products

import (
	"fmt"
	"goweb/internal/domain"
	"goweb/pkg/store"
)

var ps []domain.Products

type Repository interface {
	GetAll() ([]domain.Products, error)
	CreateProduct(id int, nombre, color string, precio float64, stock int, código string, publicado bool, fecha_de_creación string) (domain.Products, error)
	LastID() (int, error)
	Update(id int, nombre, color string, precio float64, stock int, código string, publicado bool, fecha_de_creación string) (domain.Products, error)
	Delete(id int) error
	UpdateOne(id int, nombre string, precio float64) (domain.Products, error)
}

const (
	ProductNotFound = "Producto %d no encontrado"
	FailReading     = "No se pudo leer el archivo"
	FailWriting     = "No se pudo escribir el archivo, error: %w"
)

type repository struct {
	db store.Store
}

func InitRepository(db store.Store) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll() ([]domain.Products, error) {
	var ps2 []domain.Products
	if err := r.db.Read(&ps2); err != nil {
		fmt.Println("Ocurrio un error")
		return nil, fmt.Errorf(FailReading)
	}
	fmt.Println("repository: ", ps2)
	return ps2, nil
}

func (r *repository) LastID() (int, error) {
	var ps2 []domain.Products
	if err := r.db.Read(&ps2); err != nil {
		return 0, fmt.Errorf(FailReading)
	}
	if len(ps2) == 0 {
		return 0, nil
	}
	return ps2[len(ps2)-1].Id, nil
}

func (r *repository) CreateProduct(id int, nombre, color string, precio float64, stock int, código string, publicado bool, fecha_de_creación string) (domain.Products, error) {
	var ps2 []domain.Products
	if err := r.db.Read(&ps2); err != nil {
		return domain.Products{}, fmt.Errorf(FailReading)
	}
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
	ps2 = append(ps2, p)
	if err := r.db.Write(ps2); err != nil {
		return domain.Products{}, fmt.Errorf(FailWriting, err)
	}

	return p, nil
}

func (r *repository) Update(id int, nombre, color string, precio float64, stock int, código string, publicado bool, fecha_de_creación string) (domain.Products, error) {
	var ps2 []domain.Products
	if err := r.db.Read(&ps2); err != nil {
		return domain.Products{}, fmt.Errorf(FailReading)
	}
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
	for i := range ps2 {
		if ps2[i].Id == id {
			p.Id = id
			ps2[i] = p
			updated = true
		}
	}
	if !updated {
		return domain.Products{}, fmt.Errorf("Producto %d no encontrado", id)
	}
	if err := r.db.Write(ps2); err != nil {
		return domain.Products{}, fmt.Errorf(FailWriting, err)
	}
	return p, nil
}

func (r *repository) Delete(id int) error {
	var ps2 []domain.Products
	if err := r.db.Read(&ps2); err != nil {
		return fmt.Errorf(FailReading)
	}
	deleted := false
	var index int
	for i := range ps2 {
		if ps2[i].Id == id {
			index = i
			deleted = true
		}
	}
	if !deleted {
		return fmt.Errorf("Producto %d no encontrado", id)
	}
	ps2 = append(ps2[:index], ps2[index+1:]...)
	if err := r.db.Write(ps2); err != nil {
		return fmt.Errorf(FailWriting, err)
	}
	return nil

}

func (r *repository) UpdateOne(id int, nombre string, precio float64) (domain.Products, error) {
	var ps2 []domain.Products
	if err := r.db.Read(&ps2); err != nil {
		return domain.Products{}, fmt.Errorf(FailReading)
	}
	var p domain.Products
	updated := false
	for i := range ps2 {
		if ps2[i].Id == id {
			ps2[i].Nombre = nombre
			ps2[i].Precio = precio
			updated = true
			p = ps2[i]
		}
	}
	if !updated {
		return domain.Products{}, fmt.Errorf("Producto %d no encontrado", id)
	}
	if err := r.db.Write(ps2); err != nil {
		return domain.Products{}, fmt.Errorf(FailWriting, err)
	}
	return p, nil
}
