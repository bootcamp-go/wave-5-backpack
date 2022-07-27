package service

import (
	"errors"
	"testing"
	"testing/4/tm/test_funcional/internal/domain"
	"testing/4/tm/test_funcional/internal/repository"
	"testing/4/tm/test_funcional/pkg/store"

	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	mockStore := store.MockStore{}
	r := repository.NewRepository(&mockStore)
	s := NewService(r)

	productEsperado := domain.NewProduct(1, "Banana", 1.5, 2)

	resultado, err := s.Create("Banana", 1.5, 2)

	assert.Equal(t, productEsperado, resultado)
	assert.True(t, mockStore.WriteWasCalled)
	assert.Nil(t, err)
}

func TestUpdate(t *testing.T) {
	mockStore := store.MockStore{}
	r := repository.NewRepository(&mockStore)
	s := NewService(r)

	productEsperado := domain.NewProduct(1, "After Update", 3, 5)

	resultado, err := s.Update(1, "After Update", 3, 5)

	assert.Equal(t, productEsperado, resultado)
	assert.True(t, mockStore.ReadWasCalled)
	assert.True(t, mockStore.WriteWasCalled)
	assert.Nil(t, err)
}

func TestDelete(t *testing.T) {
	mockStore := store.MockStore{}
	r := repository.NewRepository(&mockStore)
	s := NewService(r)

	err := s.Delete(1)

	assert.True(t, mockStore.WriteWasCalled)
	assert.Nil(t, err)

	productsEsperado := []domain.Product{}
	errorEsperado := errors.New("no se han encontrado productos en el listado")
	resultado, err := s.ReadAll()

	assert.Equal(t, productsEsperado, resultado)
	assert.True(t, mockStore.ReadWasCalled)
	assert.Equal(t, errorEsperado, err)
}

func TestDeleteFail(t *testing.T) {
	mockStore := store.MockStore{}
	r := repository.NewRepository(&mockStore)
	s := NewService(r)

	errorEsperado := errors.New("no se encontro el producto de id 2")

	err := s.Delete(2)

	assert.Equal(t, err, errorEsperado)
}
