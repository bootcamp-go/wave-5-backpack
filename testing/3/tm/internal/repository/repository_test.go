package repository

import (
	"errors"
	"testing"
	"testing/3/tm/internal/domain"
	"testing/3/tm/pkg/store"

	"github.com/stretchr/testify/assert"
)

func TestUpdate(t *testing.T) {
	mockStore := store.MockStore{}
	r := NewRepository(&mockStore)

	productEsperado := domain.NewProduct(1, "After Update", 3, 5)

	resultado, err := r.Update(1, "After Update", 3, 5)

	assert.Equal(t, productEsperado, resultado)
	assert.True(t, mockStore.ReadWasCalled)
	assert.True(t, mockStore.WriteWasCalled)
	assert.Nil(t, err)
}

func TestDelete(t *testing.T) {
	mockStore := store.MockStore{}
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

func TestDeleteFail(t *testing.T) {
	mockStore := store.MockStore{}
	r := NewRepository(&mockStore)

	errorEsperado := errors.New("no se encontro el producto de id 2")

	err := r.Delete(2)

	assert.Equal(t, errorEsperado, err)
}
