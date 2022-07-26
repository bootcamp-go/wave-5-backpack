package products

import (
	"testing"

	"github.com/bootcamp-go/wave-5-backpack/goweb/internal/domain"
	"github.com/stretchr/testify/assert"
)

type StubRepository struct{}

func (s StubRepository) GetAll() ([]domain.Product, error) {
	return []domain.Product{
		{Nombre: "TV Samsung", Color: "Negro", Precio: 10},
		{Nombre: "TV LG", Color: "Gris", Precio: 15}}, nil
}

func (s StubRepository) GetProduct(id int) (domain.Product, error) {
	return domain.Product{}, nil
}

func (s StubRepository) Store(id int, nombre string, color string, precio float64, stock int, codigo string, publicado bool) (domain.Product, error) {
	return domain.Product{}, nil
}

func (s StubRepository) LastID() (int, error) {
	return 0, nil
}

func (s StubRepository) UpdateAll(id int, nombre string, color string, precio float64, stock int, codigo string, publicado bool) (domain.Product, error) {
	return domain.Product{}, nil

}

func (s StubRepository) Delete(id int) error {
	return nil
}

func (s StubRepository) Update(id int, nombre string, precio float64) (domain.Product, error) {
	return domain.Product{}, nil
}

func TestSGetAll(t *testing.T) {
	stub := StubRepository{}
	service := NewService(stub)

	productsExpected := []domain.Product{{Nombre: "TV Samsung", Color: "Negro", Precio: 10}, {Nombre: "TV LG", Color: "Gris", Precio: 15}}

	products, _ := service.GetAll()

	assert.Equal(t, products, productsExpected)

}
