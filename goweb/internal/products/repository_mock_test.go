package products

import (
	"errors"
	"fmt"
	"testing"

	"github.com/bootcamp-go/wave-5-backpack/goweb/internal/domain"
	"github.com/stretchr/testify/assert"
)

type MockStore struct {
	readCalled  bool
	writeCalled bool
	db          []domain.Product
	errRead     string
	errWrite    string
}

func (ms MockStore) Ping() error {
	return nil
}

func (ms *MockStore) Write(data interface{}) error {
	ms.writeCalled = true
	if ms.errWrite != "" {
		return errors.New(ms.errWrite)
	}

	product := data.(*[]domain.Product)
	ms.db = *product

	return nil
}

func (ms *MockStore) Read(data interface{}) error {
	ms.readCalled = true
	if ms.errRead != "" {
		return errors.New(ms.errRead)
	}

	products := data.(*[]domain.Product)
	*products = []domain.Product{
		{ID: 1, Nombre: "TV SAMSUNG", Color: "Negro", Precio: 10},
		{ID: 2, Nombre: "TV LG", Color: "Negro", Precio: 15}}

	ms.db = append(ms.db, *products...)

	return nil
}

func TestUpdateSuccess(t *testing.T) {
	mock := &MockStore{}
	repository := NewRepository(mock)

	errorUpdate, errorID := repository.Update(3, "", 0)
	afterUpdate, err := repository.Update(1, "TV LG", 15)

	assert.Nil(t, err)
	assert.Empty(t, errorUpdate)
	assert.ErrorContains(t, errorID, ERROR_ID_NOT_EXIST)
	assert.Equal(t, mock.db[0], afterUpdate)
	assert.True(t, mock.readCalled)
	assert.True(t, mock.writeCalled)
}

func TestUpdateErrorRead(t *testing.T) {
	mock := &MockStore{
		errRead: "error al leer el archivo",
	}
	repository := NewRepository(mock)

	errorExpected := fmt.Sprintf("%s: %s", mock.errRead, ERROR_UPDATE)
	afterUpdate, err := repository.Update(1, "TV LG", 15)

	assert.Empty(t, afterUpdate)
	assert.ErrorContains(t, err, errorExpected)
	assert.True(t, mock.readCalled)
}

func TestUpdateErrorWrite(t *testing.T) {
	mock := &MockStore{
		errWrite: "error al escribir el archivo",
	}
	repository := NewRepository(mock)

	errorExpected := fmt.Sprintf("%s: %s", mock.errWrite, ERROR_UPDATE)
	afterUpdate, err := repository.Update(1, "TV LG", 15)

	assert.Empty(t, afterUpdate)
	assert.ErrorContains(t, err, errorExpected)
	assert.True(t, mock.writeCalled)
}

func TestDeleteErrorWrite(t *testing.T) {
	mock := &MockStore{
		errWrite: "error al escribir el archivo",
	}
	repository := NewRepository(mock)

	errorExpected := fmt.Sprintf("%s: %s", mock.errWrite, ERROR_DELETE)
	err := repository.Delete(1)

	assert.ErrorContains(t, err, errorExpected)
	assert.True(t, mock.writeCalled)
}

func TestDeleteErrorRead(t *testing.T) {
	mock := &MockStore{
		errRead: "error al leer el archivo",
	}
	repository := NewRepository(mock)

	errorExpected := fmt.Sprintf("%s: %s", mock.errRead, ERROR_DELETE)
	err := repository.Delete(1)

	assert.ErrorContains(t, err, errorExpected)
	assert.True(t, mock.readCalled)
}
