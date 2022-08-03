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
}

type repository struct {
	db store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{db: db}
}

func (r *repository) GetAll() ([]*Product, error) {
	var products []*Product
	if err := r.db.Read(&products); err != nil {
		return nil, err
	}

	return products, nil
}

func (r *repository) UpdateName(id int, nombre string) (*Product, error) {
	var updated bool = false
	var products []*Product
	if err := r.db.Read(&products); err != nil {
		return nil, err
	}

	var product *Product
	for _, value := range products {
		if value.Id == id {
			value.Nombre = nombre
			product = value
			updated = true
		}
	}

	if !updated {
		return nil, fmt.Errorf("producto id %d no encontrado", id)
	}

	if err := r.db.Write(&products); err != nil {
		return nil, err
	}

	return product, nil
}
