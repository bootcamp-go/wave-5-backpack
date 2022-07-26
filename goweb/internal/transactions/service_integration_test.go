package transactions

import (
	"testing"

	"github.com/bootcamp-go/wave-5-backpack/tree/lopez_cristian/goweb/internal/models"
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
