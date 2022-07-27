package transactions

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"goweb/internal/domain"
)

type MockStorage struct {
	transactions *[]domain.Transaction
	readWasCalled bool
}

func (ms *MockStorage) Read(data interface{}) error {
	//Des-referencio
	transactions := data.(*[]domain.Transaction)
	*transactions = *ms.transactions
	ms.readWasCalled = true
	return nil
}
func (ms *MockStorage) Write(data interface{}) error {
	return nil
}
func (ms *MockStorage) Ping() error {
	return nil
}

func TestUpdate(t *testing.T) {
	//arrange
	transactionBefore := &[]domain.Transaction{
		{
			Id:             1000,
			CodTransaction: "string",
			Currency:       "string",
			Amount:         1111,
			Sender:         "string",
			Receiver:       "string",
			DateOrder:      "string",
		},
	}

	transactionExpected := domain.Transaction{
		Id:             1000,
		CodTransaction: "string",
		Currency:       "string",
		Amount:         15000,
		Sender:         "string",
		Receiver:       "string",
		DateOrder:      "string",
	}

	infoMock := MockStorage{
		transactions: transactionBefore,
		readWasCalled: false,
	}
	repository := NewRepository(&infoMock)

	//Act
	transactionUpdated, err := repository.UpdateAmount(1000, 15000)

	//Assert
	assert.Nil(t, err)
	assert.Equal(t, transactionExpected, transactionUpdated)
	assert.True(t, infoMock.readWasCalled, "function read not was called")
}