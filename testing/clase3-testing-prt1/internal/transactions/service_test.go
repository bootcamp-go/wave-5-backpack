package transactions

import (
	"errors"
	"fmt"
	"testing"

	"clase3-testing-prt1/internal/domain"

	"github.com/stretchr/testify/assert"
)

func TestServiceGetAll(t *testing.T) {
	//	Arrange
	database := []domain.Transaction{
		{
			Id:                1,
			CodigoTransaccion: "abc123",
			Monto:             1100.5,
			Moneda:            "MXN",
			Emisor:            "HSBC",
			Receptor:          "Banorte",
			Fecha:             "2022/05/16",
		}, {
			Id:                2,
			CodigoTransaccion: "cde456",
			Monto:             2500.5,
			Moneda:            "USD",
			Emisor:            "BBVA",
			Receptor:          "Santander",
			Fecha:             "2020/07/08",
		}}

	mockStorage := MockStorage{
		dataMock:       database,
		errWrite:       "",
		errRead:        "",
		readMockCalled: false,
	}

	// act
	repo := NewRepository(&mockStorage)
	service := NewService(repo)
	result, err := service.GetAll()
	// assert
	assert.Nil(t, err)
	assert.Equal(t, mockStorage.dataMock, result)
}

func TestServiceGetAllFail(t *testing.T) {
	// arrange
	expectedError := errors.New("cant read database")
	mockStorage := MockStorage{
		dataMock: nil,
		errWrite: "",
		errRead:  "cant read database",
	}
	// act
	repo := NewRepository(&mockStorage)
	service := NewService(repo)
	result, err := service.GetAll()
	// assert
	assert.ErrorContains(t, err, expectedError.Error())
	assert.Nil(t, result)
}

func TestServiceGetOne(t *testing.T) {
	//	Arrange
	database := []domain.Transaction{
		{
			Id:                1,
			CodigoTransaccion: "abc123",
			Monto:             1100.5,
			Moneda:            "MXN",
			Emisor:            "HSBC",
			Receptor:          "Banorte",
			Fecha:             "2022/05/16",
		},
	}
	mockStorage := MockStorage{
		dataMock:       database,
		errWrite:       "",
		errRead:        "",
		readMockCalled: false,
	}

	// act
	idSelected := 1
	repo := NewRepository(&mockStorage)
	service := NewService(repo)
	result, err := service.GetOne(idSelected)

	// assert
	assert.Nil(t, err)
	assert.Equal(t, mockStorage.dataMock[0], result)
}

func TestServiceGetOneFail(t *testing.T) {
	// arrange
	database := []domain.Transaction{
		{
			Id:                1,
			CodigoTransaccion: "abc123",
			Monto:             1100.5,
			Moneda:            "MXN",
			Emisor:            "HSBC",
			Receptor:          "Banorte",
			Fecha:             "2022/05/16",
		},
	}
	mockStorage := MockStorage{
		dataMock: database,
		errWrite: "",
		errRead:  "",
	}
	// act
	idSelected := 3
	repo := NewRepository(&mockStorage)
	service := NewService(repo)
	result, err := service.GetOne(idSelected)

	expectedError := fmt.Errorf("transaction with id: %d, not found 😵‍💫", idSelected)
	// assert
	assert.Error(t, err, expectedError.Error())
	assert.Empty(t, result)
}

func TestServiceUpdate(t *testing.T) {
	//	Arrange
	testBeforeUpdate := []domain.Transaction{
		{
			Id:                1,
			CodigoTransaccion: "abc123",
			Monto:             1100.5,
			Moneda:            "MXN",
			Emisor:            "HSBC",
			Receptor:          "Banorte",
			Fecha:             "2022/05/16",
		},
		{
			Id:                2,
			CodigoTransaccion: "cde456",
			Monto:             2500.5,
			Moneda:            "USD",
			Emisor:            "BBVA",
			Receptor:          "Santander",
			Fecha:             "2020/07/08",
		},
	}

	testAfterUpdate := []domain.Transaction{
		{
			Id:                1,
			CodigoTransaccion: "abc123",
			Monto:             1100.5,
			Moneda:            "MXN",
			Emisor:            "HSBC",
			Receptor:          "Banorte",
			Fecha:             "2022/05/16",
		},
		{
			Id:                2,
			CodigoTransaccion: "fgh789",
			Monto:             81274.75,
			Moneda:            "MXN",
			Emisor:            "Nu bank",
			Receptor:          "Bancomer",
			Fecha:             "2019/02/11",
		},
	}

	mockStorage := MockStorage{
		dataMock:       testBeforeUpdate,
		errWrite:       "",
		errRead:        "",
		readMockCalled: false,
	}

	// act
	repo := NewRepository(&mockStorage)
	service := NewService(repo)
	result, err := service.Update(
		2,
		testAfterUpdate[1].CodigoTransaccion,
		testAfterUpdate[1].Moneda,
		testAfterUpdate[1].Monto,
		testAfterUpdate[1].Emisor,
		testAfterUpdate[1].Receptor,
		testAfterUpdate[1].Fecha,
	)

	// assert
	assert.Nil(t, err)
	assert.Equal(t, testAfterUpdate[1], mockStorage.dataMock[1])
	assert.Equal(t, testAfterUpdate[1], result, "Deben coincidir")
	assert.True(t, mockStorage.readMockCalled)

}

