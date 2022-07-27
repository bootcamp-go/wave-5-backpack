package products

import (
	"errors"
	"fmt"
	"testing"

	"github.com/bootcamp-go/wave-5-backpack/goweb/internal/domain"
	"github.com/stretchr/testify/assert"
)

type StubStore struct {
	errGetAll string
}

func (ss StubStore) Ping() error {
	return nil
}

func (ss StubStore) Write(data interface{}) error {
	return nil
}

func (ss *StubStore) Read(data interface{}) error {
	if ss.errGetAll != "" {
		return errors.New(ss.errGetAll)
	}
	products := data.(*[]domain.Product)
	*products = []domain.Product{
		{ID: 1, Nombre: "TV SAMSUNG", Color: "Negro", Precio: 10},
		{ID: 2, Nombre: "TV LG", Color: "Negro", Precio: 15}}

	return nil
}

func TestGetAllSuccess(t *testing.T) {
	stub := &StubStore{}
	repository := NewRepository(stub)

	productsExpected := []domain.Product{
		{ID: 1, Nombre: "TV SAMSUNG", Color: "Negro", Precio: 10},
		{ID: 2, Nombre: "TV LG", Color: "Negro", Precio: 15}}

	products, err := repository.GetAll()

	assert.Nil(t, err)
	assert.Equal(t, productsExpected, products)
}

func TestGetAllErrorRead(t *testing.T) {
	stub := &StubStore{
		errGetAll: "error al leer el archivo",
	}
	repository := NewRepository(stub)

	errorExpected := fmt.Sprintf("%s: %s", stub.errGetAll, ERROR_GET_ALL)
	product, err := repository.GetAll()

	assert.Nil(t, product)
	assert.ErrorContains(t, err, errorExpected)
}
