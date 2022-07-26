package transactions

import (
	"fmt"
	"goweb/internal/domain"
	"goweb/pkg/store"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var timeTest, _ = time.Parse("2006-01-02T15:04:05-07:00", "2020-11-02T10:44:48+01:00")

func TestServiceUpdate(t *testing.T) {
	transactions := []domain.Transaction{
		{
			Id:              12,
			TransactionCode: "beforeUpdate",
			Currency:        "USD",
			Amount:          1232,
			Sender:          "Anonimo",
			Reciever:        "Anonimo",
			TransactionDate: timeTest,
		},
	}
	mockStore := store.Mock{
		ReadFlag: false,
		Db:       &transactions,
	}
	expectedReciever := "afterUpdate"
	expectedSender := "afterUpdate"
	expectedCurrency := "ARG"
	expectedAmount := 1500.00

	repo := NewRepository(&mockStore)
	serv := NewService(repo)

	updatedTransaction, err := serv.Update(transactions[0].Id, expectedCurrency, expectedAmount, expectedSender, expectedReciever)

	updatedTransactionDatabase := transactions[0]
	assert.Nil(t, err, fmt.Sprint("error in TestServiceUpdate: %w", err))
	assert.True(t, mockStore.ReadFlag)

	assert.Equal(t, expectedAmount, updatedTransaction.Amount, "amount not updated")
	assert.Equal(t, expectedCurrency, updatedTransaction.Currency, "currency not updated")
	assert.Equal(t, expectedSender, updatedTransaction.Sender, "sender not updated")
	assert.Equal(t, expectedReciever, updatedTransaction.Reciever, "reciever not updated")
	assert.Equal(t, updatedTransaction, updatedTransactionDatabase, "service and database data is not the same")
}

func TestServiceDelete(t *testing.T) {
	transactions := []domain.Transaction{
		{
			Id:              12,
			TransactionCode: "beforeUpdate",
			Currency:        "USD",
			Amount:          1232,
			Sender:          "Anonimo",
			Reciever:        "Anonimo",
			TransactionDate: timeTest,
		},
	}
	mockStore := store.Mock{
		ReadFlag: false,
		Db:       &transactions,
	}
	repo := NewRepository(&mockStore)
	serv := NewService(repo)

	err := serv.Delete(transactions[0].Id)

	assert.Nil(t, err, fmt.Sprint("error in TestServiceUpdate: %w", err))
	assert.True(t, mockStore.ReadFlag)
	assert.True(t, len(transactions) == 0, "element is not correctly eliminated")
}

func TestServiceDeleteNotFound(t *testing.T) {
	transactions := []domain.Transaction{
		{
			Id:              12,
			TransactionCode: "beforeUpdate",
			Currency:        "USD",
			Amount:          1232,
			Sender:          "Anonimo",
			Reciever:        "Anonimo",
			TransactionDate: timeTest,
		},
	}
	mockStore := store.Mock{
		ReadFlag: false,
		Db:       &transactions,
	}

	fakeId := 5

	repo := NewRepository(&mockStore)
	serv := NewService(repo)

	err := serv.Delete(fakeId)

	expectedError := fmt.Errorf("error: cannot be deleted id %d %w", fakeId, &NotFound{searchValue: strconv.Itoa(fakeId), fieldName: "Id"})

	assert.NotNil(t, err, "error is not raised")
	assert.Equal(t, expectedError, err, "error is not equal to expected")
}
