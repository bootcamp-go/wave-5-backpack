package transactions

import (
	"goweb/internal/domain"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var timeNow, _ = time.Parse("2006-01-02T15:04:05-07:00", "2020-11-02T10:44:48+01:00")

type Store struct {
	ReadFlag bool
	Db       *[]domain.Transaction
}

func (s *Store) Ping() error {
	return nil
}
func (s *Store) Read(data interface{}) error {
	s.ReadFlag = true
	listProducts := data.(*[]domain.Transaction)
	*listProducts = *s.Db
	return nil
}
func (s *Store) Write(data interface{}) error {
	listProducts := data.([]domain.Transaction)
	*s.Db = listProducts
	return nil
}

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
	mockStore := Store{
		ReadFlag: false,
		Db:       &dataTransaction,
	}
	repo := NewRepository(&mockStore)

	response, err := repo.GetAll()
	assert.Nil(t, err, "error debe ser nulo")
	assert.Equal(t, dataTransaction, response, "deben ser iguales")
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
	mockStore := Store{
		ReadFlag: false,
		Db:       &transactions,
	}
	expectedCurrency := "ARG"
	expectedAmount := 1500.00
	repo := NewRepository(&mockStore)

	repo.UpdateCurrencyAndAmount(transactions[0].Id, expectedCurrency, expectedAmount)

	updatedTransaction := transactions[0]
	assert.True(t, mockStore.ReadFlag)
	assert.Equal(t, expectedAmount, updatedTransaction.Amount, "amount not updated")
	assert.Equal(t, expectedCurrency, updatedTransaction.Currency, "currency not updated")
}
