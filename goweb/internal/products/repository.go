package products

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"time"

	"github.com/bootcamp-go/wave-5-backpack/goweb/internal/domain"
)

var products []domain.Product
var idLast int

type Repository interface {
	GetAll() ([]domain.Product, error)
	Store(id int, nombre string, color string, precio float64, stock int, codigo string, publicado bool) (domain.Product, error)
	LastID() (int, error)
	UpdateAll(id int, nombre string, color string, precio float64, stock int, codigo string, publicado bool) (domain.Product, error)
	Delete(id int) error
	Update(id int, nombre string, precio float64) (domain.Product, error)
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r repository) GetAll() ([]domain.Product, error) {

	// if err := Read(); err != nil {
	// 	return nil, err
	// }

	return products, nil
}

func (r repository) Store(id int, nombre string, color string, precio float64, stock int, codigo string, publicado bool) (domain.Product, error) {
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
	return product, nil
}

func (r repository) LastID() (int, error) {
	maxID := 0
	for _, product := range products {
		if product.ID > maxID {
			maxID = product.ID
		}
	}

	return (maxID + 1), nil
}

func Read() error {
	jsonData, err := ioutil.ReadFile("./products.json")

	if err != nil {
		return err
	}

	if err := json.Unmarshal(jsonData, &products); err != nil {
		return err
	}

	return nil
}

func (r repository) UpdateAll(id int, nombre string, color string, precio float64, stock int, codigo string, publicado bool) (domain.Product, error) {
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
			return product, nil
		}
	}

	return domain.Product{}, errors.New("No fue posible encotrar el producto a modificar")

}

func (r repository) Delete(id int) error {
	for i, p := range products {
		if p.ID == id {
			products = append(products[:i], products[i+1:]...)
			return nil
		}
	}
	return errors.New("no fue posible encontrar el producto a modificar")
}

func (r repository) Update(id int, nombre string, precio float64) (domain.Product, error) {

	for i, p := range products {
		if p.ID == id {
			p.Nombre = nombre
			p.Precio = precio
			products[i] = p
			return p, nil
		}
	}

	return domain.Product{}, errors.New("no fue posible encontrar el producto a modificar")
}
