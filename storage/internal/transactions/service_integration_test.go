package transactions

import (
	"fmt"
	"testing"

	"github.com/bootcamp-go/wave-5-backpack/tree/lopez_cristian/storage/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestIntegrationUpdate(t *testing.T) {
	//Arrange
	data := []models.Transaction{
		{
			ID:       1,
			Monto:    500.5,
			Cod:      "aaa111",
			Moneda:   "ARS",
			Emisor:   "Mercado Pago",
			Receptor: "BBVA",
			Fecha:    "2022-07-25 12:00",
		},
	}
	transactionExpected := models.Transaction{
		ID:       1,
		Monto:    1000,
		Cod:      "aaa112",
		Moneda:   "USD",
		Emisor:   "BBVA",
		Receptor: "Mercado Pago",
		Fecha:    "2022-07-25 12:00",
	}

	// Act
	storage := MockStorage{data, false, false, "", ""}
	repo := NewRepository(&storage)
	service := NewService(repo)
	transactionUpdated, err := service.Update(1, 1000, "aaa112", "USD", "BBVA", "Mercado Pago")

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, transactionExpected, transactionUpdated)
	assert.True(t, storage.ReadCalled)
	assert.True(t, storage.WriteCalled)
}

func TestIntegrationDelete(t *testing.T) {
	//Arrange
	data := []models.Transaction{
		{
			ID:       1,
			Monto:    1000.5,
			Cod:      "aaa111",
			Moneda:   "ARS",
			Emisor:   "Mercado Pago",
			Receptor: "BBVA",
			Fecha:    "2020-25-07",
		},
		{
			ID:       2,
			Monto:    500,
			Cod:      "aaa112",
			Moneda:   "ARS",
			Emisor:   "BBVA",
			Receptor: "Mercado Pago",
			Fecha:    "2020-25-07",
		},
	}

	//Act
	storage := MockStorage{Data: data}
	repo := NewRepository(&storage)
	service := NewService(repo)
	id, err := service.Delete(2)

	//Assert
	assert.True(t, storage.ReadCalled)
	assert.True(t, storage.WriteCalled)
	assert.Equal(t, 2, id)
	assert.Len(t, storage.Data, 1)
	assert.Nil(t, err)
}

func TestIntegrationDeleteNotFound(t *testing.T) {
	//Arrange
	storage := MockStorage{}
	errExpected := fmt.Errorf("error: ID %v no existe\n", 1)

	//Act
	repo := NewRepository(&storage)
	service := NewService(repo)
	id, err := service.Delete(1)

	//Assert
	if assert.Errorf(t, errExpected, "1") {
		assert.Equal(t, errExpected, err)
	}
	assert.Equal(t, 0, id)
	assert.Len(t, storage.Data, 0)
	assert.True(t, storage.ReadCalled)
	assert.False(t, storage.WriteCalled)
}
