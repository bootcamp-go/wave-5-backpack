package products

import (
	"encoding/json"
	"io/ioutil"

	"github.com/bootcamp-go/wave-5-backpack/goweb/internal/domain"
)

var products []domain.Product
var idLast int

type Repository interface {
	GetAll() ([]domain.Product, error)
	Store(id int, nombre string, color string, precio float64, stock int, codigo string, publicado bool, FechaCreacion string) (domain.Product, error)
	LastID() (int, error)
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r repository) GetAll() ([]domain.Product, error) {

	if err := Read(); err != nil {
		return nil, err
	}

	return products, nil
}

func (r repository) Store(id int, nombre string, color string, precio float64, stock int, codigo string, publicado bool, FechaCreacion string) (domain.Product, error) {
	product := domain.Product{
		ID:            id,
		Nombre:        nombre,
		Color:         color,
		Precio:        precio,
		Stock:         stock,
		Codigo:        codigo,
		Publicado:     publicado,
		FechaCreacion: FechaCreacion}

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