func TestServiceUpdateFail(t *testing.T) {
	database := []domain.Transaction{
		{
			Id:                1,
			CodigoTransaccion: "abc123",
			Monto:             1100.5,
			Moneda:            "MXN",
			Emisor:            "HSBC",
			Receptor:          "Banorte",
			Fecha:             "2022/05/16",
		}, {
			Id:                2,
			CodigoTransaccion: "cde456",
			Monto:             2500.5,
			Moneda:            "USD",
			Emisor:            "BBVA",
			Receptor:          "Santander",
			Fecha:             "2020/07/08",
		}}
	testUpdate := []domain.Transaction{
		{
			Id:                1,
			CodigoTransaccion: "abc123",
			Monto:             81274.75,
			Moneda:            "USD",
			Emisor:            "Nu bank",
			Receptor:          "Bancomer",
			Fecha:             "2019/02/11",
		},
	}

	// Arrange
	mockStorage := MockStorage{
		dataMock: database,
		errWrite: "",
		errRead:  "",
	}
	// Act
	repo := NewRepository(&mockStorage)
	service := NewService(repo)
	result, err := service.Update(
		1,
		testUpdate[0].CodigoTransaccion,
		testUpdate[0].Moneda,
		testUpdate[0].Monto,
		testUpdate[0].Emisor,
		testUpdate[0].Receptor,
		testUpdate[0].Fecha,
	)

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, testUpdate[0], mockStorage.dataMock[0])
	assert.Equal(t, testUpdate[0], result, "Deben coincidir")
	assert.True(t, mockStorage.readMockCalled)
}

func TestServiceDelete(t *testing.T) {
	//	Arrange
	testBeforeUpdate := []domain.Transaction{
		{
			Id:                1,
			CodigoTransaccion: "abc123",
			Monto:             1100.5,
			Moneda:            "MXN",
			Emisor:            "HSBC",
			Receptor:          "Banorte",
			Fecha:             "2022/05/16",
		},
		{
			Id:                2,
			CodigoTransaccion: "cde456",
			Monto:             2500.5,
			Moneda:            "USD",
			Emisor:            "BBVA",
			Receptor:          "Santander",
			Fecha:             "2020/07/08",
		},
	}

	mockStorage := MockStorage{
		dataMock:       testBeforeUpdate,
		errWrite:       "",
		errRead:        "",
		readMockCalled: false,
	}

	// act
	totalBeforeDeleted := mockStorage.dataMock
	repo := NewRepository(&mockStorage)
	service := NewService(repo)
	err := service.Delete(2)

	// assert
	assert.Nil(t, err)
	assert.NotEqual(t, totalBeforeDeleted, len(mockStorage.dataMock), "Deben contener distinto tamaño en la lista")
	assert.True(t, mockStorage.readMockCalled)

}

func TestServiceDeleteFail(t *testing.T) {
	//	Arrange
	database := []domain.Transaction{
		{
			Id:                1,
			CodigoTransaccion: "abc123",
			Monto:             1100.5,
			Moneda:            "MXN",
			Emisor:            "HSBC",
			Receptor:          "Banorte",
			Fecha:             "2022/05/16",
		},
	}

	mockStorage := MockStorage{
		dataMock: database,
		errWrite: "",
		errRead:  "",
	}

	// Act
	idSelected := 3
	repo := NewRepository(&mockStorage)
	service := NewService(repo)
	err := service.Delete(idSelected)

	expectedError := fmt.Errorf("transaction with id: %d, not found 😵‍💫", idSelected)

	assert.ErrorContains(t, err, expectedError.Error())
}
