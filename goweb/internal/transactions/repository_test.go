package transactions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type StubTransactions struct{}

func (st *StubTransactions) Read(data interface{}) error {
	transaction := data.(*[]*Transactions)
	t1 := &Transactions{
		Id:              1,
		TransactionCode: "TRM01",
		TypeCurrency:    "MXN",
		Amount:          67800,
		Transmitter:     "BANREGIO",
		Receiver:        "BBVA",
		Date:            "17/07/2022",
		Completed:       true,
	}
	t2 := &Transactions{
		Id:              2,
		TransactionCode: "TRM02",
		TypeCurrency:    "MXN",
		Amount:          12500,
		Transmitter:     "SCOTIABANK",
		Receiver:        "HSBC",
		Date:            "20/07/2022",
		Completed:       true,
	}

	*transaction = append(*transaction, t1)
	*transaction = append(*transaction, t2)

	return nil
}

func (st *StubTransactions) Write(data interface{}) error {
	return nil
}

func TestReadAll(t *testing.T) {
	stub := &StubTransactions{}
	repo := NewRepository(stub)
	expected := []*Transactions{
		{
			Id:              1,
			TransactionCode: "TRM01",
			TypeCurrency:    "MXN",
			Amount:          67800,
			Transmitter:     "BANREGIO",
			Receiver:        "BBVA",
			Date:            "17/07/2022",
			Completed:       true,
		},
		{
			Id:              2,
			TransactionCode: "TRM02",
			TypeCurrency:    "MXN",
			Amount:          12500,
			Transmitter:     "SCOTIABANK",
			Receiver:        "HSBC",
			Date:            "20/07/2022",
			Completed:       true,
		},
	}

	tsGetAll, err := repo.GetAll()
	assert.Nil(t, err)
	assert.Equal(t, expected, tsGetAll)

}

type MockStore struct {
	Invoked bool
	Data    []*Transactions
}

func (ms MockStore) Read(data interface{}) error {
	ms.Invoked = true
	read := data.(*[]*Transactions)
	*read = ms.Data
	return nil
}

func (ms *MockStore) Write(data interface{}) error {
	return nil
}

func TestUpdateName(t *testing.T) {
	id, transactionsCode := 1, "Before_Update"
	transaction := []*Transactions{{
		Id:              1,
		TransactionCode: "Before_Update",
		TypeCurrency:    "MXN",
		Amount:          67800,
		Transmitter:     "BANREGIO",
		Receiver:        "BBVA",
		Date:            "17/07/2022",
		Completed:       true,
	}}

	mock := MockStore{Data: transaction}

	repo := NewRepository(&mock)

	update, err := repo.UpdateTransmitter(int64(id), transactionsCode)

	assert.Nil(t, err)
	assert.Equal(t, transactionsCode, update.TransactionCode)
	assert.True(t, true, mock.Invoked)
}
