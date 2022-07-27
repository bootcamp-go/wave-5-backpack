package transactions

import (
	"ejer02-TT/internal/domain"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntegracionGetAll(t *testing.T) {
	// arrange
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
		errRead:  "",
		errWrite: "",
	}

	// act
	repo := NewRepository(&mockStorage)
	service := NewService(repo)
	res, err := service.GetAll()

	// assert

	assert.Nil(t, err)
	assert.Equal(t, mockStorage.dataMock, res)
}

func TestIntegracionGetAllFail(t *testing.T) {
	// arrange
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
		errRead:  "can't read database",
		errWrite: "",
	}

	// act
	repo := NewRepository(&mockStorage)
	service := NewService(repo)
	_, err := service.GetAll()
	errExpected := fmt.Errorf("can't read database")
	// assert

	assert.NotNil(t, err)
	assert.Equal(t, errExpected, err)
	//assert.Equal(t, mockStorage.dataMock, res)
}

func TestIntegracionStore(t *testing.T) {
	newTransaction := domain.Transaction{

		TranCode:    "tranCode3",
		Currency:    "moneda",
		Amount:      11.5,
		Transmitter: "transmitter",
		Reciever:    "reciever",
		TranDate:    "tranDate",
	}

	mockStorage := MockStorage{
		dataMock: []domain.Transaction{},
		errRead:  "",
		errWrite: "",
	}

	// act
	repo := NewRepository(&mockStorage)
	service := NewService(repo)
	result, err := service.Store(
		newTransaction.TranCode,
		newTransaction.Currency,
		newTransaction.Amount,
		newTransaction.Transmitter,
		newTransaction.Reciever,
		newTransaction.TranDate,
	)

	// assert

	assert.Nil(t, err)
	assert.Equal(t, mockStorage.dataMock[0], result)
	assert.Equal(t, mockStorage.dataMock[0].Id, result.Id)

}

func TestIntegracionStoreFail(t *testing.T) {
	newTransaction := domain.Transaction{

		TranCode:    "tranCode3",
		Currency:    "moneda",
		Amount:      11.5,
		Transmitter: "transmitter",
		Reciever:    "reciever",
		TranDate:    "tranDate",
	}

	mockStorage := MockStorage{
		dataMock: []domain.Transaction{},
		errRead:  "",
		errWrite: "can't write to the database",
	}

	// act
	repo := NewRepository(&mockStorage)
	service := NewService(repo)
	errExpected := fmt.Errorf("can't write to the database")
	_, err := service.Store(
		newTransaction.TranCode,
		newTransaction.Currency,
		newTransaction.Amount,
		newTransaction.Transmitter,
		newTransaction.Reciever,
		newTransaction.TranDate,
	)

	// assert
	assert.NotNil(t, err)
	assert.Equal(t, errExpected, err)

}

func TestIntegracionUpdate(t *testing.T) {
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
		dataMock: []domain.Transaction{newTransaction},
		errRead:  "",
		errWrite: "",
	}

	// act
	repo := NewRepository(&mockStorage)
	service := NewService(repo)
	res, err := service.Update(
		newTransaction.Id,
		newTransaction.TranCode,
		newTransaction.Currency,
		newTransaction.Amount,
		newTransaction.Transmitter,
		newTransaction.Reciever,
		newTransaction.TranDate)

	// assert

	assert.Nil(t, err)
	assert.Equal(t, true, mockStorage.ReadCalled)
	assert.Equal(t, mockStorage.dataMock[0], res)
	assert.Equal(t, mockStorage.dataMock[0].Id, res.Id)

}

func TestIntegracionUpdateFail(t *testing.T) {
	newTransaction := domain.Transaction{
		Id:       0,
		TranCode: "PRUEBA",
	}

	mockStorage := MockStorage{
		dataMock: []domain.Transaction{newTransaction},
		errRead:  "",
		errWrite: "",
	}

	// act
	repo := NewRepository(&mockStorage)
	service := NewService(repo)
	_, err := service.Update(
		4,
		newTransaction.TranCode,
		newTransaction.Currency,
		newTransaction.Amount,
		newTransaction.Transmitter,
		newTransaction.Reciever,
		newTransaction.TranDate)

	// assert

	assert.NotNil(t, err)
	assert.Equal(t, true, mockStorage.ReadCalled)

}

func TestIntegracionDelete(t *testing.T) {
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
		dataMock: []domain.Transaction{newTransaction},
		errRead:  "",
		errWrite: "",
	}

	// act
	repo := NewRepository(&mockStorage)
	service := NewService(repo)
	err := service.Delete(newTransaction.Id)

	// assert
	assert.Equal(t, true, mockStorage.ReadCalled)
	assert.Nil(t, err)
}
