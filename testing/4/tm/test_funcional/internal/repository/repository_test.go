package repository

import (
	"errors"
	"testing"
	"testing/4/tm/test_funcional/internal/domain"
	"testing/4/tm/test_funcional/pkg/store"

	"github.com/stretchr/testify/assert"
)

func TestCreateFailRead(t *testing.T) {
	mockStore := store.MockStore{
		ReadWasCalled:  false,
		WriteWasCalled: false,
		ErrRead:        "error de lectura",
		ErrWrite:       "",
	}
	r := NewRepository(&mockStore)

	productEsperado := domain.Product{}
	errEsperado := errors.New("error de lectura")

	resultado, err := r.Create("Banana", 1.5, 2)

	assert.Equal(t, productEsperado, resultado)
	assert.True(t, mockStore.ReadWasCalled)
	assert.Equal(t, errEsperado, err)
}

func TestCreateFailWrite(t *testing.T) {
	mockStore := store.MockStore{
		ReadWasCalled:  false,
		WriteWasCalled: false,
		ErrRead:        "",
		ErrWrite:       "error de escritura",
	}
	r := NewRepository(&mockStore)

	productEsperado := domain.Product{}
	errEsperado := errors.New("error de escritura")

	resultado, err := r.Create("Banana", 1.5, 2)

	assert.Equal(t, productEsperado, resultado)
	assert.True(t, mockStore.ReadWasCalled)
	assert.True(t, mockStore.WriteWasCalled)
	assert.Equal(t, errEsperado, err)
}

func TestCreate(t *testing.T) {
	mockStore := store.MockStore{
		ReadWasCalled:  false,
		WriteWasCalled: false,
		ErrRead:        "",
		ErrWrite:       "",
	}
	r := NewRepository(&mockStore)

	productEsperado := domain.NewProduct(1, "Banana", 1.5, 2)

	resultado, err := r.Create("Banana", 1.5, 2)

	assert.Equal(t, productEsperado, resultado)
	assert.True(t, mockStore.WriteWasCalled)
	assert.Nil(t, err)
}

func TestReadAllFailRead(t *testing.T) {
	mockStore := store.MockStore{
		ReadWasCalled:  false,
		WriteWasCalled: false,
		ErrRead:        "error de lectura",
		ErrWrite:       "",
	}
	r := NewRepository(&mockStore)

	productsEsperado := []domain.Product{}
	errEsperado := errors.New("error de lectura")

	resultado, err := r.ReadAll()

	assert.Equal(t, productsEsperado, resultado)
	assert.True(t, mockStore.ReadWasCalled)
	assert.Equal(t, err, errEsperado)
}

func TestReadAll(t *testing.T) {
	mockStore := store.MockStore{
		ReadWasCalled:  false,
		WriteWasCalled: false,
		ErrRead:        "",
		ErrWrite:       "",
	}
	r := NewRepository(&mockStore)

	var products []domain.Product
	productsEsperado := append(products, domain.NewProduct(1, "Banana", 1.5, 2))

	resultado, err := r.ReadAll()

	assert.Equal(t, productsEsperado, resultado)
	assert.True(t, mockStore.ReadWasCalled)
	assert.Nil(t, err)
}

func TestReadFailRead(t *testing.T) {
	mockStore := store.MockStore{
		ReadWasCalled:  false,
		WriteWasCalled: false,
		ErrRead:        "error de lectura",
		ErrWrite:       "",
	}
	r := NewRepository(&mockStore)

	productEsperado := domain.Product{}
	errEsperado := errors.New("error de lectura")

	resultado, err := r.Read(1)

	assert.Equal(t, productEsperado, resultado)
	assert.True(t, mockStore.ReadWasCalled)
	assert.Equal(t, err, errEsperado)
}

func TestRead(t *testing.T) {
	mockStore := store.MockStore{
		ReadWasCalled:  false,
		WriteWasCalled: false,
		ErrRead:        "",
		ErrWrite:       "",
	}
	r := NewRepository(&mockStore)

	productEsperado := domain.NewProduct(1, "Banana", 1.5, 2)

	resultado, err := r.Read(1)

	assert.Equal(t, productEsperado, resultado)
	assert.True(t, mockStore.ReadWasCalled)
	assert.Nil(t, err)
}

