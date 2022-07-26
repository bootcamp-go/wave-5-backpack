package products

import (
	"reflect"
	"testing"

	"github.com/bootcamp-go/wave-5-backpack/goweb/internal/domain"
	"github.com/stretchr/testify/assert"
)

type MockStore struct {
	updateCall bool
	db         []domain.Product
}

func (ms MockStore) Ping() error {
	return nil
}

func (ms MockStore) Write(data interface{}) error {
	return nil
}

func (ms *MockStore) Read(data interface{}) error {
	products := reflect.ValueOf(data)
	products = reflect.Indirect(products)
	products.Set(reflect.ValueOf(ms.db))

	ms.updateCall = true

	return nil
}

func TestUpdate(t *testing.T) {
	mock := &MockStore{
		db: []domain.Product{{ID: 1, Nombre: "TV SAMSUNG", Color: "Negro", Precio: 10}},
	}
	repository := NewRepository(mock)

	afterUpdate, _ := repository.Update(1, "TV LG", 15)

	assert.Equal(t, mock.db[0], afterUpdate)
	assert.True(t, mock.updateCall)
}
