package service

import (
	"testing"
	"testing/3/tm/internal/domain"
	"testing/3/tm/internal/repository"
	"testing/3/tm/pkg/store"

	"github.com/stretchr/testify/assert"
)

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
	resultado, err := s.ReadAll()

	assert.Equal(t, productsEsperado, resultado)
	assert.True(t, mockStore.ReadWasCalled)
	assert.NotNil(t, err)
}

func TestDeleteFail(t *testing.T) {
	mockStore := store.MockStore{}
	r := repository.NewRepository(&mockStore)
	s := NewService(r)

	err := s.Delete(2)

	assert.Error(t, err)
}
