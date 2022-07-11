package products

import (
	"errors"
	"time"

	"github.com/bootcamp-go/wave-5-backpack/goweb/internal/domain"
	"github.com/bootcamp-go/wave-5-backpack/goweb/pkg/store"
)

var idLast int

type Repository interface {
	GetAll() ([]domain.Product, error)
	GetProduct(id int) (domain.Product, error)
	Store(id int, nombre string, color string, precio float64, stock int, codigo string, publicado bool) (domain.Product, error)
	LastID() (int, error)
	UpdateAll(id int, nombre string, color string, precio float64, stock int, codigo string, publicado bool) (domain.Product, error)
	Delete(id int) error
	Update(id int, nombre string, precio float64) (domain.Product, error)
}

type repository struct {
	db store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{
		db: db,
	}
}

func (r repository) GetAll() ([]domain.Product, error) {
	var products []domain.Product
	if err := r.db.Read(&products); err != nil {
		return nil, err
	}

	return products, nil
}

func (r repository) GetProduct(id int) (domain.Product, error) {
	var products []domain.Product
	if err := r.db.Read(&products); err != nil {
		return domain.Product{}, err
	}

	for _, p := range products {
		if p.ID == id {
			return p, nil
		}
	}

	return domain.Product{}, errors.New("El producto no existe")
}

func (r repository) Store(id int, nombre string, color string, precio float64, stock int, codigo string, publicado bool) (domain.Product, error) {
	var products []domain.Product
	r.db.Read(&products)

	product := domain.Product{
		ID:            id,
		Nombre:        nombre,
		Color:         color,
		Precio:        precio,
		Stock:         stock,
		Codigo:        codigo,
		Publicado:     publicado,
		FechaCreacion: time.Now().Local().String(),
	}

	products = append(products, product)
	if err := r.db.Write(products); err != nil {
		return domain.Product{}, err
	}

	return product, nil
}

func (r repository) LastID() (int, error) {
	var products []domain.Product
	if err := r.db.Read(&products); err != nil {
		return 0, err
	}

	if len(products) == 0 {
		return 1, nil
	}
	maxID := products[len(products)-1].ID

	return (maxID + 1), nil
}

func (r repository) UpdateAll(id int, nombre string, color string, precio float64, stock int, codigo string, publicado bool) (domain.Product, error) {
	var products []domain.Product
	if err := r.db.Read(&products); err != nil {
		return domain.Product{}, err
	}

	product := domain.Product{
		Nombre:    nombre,
		Color:     color,
		Precio:    precio,
		Stock:     stock,
		Codigo:    codigo,
		Publicado: publicado,
	}

	for i, p := range products {
		if p.ID == id {
			product.ID = id
			product.FechaCreacion = p.FechaCreacion
			products[i] = product
			r.db.Write(&products)
			return product, nil
		}
	}

	return domain.Product{}, errors.New("No fue posible encotrar el producto a modificar")

}

func (r repository) Delete(id int) error {
	var products []domain.Product
	if err := r.db.Read(&products); err != nil {
		return err
	}

	for i, p := range products {
		if p.ID == id {
			products = append(products[:i], products[i+1:]...)
			r.db.Write(&products)
			return nil
		}
	}
	return errors.New("no fue posible encontrar el producto a modificar")
}

func (r repository) Update(id int, nombre string, precio float64) (domain.Product, error) {
	var products []domain.Product
	if err := r.db.Read(&products); err != nil {
		return domain.Product{}, err
	}

	for i, p := range products {
		if p.ID == id {
			p.Nombre = nombre
			p.Precio = precio
			products[i] = p
			r.db.Write(&products)
			return p, nil
		}
	}

	return domain.Product{}, errors.New("no fue posible encontrar el producto a modificar")
}
