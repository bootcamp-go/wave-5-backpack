package products

import (
	"fmt"
	"goweb/productos_capas/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockStorage struct {
	dataMock      []domain.Product
	errWrite      string
	errRead       string
	ReadWasCalled bool
}

func (m *MockStorage) Read(data interface{}) error {
	if m.errRead != "" {
		return fmt.Errorf(m.errRead)
	}
	m.ReadWasCalled = true
	a := data.(*[]domain.Product)
	*a = m.dataMock
	return nil
}

func (m *MockStorage) Write(data interface{}) error {
	if m.errWrite != "" {
		return fmt.Errorf(m.errWrite)
	}
	a := data.([]domain.Product)
	//m.dataMock = append(m.dataMock, a...)
	m.dataMock = a
	return nil
}

func (m *MockStorage) Ping() error {
	return nil
}

func TestGetAll(t *testing.T) {
	//arrange
	database := []domain.Product{
		{Id: 1, Nombre: "Before Update", Color: "Negro", Precio: 100, Stock: 10, Codigo: "000", Publicado: true, FechaCreacion: "02-11-1999"},
		{Id: 2, Nombre: "Televisor", Color: "Azul", Precio: 200, Stock: 5, Codigo: "111", Publicado: false, FechaCreacion: "12-04-1989"},
	}
	mockStorage := MockStorage{
		dataMock: database,
		errWrite: "",
		errRead:  "",
	}

	//act
	repository := NewRepository(&mockStorage)
	resultado, err := repository.GetAll("", "", 0, 0, "", false, "")

	//assert
	assert.Equal(t, mockStorage.dataMock, resultado, "deben ser iguales")
	assert.Equal(t, nil, err, "deben ser iguales")
}

func TestUpdate(t *testing.T) {
	//arrange
	database := []domain.Product{
		{Id: 1, Nombre: "Before Update", Color: "Negro", Precio: 100, Stock: 10, Codigo: "000", Publicado: true, FechaCreacion: "02-11-1999"},
		{Id: 2, Nombre: "Televisor", Color: "Azul", Precio: 200, Stock: 5, Codigo: "111", Publicado: false, FechaCreacion: "12-04-1989"},
	}
	mockStorage := MockStorage{
		dataMock: database,
		errWrite: "",
		errRead:  "",
	}

	//act
	nombreEsperado := "After Update"
	precioEsperado := 50
	repository := NewRepository(&mockStorage)
	resultado, err := repository.UpdateNamePrice(1, nombreEsperado, precioEsperado)

	//assert
	assert.Equal(t, true, mockStorage.ReadWasCalled)
	assert.Equal(t, nombreEsperado, resultado.Nombre)
	assert.Equal(t, precioEsperado, resultado.Precio)
	assert.Equal(t, nil, err)
}
