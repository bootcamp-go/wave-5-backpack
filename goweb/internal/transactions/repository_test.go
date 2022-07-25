package transactions

import (
	"encoding/json"
	"goweb/internal/domain"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var timeNow, _ = time.Parse("2006-01-02T15:04:05-07:00", "2020-11-02T10:44:48+01:00")

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

type Store struct {
	ReadFlag bool
}

func (s *Store) Ping() error {
	return nil
}
func (s *Store) Read(data interface{}) error {
	s.ReadFlag = true
	bytes, err := json.Marshal(dataTransaction)
	json.Unmarshal(bytes, data)
	return err
}
func (s *Store) Write(data interface{}) error {

	transactionMarshal, err := json.Marshal(data)
	json.Unmarshal(transactionMarshal, &dataTransaction)
	return err
}

func TestGetAll(t *testing.T) {
	mockStore := Store{
		ReadFlag: false,
	}
	repo := NewRepository(&mockStore)

	response, err := repo.GetAll()
	assert.Nil(t, err, "error debe ser nulo")
	assert.Equal(t, dataTransaction, response, "deben ser iguales")
}

func TestUpdate(t *testing.T) {
	mockStore := Store{
		ReadFlag: false,
	}
	expectedCurrency := "ARG"
	expectedAmount := 1500.00
	repo := NewRepository(&mockStore)

	repo.UpdateCurrencyAndAmount(dataTransaction[0].Id, expectedCurrency, expectedAmount)

	updatedTransaction := dataTransaction[0]
	assert.True(t, mockStore.ReadFlag)
	assert.Equal(t, expectedAmount, updatedTransaction.Amount, "amount not updated")
	assert.Equal(t, expectedCurrency, updatedTransaction.Currency, "currency not updated")
}
