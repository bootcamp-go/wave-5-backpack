package products

import (
	"fmt"
	"testing-goweb/internal/domain"
	"testing-goweb/pkg/store"
)

//Mensajes de Respuestas Estandar
const (
	NoReading       = "no se puede leer DB"
	NoWriting       = "no se puede escribir DB, error: %v"
	ProductNotFound = "producto con id %d no se encuentra en DB"
)

//Interfaz de Repository
type Repository interface {
	GetAll() ([]domain.Product, error)
	Store(id int, nombre, color string, precio, stock int, codigo string, publicdo bool, fechaCreacion string) (domain.Product, error)
	Update(id int, nombre, color string, precio, stock int, codigo string, publicdo bool, fechaCreacion string) (domain.Product, error)
	UpdatePrecioStock(id, precio, stock int) (domain.Product, error)
	Delete(id int) error
	LastID() (int, error)
}

type repository struct {
	db store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{
		db: db,
	}
}

//GetAll - listar todos los productos
func (r *repository) GetAll() ([]domain.Product, error) {
	var list []domain.Product
	if err := r.db.Read(&list); err != nil {
		return nil, fmt.Errorf(NoReading)
	}
	return list, nil
}

//Store - Crear producto nuevo
func (r *repository) Store(id int, nombre, color string, precio, stock int, codigo string, publicdo bool, fechaCreacion string) (domain.Product, error) {
	var list []domain.Product
	if err := r.db.Read(&list); err != nil {
		return domain.Product{}, fmt.Errorf(NoReading)
	}
	p := domain.Product{
		ID:            id,
		Nombre:        nombre,
		Color:         color,
		Precio:        precio,
		Stock:         stock,
		Codigo:        codigo,
		Publicado:     publicdo,
		FechaCreacion: fechaCreacion,
	}
	list = append(list, p)

	if err := r.db.Write(list); err != nil {
		return domain.Product{}, fmt.Errorf(NoWriting, err)
	}
	return p, nil
}

//Update - se actualiza un producto por completo
func (r *repository) Update(id int, nombre, color string, precio, stock int, codigo string, publicdo bool, fechaCreacion string) (domain.Product, error) {
	var list []domain.Product
	if err := r.db.Read(&list); err != nil {
		return domain.Product{}, fmt.Errorf(NoReading)
	}
	p := domain.Product{
		ID:            id,
		Nombre:        nombre,
		Color:         color,
		Precio:        precio,
		Stock:         stock,
		Codigo:        codigo,
		Publicado:     publicdo,
		FechaCreacion: fechaCreacion,
	}
	updated := false
	for i := range list {
		if list[i].ID == id {
			list[i] = p
			updated = true
			break
		}
	}
	if !updated {
		return domain.Product{}, fmt.Errorf(ProductNotFound, id)
	}

	if err := r.db.Write(list); err != nil {
		return domain.Product{}, fmt.Errorf(NoWriting, err)
	}
	return p, nil
}

//UpdatePrecioStock - se debe actualizar precio y stock
func (r *repository) UpdatePrecioStock(id, precio, stock int) (domain.Product, error) {
	var list []domain.Product
	if err := r.db.Read(&list); err != nil {
		return domain.Product{}, fmt.Errorf(NoReading)
	}

	updated := false
	var p domain.Product
	for i := range list {
		if list[i].ID == id {
			list[i].Precio = precio
			list[i].Stock = stock
			p = list[i]
			updated = true
		}
	}

	if !updated {
		return domain.Product{}, fmt.Errorf(ProductNotFound, id)
	}
	if err := r.db.Write(list); err != nil {
		return domain.Product{}, fmt.Errorf(NoWriting, err)
	}
	return p, nil
}

//Delete - Se elimina la ID que se buscara
func (r *repository) Delete(id int) error {
	var list []domain.Product
	if err := r.db.Read(&list); err != nil {
		return fmt.Errorf(NoReading)
	}

	deleted := false
	var index int
	for i := range list {
		if list[i].ID == id {
			index = i
			deleted = true
		}
	}

	if !deleted {
		return fmt.Errorf(ProductNotFound, id)
	}
	list = append(list[:index], list[index+1:]...)

	if err := r.db.Write(list); err != nil {
		return fmt.Errorf(NoWriting, err)
	}
	return nil
}

//LastID - se obtiene la ultima id desde la DB
func (r *repository) LastID() (int, error) {
	var list []domain.Product
	if err := r.db.Read(&list); err != nil {
		return 0, fmt.Errorf(NoReading)
	}
	if len(list) == 0 {
		return 0, nil
	}
	return list[len(list)-1].ID, nil
}
