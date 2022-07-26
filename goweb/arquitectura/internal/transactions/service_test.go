package transactions

import (
	"arquitectura/internal/domain"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServiceUpdate(t *testing.T) {
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
	service := NewService(repo)
	expectedResult := domain.Transaction{
		Id:          1,
		TranCode:    "AFTER",
		Currency:    "USD",
		Amount:      500,
		Transmitter: "jperez",
		Reciever:    "cmonsalve",
		TranDate:    "10-10-2022",
	}

	result, err := service.Update(1, "AFTER", "USD", 500, "jperez", "cmonsalve", "10-10-2022")
	assert.Nil(t, err)
	assert.Equal(t, expectedResult, result)
	assert.True(t, mock.ReadWasCalled)
}

func TestServiceDelete(t *testing.T) {
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
	service := NewService(repo)

	err := service.Delete(1)

	assert.Nil(t, err)
	assert.Equal(t, 0, len(mock.Data))
	assert.True(t, mock.ReadWasCalled)
}

func TestServiceDeleteInexistent(t *testing.T) {
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
	service := NewService(repo)
	expectedError := fmt.Errorf("transaction with id %d doesn`t exists en database", 2)

	err := service.Delete(2)

	assert.ErrorContains(t, err, expectedError.Error())
}
