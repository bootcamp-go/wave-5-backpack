package products

import (
	"fmt"
	"goweb/clase1_clase2/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockStorage struct {
	ReadWasCalled bool
	dataMock      []domain.Product
	errWrite      string
	errRead       string
}

var products = []domain.Product{
	{Id: 1, Nombre: "Before Update", Color: "Blanco", Precio: 600, Stock: 4, Codigo: "B453", Publicado: true, Fecha: "05-05-2022"},
	{Id: 2, Nombre: "Televisor", Color: "Verde", Precio: 100, Stock: 4, Codigo: "V479", Publicado: false, Fecha: "17-06-2022"},
}

func (m *MockStorage) Write(data interface{}) error {
	if m.errWrite != "" {
		return fmt.Errorf(m.errWrite)
	}
	a := data.([]domain.Product)
	m.dataMock = a
	return nil
}

func (m *MockStorage) Read(data interface{}) error {
	m.ReadWasCalled = true
	if m.errRead != "" {
		return fmt.Errorf(m.errRead)
	}
	a := data.(*[]domain.Product)
	*a = m.dataMock
	return nil
}

func (m *MockStorage) Ping() error {
	return nil
}

func TestMock(t *testing.T) {
	//arrange
	database := []domain.Product{
		{Id: 1, Nombre: "Before Update", Color: "Blanco", Precio: 600, Stock: 4, Codigo: "B453", Publicado: true, Fecha: "05-05-2022"},
		{Id: 2, Nombre: "Televisor", Color: "Verde", Precio: 100, Stock: 4, Codigo: "V479", Publicado: false, Fecha: "17-06-2022"},
	}
	mockStorage := MockStorage{
		dataMock: database,
		errWrite: "",
		errRead:  "",
	}
	nombreEsperado := "After Update"
	precioEsperado := 3
	//act
	repo := NewRepository(&mockStorage)
	resultado, err := repo.UpdateFields(1, nombreEsperado, precioEsperado)

	assert.Equal(t, true, mockStorage.ReadWasCalled)
	assert.Equal(t, nombreEsperado, resultado.Nombre)
	assert.Equal(t, precioEsperado, resultado.Precio)
	assert.Equal(t, nil, err)
}

func TestGetAll(t *testing.T) {
	//arrange
	database := []domain.Product{
		{
			Id:        1,
			Nombre:    "Nevera",
			Color:     "Blanco",
			Precio:    600,
			Stock:     4,
			Codigo:    "B453",
			Publicado: true,
			Fecha:     "05-05-2022",
		},
		{
			Id:        1,
			Nombre:    "Televisor",
			Color:     "Verde",
			Precio:    400,
			Stock:     9,
			Codigo:    "T543",
			Publicado: true,
			Fecha:     "02-06-2022",
		},
	}
	mockStorage := MockStorage{
		dataMock: database,
		errWrite: "",
		errRead:  "",
	}

	//act
	repo := NewRepository(&mockStorage)
	result, err := repo.GetAll("", "", 0, 0, "", false, "")

	//assert
	assert.Nil(t, err)
	assert.Equal(t, mockStorage.dataMock, result)
}
