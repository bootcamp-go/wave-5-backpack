package products

import (
	"testing"

	"github.com/bootcamp-go/wave-5-backpack/goweb/internal/domain"
	"github.com/stretchr/testify/assert"
)

type StubStore struct{}

func (ss StubStore) Ping() error {
	return nil
}

func (ss StubStore) Write(data interface{}) error {
	return nil
}

func (ss StubStore) Read(data interface{}) error {
	products := data.(*[]domain.Product)
	*products = []domain.Product{
		{ID: 1, Nombre: "TV SAMSUNG", Color: "Negro", Precio: 10},
		{ID: 2, Nombre: "TV LG", Color: "Negro", Precio: 15}}

	return nil
}

func TestRGetAll(t *testing.T) {
	stub := &StubStore{}
	repository := NewRepository(stub)

	productsExpected := []domain.Product{
		{ID: 1, Nombre: "TV SAMSUNG", Color: "Negro", Precio: 10},
		{ID: 2, Nombre: "TV LG", Color: "Negro", Precio: 15}}

	products, err := repository.GetAll()

	assert.Nil(t, err)
	assert.Equal(t, productsExpected, products)
}
