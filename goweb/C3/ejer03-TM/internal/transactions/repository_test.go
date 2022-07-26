package transactions

import (
	"ejer02-TT/internal/domain"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockStorage struct {
	dataMock   []domain.Transaction
	errWrite   string
	errRead    string
	ReadCalled bool
}

func (fs *MockStorage) Read(data interface{}) error {
	fs.ReadCalled = true
	if fs.errRead != "" {
		return fmt.Errorf(fs.errRead)
	}
	a := data.(*[]domain.Transaction)
	*a = fs.dataMock
	return nil

}

func (fs *MockStorage) Write(data interface{}) error {
	return nil
}

func (fs *MockStorage) Ping() error {
	return nil
}

func TestStore(t *testing.T) {

	database := []domain.Transaction{
		{
			Id:          1,
			TranCode:    "tranCode",
			Currency:    "moneda",
			Amount:      12.5,
			Transmitter: "transmitter",
			Reciever:    "reciever",
			TranDate:    "tranDate",
		},
		{
			Id:          2,
			TranCode:    "tranCode2",
			Currency:    "moneda",
			Amount:      10.5,
			Transmitter: "transmitter",
			Reciever:    "reciever",
			TranDate:    "tranDate",
		},
	}

	newTransaction := domain.Transaction{

		Id:          3,
		TranCode:    "tranCode3",
		Currency:    "moneda",
		Amount:      11.5,
		Transmitter: "transmitter",
		Reciever:    "reciever",
		TranDate:    "tranDate",
	}

	mockStorage := MockStorage{
		dataMock: database,
	}
	// act
	repo := NewRepository(&mockStorage)
	result, err := repo.Store(
		newTransaction.Id,
		newTransaction.TranCode,
		newTransaction.Currency,
		newTransaction.Amount,
		newTransaction.Transmitter,
		newTransaction.Reciever,
		newTransaction.TranDate,
	)
	// assert
	fmt.Printf("\n%+v", mockStorage.dataMock)
	assert.Nil(t, err)
	assert.Equal(t, newTransaction, result)
	//assert.Equal(t, mockStorage.dataMock[len(mockStorage.dataMock)-1], newTransaction)
}

func TestGetAll(t *testing.T) {

	database := []domain.Transaction{
		{
			Id:          1,
			TranCode:    "tranCode",
			Currency:    "moneda",
			Amount:      12.5,
			Transmitter: "transmitter",
			Reciever:    "reciever",
			TranDate:    "tranDate",
		},
		{
			Id:          2,
			TranCode:    "tranCode2",
			Currency:    "moneda",
			Amount:      10.5,
			Transmitter: "transmitter",
			Reciever:    "reciever",
			TranDate:    "tranDate",
		},
	}

	mockStorage := MockStorage{
		dataMock: database,
		errWrite: "",
		errRead:  "",
	}
	// act
	repo := NewRepository(&mockStorage)
	result, err := repo.GetAll()
	// assert
	assert.Nil(t, err)
	assert.Equal(t, mockStorage.dataMock, result)
}

func TestUpdate(t *testing.T) {

	database := []domain.Transaction{
		{
			Id:          1,
			TranCode:    "tranCode",
			Currency:    "moneda",
			Amount:      12.5,
			Transmitter: "transmitter",
			Reciever:    "reciever",
			TranDate:    "tranDate",
		},
		{
			Id:          2,
			TranCode:    "tranCode2",
			Currency:    "moneda",
			Amount:      10.5,
			Transmitter: "transmitter",
			Reciever:    "reciever",
			TranDate:    "tranDate",
		},
	}

	mockStorage := MockStorage{
		dataMock: database,
		errWrite: "",
		errRead:  "",
	}
	// act
	repo := NewRepository(&mockStorage)

	expected := domain.Transaction{Id: 1, TranCode: "cambio", Currency: "moneda", Amount: 10.5, Transmitter: "transmitter", Reciever: "reciever", TranDate: "tranDate"}

	a, err := repo.UpdateCodeAndAmount(1, "cambio", 10.5)

	assert.Nil(t, err)
	assert.Equal(t, expected, a)
	//assert.True(t, stub.ReadCalled)
}
