package transactions

import (
	"goweb/internal/domain"
	"goweb/pkg/store"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var timeNow, _ = time.Parse("2006-01-02T15:04:05-07:00", "2020-11-02T10:44:48+01:00")

func TestGetAll(t *testing.T) {
	var dataTransaction = []domain.Transaction{
		{
			Id:              12,
			TransactionCode: "beforeUpdate",
			Currency:        "USD",
			Amount:          1232,
			Sender:          "Anonimo",
			Reciever:        "Anonimo",
			TransactionDate: timeNow,
		},
		{
			Id:              1,
			TransactionCode: "kdas23kda",
			Currency:        "CLP",
			Amount:          120000,
			Sender:          "Anonimo 2",
			Reciever:        "Anonimo 3",
			TransactionDate: timeNow,
		},
	}
	mockStore := store.Mock{
		ReadFlag: false,
		Db:       &dataTransaction,
	}
	repo := NewRepository(&mockStore)

	response, err := repo.GetAll()
	assert.Nil(t, err, "error debe ser nulo")
	assert.Equal(t, dataTransaction, response, "deben ser iguales")
}

func TestGetAllFail(t *testing.T) {
	var dataTransaction = []domain.Transaction{}
	mockStore := store.Mock{
		FailRead: true,
		ReadFlag: false,
		Db:       &dataTransaction,
	}
	repo := NewRepository(&mockStore)

	_, err := repo.GetAll()
	assert.ErrorContains(t, err, "error: at read")
}

func TestStore(t *testing.T) {
	var dataTransaction = []domain.Transaction{}
	mockStore := store.Mock{
		ReadFlag: false,
		Db:       &dataTransaction,
	}
	repo := NewRepository(&mockStore)

	expectedTransaction := domain.Transaction{
		Id:              1,
		TransactionCode: "kdas23kda",
		Currency:        "CLP",
		Amount:          120000,
		Sender:          "Anonimo 2",
		Reciever:        "Anonimo 3",
		TransactionDate: timeNow,
	}
	_, err := repo.Store(expectedTransaction.Id,
		expectedTransaction.TransactionCode,
		expectedTransaction.Currency,
		expectedTransaction.Amount,
		expectedTransaction.Sender,
		expectedTransaction.Reciever,
		expectedTransaction.TransactionDate)
	assert.Nil(t, err, "error at store")
	assert.True(t, len(dataTransaction) == 1, "not writed")
	assert.Equal(t, expectedTransaction, dataTransaction[0], "not equal")
}

func TestFailWriteStore(t *testing.T) {
	var dataTransaction = []domain.Transaction{}
	mockStore := store.Mock{
		FailWrite: true,
		Db:        &dataTransaction,
	}
	repo := NewRepository(&mockStore)

	expectedTransaction := domain.Transaction{
		Id:              1,
		TransactionCode: "kdas23kda",
		Currency:        "CLP",
		Amount:          120000,
		Sender:          "Anonimo 2",
		Reciever:        "Anonimo 3",
		TransactionDate: timeNow,
	}
	_, err := repo.Store(expectedTransaction.Id,
		expectedTransaction.TransactionCode,
		expectedTransaction.Currency,
		expectedTransaction.Amount,
		expectedTransaction.Sender,
		expectedTransaction.Reciever,
		expectedTransaction.TransactionDate)

	assert.ErrorContains(t, err, "error: at write")
}

func TestFailReadStore(t *testing.T) {
	var dataTransaction = []domain.Transaction{}
	mockStore := store.Mock{
		FailRead: true,
		Db:       &dataTransaction,
	}
	repo := NewRepository(&mockStore)

	expectedTransaction := domain.Transaction{
		Id:              1,
		TransactionCode: "kdas23kda",
		Currency:        "CLP",
		Amount:          120000,
		Sender:          "Anonimo 2",
		Reciever:        "Anonimo 3",
		TransactionDate: timeNow,
	}
	_, err := repo.Store(expectedTransaction.Id,
		expectedTransaction.TransactionCode,
		expectedTransaction.Currency,
		expectedTransaction.Amount,
		expectedTransaction.Sender,
		expectedTransaction.Reciever,
		expectedTransaction.TransactionDate)

	assert.ErrorContains(t, err, "error: at read")
}

func TestUpdate(t *testing.T) {
	transactions := []domain.Transaction{
		{
			Id:              12,
			TransactionCode: "beforeUpdate",
			Currency:        "USD",
			Amount:          1232,
			Sender:          "Anonimo",
			Reciever:        "Anonimo",
			TransactionDate: timeNow,
		},
	}
	mockStore := store.Mock{
		ReadFlag: false,
		Db:       &transactions,
	}
	expectedCurrency := "ARG"
	expectedAmount := 1500.00
	repo := NewRepository(&mockStore)

	_, err := repo.UpdateCurrencyAndAmount(transactions[0].Id, expectedCurrency, expectedAmount)

	updatedTransaction := transactions[0]
	assert.Nil(t, err, "error debe ser nulo")
	assert.True(t, mockStore.ReadFlag)
	assert.Equal(t, expectedAmount, updatedTransaction.Amount, "amount not updated")
	assert.Equal(t, expectedCurrency, updatedTransaction.Currency, "currency not updated")
}

func TestUpdateFail(t *testing.T) {
	transactions := []domain.Transaction{
		{
			Id:              12,
			TransactionCode: "beforeUpdate",
			Currency:        "USD",
			Amount:          1232,
			Sender:          "Anonimo",
			Reciever:        "Anonimo",
			TransactionDate: timeNow,
		},
	}
	mockStore := store.Mock{
		FailRead: true,
		ReadFlag: false,
		Db:       &transactions,
	}
	expectedCurrency := "ARG"
	expectedAmount := 1500.00
	repo := NewRepository(&mockStore)

	_, err := repo.UpdateCurrencyAndAmount(transactions[0].Id, expectedCurrency, expectedAmount)

	assert.ErrorContains(t, err, "error: at read")

	mockStore.FailRead = false
	mockStore.FailWrite = true

	_, err = repo.UpdateCurrencyAndAmount(12, expectedCurrency, expectedAmount)
	assert.ErrorContains(t, err, "error: at write")

	mockStore.FailWrite = false

	_, err = repo.UpdateCurrencyAndAmount(15, expectedCurrency, expectedAmount)
	var notFoundErr *NotFound
	assert.ErrorAs(t, err, &notFoundErr)
}

func TestGetById(t *testing.T) {
	transactions := []domain.Transaction{
		{
			Id:              12,
			TransactionCode: "beforeUpdate",
			Currency:        "USD",
			Amount:          1232,
			Sender:          "Anonimo",
			Reciever:        "Anonimo",
			TransactionDate: timeNow,
		},
	}
	mockStore := store.Mock{
		ReadFlag: false,
		Db:       &transactions,
	}
	repo := NewRepository(&mockStore)
	data, err := repo.GetById(transactions[0].Id)
	assert.Nil(t, err, "error debe ser nulo")
	assert.Equal(t, transactions[0], data, "not equal")
}

func TestGetByIdFail(t *testing.T) {
	transactions := []domain.Transaction{}
	mockStore := store.Mock{
		FailRead: true,
		ReadFlag: false,
		Db:       &transactions,
	}
	repo := NewRepository(&mockStore)
	_, err := repo.GetById(10)

	assert.ErrorContains(t, err, "error: at read")

	mockStore.FailRead = false

	_, err = repo.GetById(10)
	var notFoundErr *NotFound
	assert.ErrorAs(t, err, &notFoundErr)
}

func TestLastId(t *testing.T) {
	transactions := []domain.Transaction{
		{
			Id:              12,
			TransactionCode: "beforeUpdate",
			Currency:        "USD",
			Amount:          1232,
			Sender:          "Anonimo",
			Reciever:        "Anonimo",
			TransactionDate: timeNow,
		},
	}
	mockStore := store.Mock{
		ReadFlag: false,
		Db:       &transactions,
	}
	repo := NewRepository(&mockStore)
	id, err := repo.lastId()
	assert.Nil(t, err, "error debe ser nulo")
	assert.Equal(t, transactions[0].Id, id, "not equal id")
}

func TestLastIdFail(t *testing.T) {
	transactions := []domain.Transaction{}
	mockStore := store.Mock{
		FailRead: true,
		ReadFlag: false,
		Db:       &transactions,
	}
	repo := NewRepository(&mockStore)
	_, err := repo.lastId()

	assert.ErrorContains(t, err, "error: at read")
}

func TestDeleteFail(t *testing.T) {
	transactions := []domain.Transaction{{
		Id:              12,
		TransactionCode: "beforeUpdate",
		Currency:        "USD",
		Amount:          1232,
		Sender:          "Anonimo",
		Reciever:        "Anonimo",
		TransactionDate: timeNow,
	}}
	mockStore := store.Mock{
		FailRead:  true,
		FailWrite: true,
		ReadFlag:  false,
		Db:        &transactions,
	}
	repo := NewRepository(&mockStore)
	err := repo.Delete(1)

	assert.ErrorContains(t, err, "error: at read")
	mockStore.FailRead = false

	err = repo.Delete(12)
	assert.ErrorContains(t, err, "error: at write")

}
