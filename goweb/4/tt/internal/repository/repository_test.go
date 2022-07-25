package repository

import (
	"goweb/4/tt/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadALl(t *testing.T) {
	stubStore := StubStore{}
	r := NewRepository(&stubStore)

	var productsEsperado []domain.Product
	product1 := domain.NewProduct(1, "Banana", 1.5, 2)
	product2 := domain.NewProduct(2, "Manzana", 1.25, 5)
	productsEsperado = append(productsEsperado, product1, product2)

	resultado, err := r.ReadAll()

	assert.Equal(t, productsEsperado, resultado)
	assert.Nil(t, err)
}

func TestUpdateNamePrice(t *testing.T) {
	mockStore := MockStore{}
	r := NewRepository(&mockStore)

	productEsperado := domain.NewProduct(1, "After Update", 1.5, 2)

	resultado, err := r.UpdateNamePrice(1, "After Update", 1.5)

	assert.Equal(t, productEsperado, resultado)
	assert.Nil(t, err)
	assert.True(t, mockStore.ReadWasCalled)
	assert.True(t, mockStore.WriteWasCalled)

}
