package transactions

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/bootcamp-go/pkg/store"
	"github.com/stretchr/testify/assert"
)

func TestUpdate(t *testing.T) {
	rT := &Transaction{
		Id:       1,
		Codigo:   "abc123",
		Moneda:   "euro",
		Emisor:   "Lucas",
		Receptor: "Sofia",
		Monto:    90,
	}

	ts := []*Transaction{
		{
			Id:       1,
			Codigo:   "abc123",
			Moneda:   "dolar",
			Emisor:   "Lucas",
			Receptor: "Luciana",
			Monto:    99,
		},
	}

	// Need Repository, StoreStub, Mock
	data, _ := json.Marshal(ts)
	mock := store.Mock{Data: data}

	storeStub := store.FileStore{Mock: &mock}
	repo := NewRepository(&storeStub)
	serv := NewService(repo)
	transaction, err := serv.Update(rT.Id, rT.Codigo, rT.Moneda, rT.Emisor, rT.Receptor, rT.Monto)

	assert.Nil(t, err)
	assert.True(t, mock.ReadInvoked)
	assert.Equal(t, rT, transaction)
}

func TestDelete(t *testing.T) {
	var (
		id          int64 = 1
		incorrectId int64 = 3
	)

	ts := []*Transaction{
		{
			Id:       1,
			Codigo:   "abc12as4",
			Moneda:   "euro",
			Emisor:   "Dante",
			Receptor: "Ivonne",
			Monto:    190000,
		},
		{
			Id:       2,
			Codigo:   "abc12as4",
			Moneda:   "dolar",
			Emisor:   "Ivonne",
			Receptor: "Juan",
			Monto:    190000,
		},
	}

	data, _ := json.Marshal(ts)
	mock := store.Mock{Data: data}
	storeStub := store.FileStore{FileName: "", Mock: &mock}

	repo := NewRepository(&storeStub)
	serv := NewService(repo)

	err := serv.Delete(id)
	assert.Nil(t, err)

	err = serv.Delete(incorrectId)
	assert.NotNil(t, err)
}

func TestDeleteError(t *testing.T) {
	var id int64 = 1
	errorExpected := errors.New("error for Delete")

	mock := store.Mock{Error: errorExpected}
	storeStub := store.FileStore{Mock: &mock}
	repo := NewRepository(&storeStub)

	err := repo.Delete(id)

	assert.NotNil(t, err)
	assert.Equal(t, errorExpected, err)
}

func TestStore(t *testing.T) {
	tsExpected := &Transaction{
		Id:       13, // lasId incrementara el id cuando recibamos enviemos datos a la función store
		Codigo:   "asfk323",
		Moneda:   "dolar",
		Emisor:   "Miguel",
		Receptor: "Luciana",
		Monto:    899,
	}

	ts := []Transaction{
		{Id: 12,
			Codigo:   "asfk323",
			Moneda:   "dolar",
			Emisor:   "Miguel",
			Receptor: "Luciana",
			Monto:    899,
		}}

	data, _ := json.Marshal(ts)
	mock := store.Mock{Data: data}
	storeStub := store.FileStore{FileName: "", Mock: &mock}
	repo := NewRepository(&storeStub)
	serv := NewService(repo)

	transaction, err := serv.Store(tsExpected.Codigo, tsExpected.Moneda, tsExpected.Emisor, tsExpected.Receptor, tsExpected.Monto)

	assert.Nil(t, err)
	assert.Equal(t, tsExpected, transaction)
}

func TestStoreError(t *testing.T) {
	var (
		codigo   string  = "abc233"
		moneda   string  = "dolar"
		emisor   string  = "Lali"
		receptor string  = "Christian"
		monto    float64 = 87.2
	)

	errorExpected := errors.New("error for Store")
	mock := store.Mock{Error: errorExpected}
	storeStub := store.FileStore{FileName: "", Mock: &mock}

	repo := NewRepository(&storeStub)
	serv := NewService(repo)
	transaction, err := serv.Store(codigo, moneda, emisor, receptor, monto)

	assert.NotNil(t, err)
	assert.Nil(t, transaction)
	assert.Equal(t, errorExpected, err)
}