func TestReadNotFound(t *testing.T) {
	mockStore := store.MockStore{
		ReadWasCalled:  false,
		WriteWasCalled: false,
		ErrRead:        "",
		ErrWrite:       "",
	}
	r := NewRepository(&mockStore)

	productEsperado := domain.Product{}
	errEsperado := errors.New("no se encontro el producto de id 2")

	resultado, err := r.Read(2)

	assert.Equal(t, productEsperado, resultado)
	assert.True(t, mockStore.ReadWasCalled)
	assert.Equal(t, err, errEsperado)
}

func TestUpdateFailRead(t *testing.T) {
	mockStore := store.MockStore{
		ReadWasCalled:  false,
		WriteWasCalled: false,
		ErrRead:        "error de lectura",
		ErrWrite:       "",
	}
	r := NewRepository(&mockStore)

	productEsperado := domain.Product{}
	errEsperado := errors.New("error de lectura")

	resultado, err := r.Update(1, "Manzana", 3, 5)

	assert.Equal(t, productEsperado, resultado)
	assert.True(t, mockStore.ReadWasCalled)
	assert.Equal(t, err, errEsperado)
}

func TestUpdateFailWrite(t *testing.T) {
	mockStore := store.MockStore{
		ReadWasCalled:  false,
		WriteWasCalled: false,
		ErrRead:        "",
		ErrWrite:       "error de escritura",
	}
	r := NewRepository(&mockStore)

	productEsperado := domain.Product{}
	errEsperado := errors.New("error de escritura")

	resultado, err := r.Update(1, "Manzana", 3, 5)

	assert.Equal(t, productEsperado, resultado)
	assert.True(t, mockStore.ReadWasCalled)
	assert.True(t, mockStore.WriteWasCalled)
	assert.Equal(t, err, errEsperado)
}

func TestUpdate(t *testing.T) {
	mockStore := store.MockStore{
		ReadWasCalled:  false,
		WriteWasCalled: false,
		ErrRead:        "",
		ErrWrite:       "",
	}
	r := NewRepository(&mockStore)

	productEsperado := domain.NewProduct(1, "Manzana", 3, 5)

	resultado, err := r.Update(1, "Manzana", 3, 5)

	assert.Equal(t, productEsperado, resultado)
	assert.True(t, mockStore.ReadWasCalled)
	assert.True(t, mockStore.WriteWasCalled)
	assert.Nil(t, err)
}

func TestUpdateNotFound(t *testing.T) {
	mockStore := store.MockStore{
		ReadWasCalled:  false,
		WriteWasCalled: false,
		ErrRead:        "",
		ErrWrite:       "",
	}
	r := NewRepository(&mockStore)

	productEsperado := domain.Product{}
	errEsperado := errors.New("no se encontro el producto de id 2")

	resultado, err := r.Update(2, "Manzana", 3, 5)

	assert.Equal(t, productEsperado, resultado)
	assert.True(t, mockStore.ReadWasCalled)
	assert.Equal(t, err, errEsperado)
}

func TestUpdateNamePriceFailRead(t *testing.T) {
	mockStore := store.MockStore{
		ReadWasCalled:  false,
		WriteWasCalled: false,
		ErrRead:        "error de lectura",
		ErrWrite:       "",
	}
	r := NewRepository(&mockStore)

	productEsperado := domain.Product{}
	errorEsperado := errors.New("error de lectura")

	resultado, err := r.UpdateNamePrice(1, "Naranja", 6)

	assert.Equal(t, productEsperado, resultado)
	assert.True(t, mockStore.ReadWasCalled)
	assert.Equal(t, errorEsperado, err)
}

func TestUpdateNamePriceFailWrite(t *testing.T) {
	mockStore := store.MockStore{
		ReadWasCalled:  false,
		WriteWasCalled: false,
		ErrRead:        "",
		ErrWrite:       "error de escritura",
	}
	r := NewRepository(&mockStore)

	productEsperado := domain.Product{}
	errorEsperado := errors.New("error de escritura")

	resultado, err := r.UpdateNamePrice(1, "Naranja", 6)

	assert.Equal(t, productEsperado, resultado)
	assert.True(t, mockStore.ReadWasCalled)
	assert.True(t, mockStore.WriteWasCalled)
	assert.Equal(t, errorEsperado, err)
}

