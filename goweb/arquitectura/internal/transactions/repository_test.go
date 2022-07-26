package transactions

import (
	"arquitectura/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

type StoreMock struct {
	ReadWasCalled bool
	Data          []domain.Transaction
}

func (sm *StoreMock) Read(data interface{}) error {
	sm.ReadWasCalled = true
	a := data.(*[]domain.Transaction)
	*a = append(*a, sm.Data...)
	return nil
}

func (sm *StoreMock) Write(data interface{}) error {
	a := data.([]domain.Transaction)
	sm.Data = a
	return nil
}

func (sm StoreMock) Ping() error {
	return nil
}

func TestRepoGetAll(t *testing.T) {
	expectedResult := []domain.Transaction{
		{
			Id:          1,
			TranCode:    "BEFORE",
			Currency:    "CLP",
			Amount:      500000,
			Transmitter: "cmonsalve",
			Reciever:    "jperez",
			TranDate:    "10-10-2021",
		},
		{
			Id:          2,
			TranCode:    "BEFORE",
			Currency:    "USD",
			Amount:      500,
			Transmitter: "jperez",
			Reciever:    "ctorres",
			TranDate:    "20-07-2022",
		},
	}
	mock := StoreMock{false, expectedResult}
	repo := NewRepository(&mock)

	result, err := repo.GetAll()
	assert.Nil(t, err)
	assert.Equal(t, expectedResult, result)
}

func TestRepoUpdateTranCode(t *testing.T) {
	list := []domain.Transaction{
		{
			Id:          1,
			TranCode:    "BEFORE",
			Currency:    "CLP",
			Amount:      500000,
			Transmitter: "cmonsalve",
			Reciever:    "jperez",
			TranDate:    "10-10-2021",
		},
	}
	mock := StoreMock{false, list}
	repo := NewRepository(&mock)
	expectedResult := domain.Transaction{
		Id:          1,
		TranCode:    "AFTER",
		Currency:    "CLP",
		Amount:      500000,
		Transmitter: "cmonsalve",
		Reciever:    "jperez",
		TranDate:    "10-10-2021",
	}
	result, err := repo.UpdateTranCode(1, "AFTER")
	assert.Nil(t, err)
	assert.Equal(t, expectedResult, result)
	assert.True(t, mock.ReadWasCalled)
}
