package transactions

import (
	"testing"

	"github.com/bootcamp-go/wave-5-backpack/tree/lopez_cristian/goweb/internal/models"
	"github.com/stretchr/testify/assert"
)

type StubStorage struct{}

func (s *StubStorage) Read(data interface{}) error {
	p := data.(*[]models.Transaction)
	*p = []models.Transaction{
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
			Monto:    1000.5,
			Cod:      "aaa111",
			Moneda:   "ARS",
			Emisor:   "Mercado Pago",
			Receptor: "BBVA",
			Fecha:    "2020-25-07",
		},
	}

	return nil
}

func (s *StubStorage) Write(data interface{}) error {
	return nil
}

func TestRepositoryGetAll(t *testing.T) {
	//Arrange
	storage := &StubStorage{}
	repository := NewRepository(storage)
	transactionsExpected := []models.Transaction{
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
			Monto:    1000.5,
			Cod:      "aaa111",
			Moneda:   "ARS",
			Emisor:   "Mercado Pago",
			Receptor: "BBVA",
			Fecha:    "2020-25-07",
		},
	}
	//Act
	transactions, err := repository.GetAll()

	//Assert
	assert.Equal(t, transactionsExpected, transactions)
	assert.Nil(t, err)
}
