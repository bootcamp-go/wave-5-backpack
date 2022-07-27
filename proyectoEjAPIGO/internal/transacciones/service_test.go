package transacciones

import (
	"fmt"
	"goweb/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntegrationUpdate(t *testing.T) {
	transactionTest := []domain.Transaccion{
		{
			ID:                 1,
			Monto:              100,
			Codigo_transaccion: "asd124",
			Moneda:             "USD",
			Emisor:             "Mercado Pago",
			Receptor:           "Santander",
			Fecha_transaccion:  "2022-07-25 12:30",
		},
	}
	transactionUpdate := domain.Transaccion{
		ID:                 1,
		Monto:              100,
		Codigo_transaccion: "124asd",
		Moneda:             "UY",
		Emisor:             "Mercado Pago",
		Receptor:           "BBVA",
		Fecha_transaccion:  "2022-08-25 12:30",
	}

	db := MockRepository{
		DataMock: transactionTest,
		ReadCall: false,
	}
	//act
	repo := NewRepository(&db)
	service := NewService(repo)
	result, err := service.Update(1, "124asd", "UY", "Mercado Pago", "BBVA", "2022-08-25 12:30", 100)
	//assert
	assert.Nil(t, err)
	assert.Equal(t, transactionUpdate, result)
	assert.True(t, db.ReadCall)

}

func TestIntegrationDelete(t *testing.T) {
	//arrange
	transactionNotFound := fmt.Errorf("error al eleminar la transaccion %v", 50)
	transactionTest := []domain.Transaccion{
		{
			ID:                 1,
			Monto:              100,
			Codigo_transaccion: "124asd",
			Moneda:             "UY",
			Emisor:             "Mercado Pago",
			Receptor:           "BBVA",
			Fecha_transaccion:  "2022-08-25 12:30",
		},
	}
	db := MockRepository{
		DataMock: transactionTest,
		ReadCall: false,
	}
	//act
	repo := NewRepository(&db)
	service := NewService(repo)
	errTransactionNotFound := service.Delete(50)
	err2 := service.Delete(1)
	//assert
	assert.Equal(t, errTransactionNotFound, transactionNotFound)
	assert.Nil(t, err2)
}
