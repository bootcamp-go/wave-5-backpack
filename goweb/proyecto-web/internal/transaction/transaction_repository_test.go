package transaction

import (
	"proyecto-web/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Ejercicio 1 - Crear un stub del storage para probar el método GetAll()
func TestGetAll(t *testing.T) {
	// arrange
	stub := StubStorage{}
	service := NewRepository(&stub)

	expected := []domain.Transaction{
		{
			Id:                1,
			CodigoTransaccion: "A1",
			Moneda:            "PESOS",
			Monto:             5.0,
			Emisor:            "ARCOR",
			Receptor:          "AFIP",
			FechaTransaccion:  "12-01-2022",
		},
		{
			Id:                2,
			CodigoTransaccion: "A2",
			Moneda:            "DOLARES",
			Monto:             20.0,
			Emisor:            "TOYOTA",
			Receptor:          "AFIP",
			FechaTransaccion:  "22-06-2022",
		},
	}

	// act
	result, _ := service.GetAll()

	// assert
	assert.Equal(t, len(result), 2)
	assert.Equal(t, result[0], expected[0])
	assert.Equal(t, result[1], expected[1])
}

// Ejercicio 2: Crear un mock del Storage para probar el métood UpdateName (UpdateParcial para este proyecto). Verificar la invocación del método Read
func TestUpdateParcial(t *testing.T) {
	//arrange
	transaction := []domain.Transaction{
		{
			Id:                1,
			CodigoTransaccion: "BEFORE UPDATE",
			Moneda:            "PESOS",
			Monto:             5.0,
			Emisor:            "ARCOR",
			Receptor:          "AFIP",
			FechaTransaccion:  "12-01-2022",
		},
	}

	StorageMock := &MockStorage{dataMock: transaction}
	repo := NewRepository(StorageMock)

	// act
	previusData, _ := repo.GetById(1)
	previusCodigo := previusData.CodigoTransaccion
	updatedData, err := repo.UpdateParcial(1, "AFTER UPDATE", 5.0)

	// assert
	assert.Equal(t, true, StorageMock.readWasCalled)
	assert.Nil(t, err)
	assert.Equal(t, "BEFORE UPDATE", previusCodigo)
	assert.Equal(t, "AFTER UPDATE", updatedData.CodigoTransaccion)
}
