package transactions

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/bootcamp-go/pkg/store"
	"github.com/stretchr/testify/assert"
)

func TestUpdateReceptorYMonto(t *testing.T) {
	var (
		receptor string  = "Laura"
		monto    float64 = 10000
		id       int64   = 1
	)

	ts := []*Transaction{
		{
			Id:       1,
			Codigo:   "abc123",
			Moneda:   "dolar",
			Emisor:   "Juan",
			Receptor: "Sol",
			Monto:    90,
		},
	}

	data, _ := json.Marshal(ts)
	mock := store.Mock{Data: data}
	storeStub := store.FileStore{Mock: &mock}
	repo := NewRepository(&storeStub)

	transaction, err := repo.UpdateReceptorYMonto(id, receptor, monto)
	assert.Nil(t, err)
	assert.True(t, mock.ReadInvoked)
	assert.Equal(t, id, transaction.Id)
	assert.Equal(t, receptor, transaction.Receptor)
	assert.Equal(t, monto, transaction.Monto)
}

func TestUpdateReceptorYMontoError(t *testing.T) {
	var (
		receptor string  = "Laura"
		monto    float64 = 10000
		id       int64   = 1
	)
	errorExpected := errors.New("error for UpdateReceptorYMonto")
	mock := store.Mock{Error: errorExpected}

	storeStub := store.FileStore{Mock: &mock}
	repo := NewRepository(&storeStub)

	transaction, err := repo.UpdateReceptorYMonto(id, receptor, monto)
	assert.Nil(t, transaction)
	assert.NotNil(t, err)
	assert.Equal(t, errorExpected, err)
}

func TestGetAll(t *testing.T) {
	ts := []*Transaction{
		{
			Id:       1,
			Codigo:   "abc23",
			Moneda:   "peso",
			Emisor:   "Juan",
			Receptor: "María",
			Monto:    100,
		},
		{
			Id:       2,
			Codigo:   "abc24",
			Moneda:   "dolar",
			Emisor:   "Sol",
			Receptor: "María",
			Monto:    100,
		},
	}

	data, _ := json.Marshal(ts)
	mock := &store.Mock{
		Data: data,
	}

	storeStub := store.FileStore{
		FileName: "",
		Mock:     mock,
	}

	repo := NewRepository(&storeStub)
	transactions, err := repo.GetAll()

	assert.Nil(t, err)
	assert.Equal(t, ts, transactions)
}

func TestGetAllError(t *testing.T) {
	expectedError := errors.New("error for GetAll")

	mock := &store.Mock{ // no ingresamos data, por lo que deberia fallar
		Error: expectedError,
	}

	storeStub := store.FileStore{
		FileName: "",
		Mock:     mock,
	}

	repo := NewRepository(&storeStub)
	ts, err := repo.GetAll()

	assert.Nil(t, ts) // Transactions deber ser nil
	assert.Equal(t, expectedError, err)
}

func TestLastId(t *testing.T) {
	var idExpected int64 = 2

	ts := []*Transaction{
		{
			Id:       2,
			Codigo:   "123asbs",
			Moneda:   "dolar",
			Emisor:   "Juliana Martinez",
			Receptor: "Federico Lopez",
			Monto:    290,
		},
	}

	data, _ := json.Marshal(ts)
	mock := store.Mock{Data: data}
	storeStub := store.FileStore{Mock: &mock}
	repo := NewRepository(&storeStub)

	LastId, err := repo.LastId()

	assert.Nil(t, err)
	assert.Equal(t, idExpected, LastId)
}

func TestLastIdError(t *testing.T) {
	var expectedId int64 = 0
	mock := store.Mock{}

	storeStub := store.FileStore{Mock: &mock}
	repo := NewRepository(&storeStub)

	id, err := repo.LastId()
	assert.Nil(t, err)
	assert.Equal(t, expectedId, id)
}
