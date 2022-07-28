package transactions

import (
	"errors"
	"goweb/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntregationGetAll(t *testing.T) {

	database := []domain.Transaction{
		{
			Id:          1,
			TranCode:    "ABC1234",
			Currency:    "USD",
			Amount:      200.0,
			Transmitter: "MERCADOPAGO",
			Receiver:    "JOSE",
			TranDate:    "05-07-22",
		},
		{
			Id:          2,
			TranCode:    "XYZ1234",
			Currency:    "CLP",
			Amount:      1500000.0,
			Transmitter: "MELI",
			Receiver:    "JUAN",
			TranDate:    "12-02-22",
		},
	}

	mockStorage := MockStore{
		dataMock: database,
	}

	// act
	repo := NewRepository(&mockStorage)
	serv := NewService(repo)
	result, err := serv.GetAll()

	// assert
	assert.Nil(t, err)
	assert.True(t, mockStorage.readCalled)
	assert.Equal(t, mockStorage.dataMock, result)

}

func TestServiceIntegrationUpdate(t *testing.T) {
	// arrange

	database := []domain.Transaction{
		{
			Id:          1,
			TranCode:    "ABC1234",
			Currency:    "USD",
			Amount:      200.0,
			Transmitter: "MERCADOPAGO",
			Receiver:    "JOSE",
			TranDate:    "05-07-22",
		},
		{
			Id:          2,
			TranCode:    "XYZ1234",
			Currency:    "CLP",
			Amount:      1500000.0,
			Transmitter: "MELI",
			Receiver:    "JUAN",
			TranDate:    "12-02-22",
		},
		{
			Id:          3,
			TranCode:    "BBB456",
			Currency:    "USD",
			Amount:      500.00,
			Transmitter: "BANCO",
			Receiver:    "JUAN",
			TranDate:    "26-07-22",
		},
	}

	mockStorage := MockStore{
		dataMock: database,
		errWrite: "",
		errRead:  "",
	}

	// act
	repo := NewRepository(&mockStorage)
	serv := NewService(repo)
	result, err := serv.Update(3, "BBB456", "USD", 500.00, "BANCO", "JUAN", "26-07-22")

	// assert
	assert.Nil(t, err)
	assert.True(t, mockStorage.readCalled)
	assert.Equal(t, mockStorage.dataMock[2], result)

}

func TestServiceIntegrationDelete(t *testing.T) {
	// arrange

	database := []domain.Transaction{
		{
			Id:          1,
			TranCode:    "ABC1234",
			Currency:    "USD",
			Amount:      200.0,
			Transmitter: "MERCADOPAGO",
			Receiver:    "JOSE",
			TranDate:    "05-07-22",
		},
		{
			Id:          2,
			TranCode:    "XYZ1234",
			Currency:    "CLP",
			Amount:      1500000.0,
			Transmitter: "MELI",
			Receiver:    "JUAN",
			TranDate:    "12-02-22",
		},
		{
			Id:          3,
			TranCode:    "BBB456",
			Currency:    "USD",
			Amount:      500.00,
			Transmitter: "BANCO",
			Receiver:    "JUAN",
			TranDate:    "26-07-22",
		},
	}

	mockStorage := MockStore{
		dataMock: database,
		errWrite: "",
		errRead:  "",
	}

	// act
	repo := NewRepository(&mockStorage)
	serv := NewService(repo)
	err := serv.Delete(3)

	// assert
	//assert.Equal(t, 2, len(mockStorage.dataMock))
	assert.Nil(t, err)

}

func TestServiceIntegrationDeleteFail(t *testing.T) {
	// arrange
	expectedError := errors.New("Producto 4 no encontrado")

	mockStorage := MockStore{
		dataMock: nil,
		errWrite: "",
		errRead:  "",
	}

	// act
	repo := NewRepository(&mockStorage)
	serv := NewService(repo)
	err := serv.Delete(4)

	// assert

	assert.ErrorContains(t, err, expectedError.Error())

}
