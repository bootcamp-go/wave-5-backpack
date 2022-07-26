package transaction

import (
	"fmt"
	"proyecto-web/internal/domain"
)

type MockStorage struct {
	dataMock      []domain.Transaction
	readWasCalled bool
	errWrite      string
	errRead       string
}

type StubStorage struct{}

// Funciones para el Stub (ejercicio 1)
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

// Funciones para el Mock (ejercicio 2)
func (m *MockStorage) Read(data interface{}) error {
	m.readWasCalled = true
	if m.errRead != "" {
		return fmt.Errorf(m.errRead)
	}

	dataCasteada := data.(*[]domain.Transaction)
	*dataCasteada = m.dataMock
	return nil
}

func (m *MockStorage) Write(data interface{}) error {
	if m.errWrite != "" {
		return fmt.Errorf(m.errWrite)
	}

	dataCasteada := data.([]domain.Transaction)
	m.dataMock = dataCasteada
	return nil
}
