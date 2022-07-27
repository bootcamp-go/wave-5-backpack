package products

import (
	"testing"

	"github.com/bootcamp-go/wave-5-backpack/goweb/internal/domain"
	"github.com/stretchr/testify/assert"
)

type mockStore struct {
	products   []domain.Product
	readCalled bool
	errID      string
}

func (ms mockStore) Ping() error {
	return nil
}

func (ms *mockStore) Write(data interface{}) error {
	product := data.(*[]domain.Product)
	ms.products = *product

	return nil
}

func (ms *mockStore) Read(data interface{}) error {
	products := data.(*[]domain.Product)
	*products = []domain.Product{
		{ID: 1, Nombre: "TV SAMSUNG", Color: "Negro", Precio: 10},
		{ID: 2, Nombre: "TV LG", Color: "Negro", Precio: 15}}

	ms.products = append(ms.products, *products...)

	ms.readCalled = true

	return nil
}

func TestIntegrationUpdate(t *testing.T) {
	mock := &mockStore{}
	repository := NewRepository(mock)
	service := NewService(repository)

	productUpdate, err := service.Update(1, "TV", 20)

	assert.Nil(t, err)
	assert.True(t, mock.readCalled)
	assert.Equal(t, productUpdate.ID, 1)
	assert.Equal(t, productUpdate, mock.products[0])

}

func TestIntegrationDelete(t *testing.T) {
	mock := &mockStore{
		errID: ERROR_ID_NOT_EXIST,
	}
	repository := NewRepository(mock)
	service := NewService(repository)

	err := service.Delete(3)
	notErr := service.Delete(1)

	assert.Nil(t, notErr)
	assert.ErrorContains(t, err, mock.errID)
	assert.Equal(t, len(mock.products), 1)
}
