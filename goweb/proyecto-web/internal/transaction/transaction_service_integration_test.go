package transaction

import (
	"proyecto-web/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

//Ejercicio 1 - Testear Update con sus casos de acierto y fallo
func TestUpdate(t *testing.T) {
	dbFake := []domain.Transaction{
		{
			Id:                1,
			CodigoTransaccion: "A-0",
			Moneda:            "PESOS",
			Monto:             5.3,
			Emisor:            "ARCOR",
			Receptor:          "AFIP",
			FechaTransaccion:  "12-03-2021",
		},
	}
	mockStorage := MockStorage{dataMock: dbFake}

	repo := NewRepository(&mockStorage)
	service := NewService(repo)

	dataUpdated, err := service.Update(1, "ACTUALIZADO", "DOLARES", 10.0, "NATURA", "BANCO GALICIA", "23-05-2022")

	assert.Nil(t, err)
	assert.Equal(t, true, mockStorage.readWasCalled)
	assert.Equal(t, 1, dataUpdated.Id)
	assert.Equal(t, "ACTUALIZADO", dataUpdated.CodigoTransaccion)
	assert.Equal(t, "DOLARES", dataUpdated.Moneda)
	assert.Equal(t, 10.0, dataUpdated.Monto)
	assert.Equal(t, "NATURA", dataUpdated.Emisor)
	assert.Equal(t, "BANCO GALICIA", dataUpdated.Receptor)
	assert.Equal(t, "23-05-2022", dataUpdated.FechaTransaccion)
}

func TestUpdateFailNotFound(t *testing.T) {
	dbFake := []domain.Transaction{
		{
			Id:                1,
			CodigoTransaccion: "A-0",
			Moneda:            "PESOS",
			Monto:             5.3,
			Emisor:            "ARCOR",
			Receptor:          "AFIP",
			FechaTransaccion:  "12-03-2021",
		},
	}
	mockStorage := MockStorage{dataMock: dbFake}

	repo := NewRepository(&mockStorage)
	service := NewService(repo)

	_, err := service.Update(89, "ACTUALIZADO", "DOLARES", 10.0, "NATURA", "BANCO GALICIA", "23-05-2022")

	assert.Equal(t, true, mockStorage.readWasCalled)
	assert.NotNil(t, err)
	assert.ErrorContains(t, err, "not found")
}

func TestUpdateFailWrite(t *testing.T) {
	dbFake := []domain.Transaction{
		{
			Id:                1,
			CodigoTransaccion: "A-0",
			Moneda:            "PESOS",
			Monto:             5.3,
			Emisor:            "ARCOR",
			Receptor:          "AFIP",
			FechaTransaccion:  "12-03-2021",
		},
	}
	mockStorage := MockStorage{dataMock: dbFake, errWrite: "cant write database"}

	repo := NewRepository(&mockStorage)
	service := NewService(repo)

	_, err := service.Update(1, "ACTUALIZADO", "DOLARES", 10.0, "NATURA", "BANCO GALICIA", "23-05-2022")

	assert.Equal(t, true, mockStorage.readWasCalled)
	assert.NotNil(t, err)
	assert.ErrorContains(t, err, "cant write database")
}

func TestUpdateFailNotRead(t *testing.T) {
	dbFake := []domain.Transaction{
		{
			Id:                1,
			CodigoTransaccion: "A-0",
			Moneda:            "PESOS",
			Monto:             5.3,
			Emisor:            "ARCOR",
			Receptor:          "AFIP",
			FechaTransaccion:  "12-03-2021",
		},
	}
	mockStorage := MockStorage{dataMock: dbFake, errRead: "cant read database"}

	repo := NewRepository(&mockStorage)
	service := NewService(repo)

	_, err := service.Update(1, "ACTUALIZADO", "DOLARES", 10.0, "NATURA", "BANCO GALICIA", "23-05-2022")

	assert.Equal(t, true, mockStorage.readWasCalled)
	assert.NotNil(t, err)
	assert.ErrorContains(t, err, "cant read database")
}

//Ejercicio 2 - Testear Delete con sus casos de acierto y fallo
func TestDelete(t *testing.T) {
	dbFake := []domain.Transaction{
		{
			Id:                1,
			CodigoTransaccion: "A-0",
			Moneda:            "PESOS",
			Monto:             5.3,
			Emisor:            "ARCOR",
			Receptor:          "AFIP",
			FechaTransaccion:  "12-03-2021",
		},
	}

	mockStorage := MockStorage{dataMock: dbFake}
	repo := NewRepository(&mockStorage)
	service := NewService(repo)

	err := service.Delete(1)
	assert.Equal(t, true, mockStorage.readWasCalled)
	assert.Nil(t, err)
	assert.Equal(t, 0, len(mockStorage.dataMock))
}

func TestDeleteFailNotFound(t *testing.T) {
	dbFake := []domain.Transaction{
		{
			Id:                1,
			CodigoTransaccion: "A-0",
			Moneda:            "PESOS",
			Monto:             5.3,
			Emisor:            "ARCOR",
			Receptor:          "AFIP",
			FechaTransaccion:  "12-03-2021",
		},
	}

	mockStorage := MockStorage{dataMock: dbFake}
	repo := NewRepository(&mockStorage)
	service := NewService(repo)

	err := service.Delete(18)

	assert.Equal(t, true, mockStorage.readWasCalled)
	assert.NotNil(t, err)
	assert.ErrorContains(t, err, "not found")
}

func TestDeleteFailWrite(t *testing.T) {
	dbFake := []domain.Transaction{
		{
			Id:                1,
			CodigoTransaccion: "A-0",
			Moneda:            "PESOS",
			Monto:             5.3,
			Emisor:            "ARCOR",
			Receptor:          "AFIP",
			FechaTransaccion:  "12-03-2021",
		},
	}

	mockStorage := MockStorage{dataMock: dbFake, errWrite: "cant write database"}
	repo := NewRepository(&mockStorage)
	service := NewService(repo)

	err := service.Delete(1)

	assert.Equal(t, true, mockStorage.readWasCalled)
	assert.NotNil(t, err)
	assert.ErrorContains(t, err, "cant write database")
}

func TestDeleteFailRead(t *testing.T) {
	dbFake := []domain.Transaction{
		{
			Id:                1,
			CodigoTransaccion: "A-0",
			Moneda:            "PESOS",
			Monto:             5.3,
			Emisor:            "ARCOR",
			Receptor:          "AFIP",
			FechaTransaccion:  "12-03-2021",
		},
	}

	mockStorage := MockStorage{dataMock: dbFake, errRead: "cant read database"}
	repo := NewRepository(&mockStorage)
	service := NewService(repo)

	err := service.Delete(1)

	assert.Equal(t, true, mockStorage.readWasCalled)
	assert.NotNil(t, err)
	assert.ErrorContains(t, err, "cant read database")
}
