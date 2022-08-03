package products

import (
	"fmt"

	"github.com/bootcamp-go/go-testing/pkg/store"
)

type Product struct {
	Id     int     `json:"id"`
	Nombre string  `json:"nombre"`
	Stock  int     `json:"stock"`
	Precio float64 `json:"precio"`
}

type Repository interface {
	GetAll() ([]*Product, error)
	UpdateName(id int, nombre string) (*Product, error)
	Update(id int, nombre string, stock int, precio float64) (*Product, error)
	Delete(id int) error
}

type repository struct {
	db store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{db: db}
}

func (r *repository) GetAll() ([]*Product, error) {
	var producp []*Product
	if err := r.db.Read(&producp); err != nil {
		return nil, err
	}

	return producp, nil
}

func (r *repository) UpdateName(id int, nombre string) (*Product, error) {
	var updated bool = false
	var producp []*Product
	if err := r.db.Read(&producp); err != nil {
		return nil, err
	}

	var product *Product
	for _, value := range producp {
		if value.Id == id {
			value.Nombre = nombre
			product = value
			updated = true
		}
	}

	if !updated {
		return nil, fmt.Errorf("producto id %d no encontrado", id)
	}

	if err := r.db.Write(&producp); err != nil {
		return nil, err
	}

	return product, nil
}

func (r *repository) Update(id int, nombre string, stock int, precio float64) (*Product, error) {
	update := false
	transactionNew := &Product{
		Id:     id,
		Nombre: nombre,
		Stock:  stock,
		Precio: precio,
	}

	var p []*Product
	if err := r.db.Read(&p); err != nil {
		return nil, err
	}

	for value := range p {
		if p[value].Id == id {
			p[value] = transactionNew
			update = true
		}
	}

	if !update {
		return nil, fmt.Errorf("transacción id %d no encontrada", id)
	}

	if err := r.db.Write(&p); err != nil {
		return nil, err
	}

	return transactionNew, nil
}

func (r *repository) Delete(id int) error {
	deleted := false
	var indice int

	var p []*Product
	if err := r.db.Read(&p); err != nil {
		return err
	}
	for value := range p {
		if p[value].Id == id {
			indice = value
			deleted = true
		}
	}

	if !deleted {
		return fmt.Errorf("la transacción id %d no existe", id)
	}

	p = append(p[:indice], p[indice+1:]...)
	if err := r.db.Write(&p); err != nil {
		return err
	}

	return nil
}
