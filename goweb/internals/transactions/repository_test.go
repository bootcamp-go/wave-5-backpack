package transactions

import (
	"goweb/internals/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

type StubStorage struct{}

func (st *StubStorage) Read(data interface{}) error {
	p := data.(*[]domain.Transaction)
	*p = []domain.Transaction{
		{
			Id:       1,
			Codigo:   "SAKMDLN89392",
			Moneda:   "ARG",
			Monto:    1000,
			Emisor:   "Cristian",
			Receptor: "Julian",
			Fecha:    "25-07-2022",
		},
		{
			Id:       2,
			Codigo:   "SASADLN89392",
			Moneda:   "MEX",
			Monto:    2000,
			Emisor:   "Julián",
			Receptor: "Cristian",
			Fecha:    "25-07-2022",
		},
	}
	return nil
}

func (st *StubStorage) Write(data interface{}) error {
	return nil
}

func (st *StubStorage) Ping() error {
	return nil
}

func TestRepositoryGetAll(t *testing.T) {
	//Arrange
	storage := StubStorage{}
	repository := NewRepository(&storage)
	TransactionExpected := []domain.Transaction{
		{
			Id:       1,
			Codigo:   "SAKMDLN89392",
			Moneda:   "ARG",
			Monto:    1000,
			Emisor:   "Cristian",
			Receptor: "Julian",
			Fecha:    "25-07-2022",
		},
		{
			Id:       2,
			Codigo:   "SASADLN89392",
			Moneda:   "MEX",
			Monto:    2000,
			Emisor:   "Julián",
			Receptor: "Cristian",
			Fecha:    "25-07-2022",
		},
	}
	//Act
	transaction, err := repository.GetAll()
	//Assert
	assert.Equal(t, TransactionExpected, transaction)
	assert.Nil(t, err)
}
