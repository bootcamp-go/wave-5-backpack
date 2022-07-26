package transaction

import (
	"proyecto-web/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Ejercicio 1 - Crear un stub del storage para probar el método GetAll()

type StubStorage struct{}

func TestGetAll(t *testing.T) {
	// arrange
	stub := StubStorage{}
	service := NewRepository(&stub)

	expected := []domain.Transaction{
		{
			Id:                0,
			CodigoTransaccion: "A1",
			Moneda:            "PESOS",
			Monto:             5.0,
			Emisor:            "ARCOR",
			Receptor:          "AFIP",
			FechaTransaccion:  "12-01-2022",
		},
		{
			Id:                1,
			CodigoTransaccion: "A2",
			Moneda:            "DOLARES",
			Monto:             20.0,
			Emisor:            "TOYOTA",
			Receptor:          "AFIP",
			FechaTransaccion:  "22-06-2022",
		},
	}

	// act
	result := service.GetAll()

	// assert
	assert.Equal(t, len(result), 2)
	assert.Equal(t, result[0], expected[0])
	assert.Equal(t, result[1], expected[1])
}

func (s *StubStorage) Read(data interface{}) error {
	dataCasteada := data.(*[]domain.Transaction)

	*dataCasteada = []domain.Transaction{
		{
			Id:                0,
			CodigoTransaccion: "A1",
			Moneda:            "PESOS",
			Monto:             5.0,
			Emisor:            "ARCOR",
			Receptor:          "AFIP",
			FechaTransaccion:  "12-01-2022",
		},
		{
			Id:                1,
			CodigoTransaccion: "A2",
			Moneda:            "DOLARES",
			Monto:             20.0,
			Emisor:            "TOYOTA",
			Receptor:          "AFIP",
			FechaTransaccion:  "22-06-2022",
		},
	}
	return nil
}
func (s *StubStorage) Write(data interface{}) error {
	return nil
}

// Ejercicio 2: Crear un mock del Storage para probar el métood UpdateName (UpdateParcial para este proyecto). Verificar la invocación del método Read

type MockStorage struct {
	Data          []domain.Transaction
	ReadWasCalled bool
}

func TestUpdateParcial(t *testing.T) {
	//arrange
	transaction := []domain.Transaction{
		{
			Id:                0,
			CodigoTransaccion: "BEFORE UPDATE",
			Moneda:            "PESOS",
			Monto:             5.0,
			Emisor:            "ARCOR",
			Receptor:          "AFIP",
			FechaTransaccion:  "12-01-2022",
		},
	}

	StorageMock := &MockStorage{Data: transaction}
	repo := NewRepository(StorageMock)

	// act
	previusData, _ := repo.GetById(0)
	previusCodigo := previusData.CodigoTransaccion
	updatedData, _ := repo.UpdateParcial(0, "AFTER UPDATE", 5.0)

	// assert
	assert.Equal(t, true, StorageMock.ReadWasCalled)
	assert.Equal(t, "BEFORE UPDATE", previusCodigo)
	assert.Equal(t, "AFTER UPDATE", updatedData.CodigoTransaccion)
}

func (m *MockStorage) Read(data interface{}) error {
	m.ReadWasCalled = true
	dataCasteada := data.(*[]domain.Transaction)
	*dataCasteada = m.Data
	return nil
}

func (s *MockStorage) Write(data interface{}) error {
	return nil
}
