package transactions

import (
	"arquitectura/internal/domain"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServiceGetAll(t *testing.T) {
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

	result, err := service.GetAll()

	assert.Nil(t, err)
	assert.Equal(t, list, result)
}

func TestServiceStore(t *testing.T) {
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
		Id:          2,
		TranCode:    "AFTER",
		Currency:    "USD",
		Amount:      500,
		Transmitter: "jperez",
		Reciever:    "cmonsalve",
		TranDate:    "10-10-2022",
	}

	result, err := service.Store("AFTER", "USD", 500, "jperez", "cmonsalve", "10-10-2022")

	assert.Nil(t, err)
	assert.Equal(t, expectedResult, result)
	assert.Equal(t, 2, len(mock.Data))
}

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

func TestServiceUpdateTranCode(t *testing.T) {
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
		Currency:    "CLP",
		Amount:      500000,
		Transmitter: "cmonsalve",
		Reciever:    "jperez",
		TranDate:    "10-10-2021",
	}

	result, err := service.UpdateTranCode(1, "AFTER")
	assert.Nil(t, err)
	assert.Equal(t, expectedResult, result)
	assert.True(t, mock.ReadWasCalled)
}

func TestServiceUpdateAmount(t *testing.T) {
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
		TranCode:    "BEFORE",
		Currency:    "CLP",
		Amount:      150000,
		Transmitter: "cmonsalve",
		Reciever:    "jperez",
		TranDate:    "10-10-2021",
	}

	result, err := service.UpdateAmount(1, 150000)
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