func TestUpdateNamePrice(t *testing.T) {
	mockStore := store.MockStore{
		ReadWasCalled:  false,
		WriteWasCalled: false,
		ErrRead:        "",
		ErrWrite:       "",
	}
	r := NewRepository(&mockStore)

	productEsperado := domain.NewProduct(1, "Naranja", 6, 5)

	resultado, err := r.UpdateNamePrice(1, "Naranja", 6)

	assert.Equal(t, productEsperado, resultado)
	assert.True(t, mockStore.ReadWasCalled)
	assert.True(t, mockStore.WriteWasCalled)
	assert.Nil(t, err)
}

func TestUpdateNamePriceNotFound(t *testing.T) {
	mockStore := store.MockStore{
		ReadWasCalled:  false,
		WriteWasCalled: false,
		ErrRead:        "",
		ErrWrite:       "",
	}
	r := NewRepository(&mockStore)

	productEsperado := domain.Product{}
	errEsperado := errors.New("no se encontro el producto de id 2")

	resultado, err := r.UpdateNamePrice(2, "Naranja", 6)

	assert.Equal(t, productEsperado, resultado)
	assert.True(t, mockStore.ReadWasCalled)
	assert.Equal(t, err, errEsperado)
}

func TestLastIdFailRead(t *testing.T) {
	mockStore := store.MockStore{
		ReadWasCalled:  false,
		WriteWasCalled: false,
		ErrRead:        "error de lectura",
		ErrWrite:       "",
	}
	r := NewRepository(&mockStore)

	IdEsperado := 0
	errorEsperado := errors.New("error de lectura")

	resultado, err := r.LastId()

	assert.Equal(t, IdEsperado, resultado)
	assert.True(t, mockStore.ReadWasCalled)
	assert.Equal(t, errorEsperado, err)
}

func TestLastId(t *testing.T) {
	mockStore := store.MockStore{
		ReadWasCalled:  false,
		WriteWasCalled: false,
		ErrRead:        "",
		ErrWrite:       "",
	}
	r := NewRepository(&mockStore)

	IdEsperado := 1

	resultado, err := r.LastId()

	assert.Equal(t, IdEsperado, resultado)
	assert.True(t, mockStore.ReadWasCalled)
	assert.Nil(t, err)
}

func TestDeleteFailRead(t *testing.T) {
	mockStore := store.MockStore{
		ReadWasCalled:  false,
		WriteWasCalled: false,
		ErrRead:        "error de lectura",
		ErrWrite:       "",
	}
	r := NewRepository(&mockStore)

	errorEsperado := errors.New("error de lectura")

	err := r.Delete(1)

	assert.True(t, mockStore.ReadWasCalled)
	assert.Equal(t, errorEsperado, err)
}

func TestDeleteFailWrite(t *testing.T) {
	mockStore := store.MockStore{
		ReadWasCalled:  false,
		WriteWasCalled: false,
		ErrRead:        "",
		ErrWrite:       "error de escritura",
	}
	r := NewRepository(&mockStore)

	errorEsperado := errors.New("error de escritura")

	err := r.Delete(1)

	assert.True(t, mockStore.ReadWasCalled)
	assert.True(t, mockStore.WriteWasCalled)
	assert.Equal(t, errorEsperado, err)
}

func TestDelete(t *testing.T) {
	mockStore := store.MockStore{
		ReadWasCalled:  false,
		WriteWasCalled: false,
		ErrRead:        "",
		ErrWrite:       "",
	}
	r := NewRepository(&mockStore)

	err := r.Delete(1)

	assert.True(t, mockStore.WriteWasCalled)
	assert.Nil(t, err)

	productsEsperado := []domain.Product{}
	resultado, err := r.ReadAll()

	assert.Equal(t, productsEsperado, resultado)
	assert.True(t, mockStore.ReadWasCalled)
	assert.Nil(t, err)
}

func TestDeleteNotFound(t *testing.T) {
	mockStore := store.MockStore{
		ReadWasCalled:  false,
		WriteWasCalled: false,
		ErrRead:        "",
		ErrWrite:       "",
	}
	r := NewRepository(&mockStore)

	errorEsperado := errors.New("no se encontro el producto de id 2")

	err := r.Delete(2)

	assert.Equal(t, errorEsperado, err)
}
