package products

import (
	"goweb/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Stub Struct

type MockStore struct {
	ReadWasCalled bool
}

func (m *MockStore) Read(data interface{}) error {
	m.ReadWasCalled = true

	products := data.(*[]domain.Product)
	*products = append(*products, domain.Product{Id: 1, Nombre: "Before Update", Color: "rojo", Precio: 10, Stock: 10, Codigo: "123", Publicado: true, FechaCreacion: "2020-01-01"})
	return nil
}

func (s *MockStore) Write(data interface{}) error {
	return nil
}

func (s *MockStore) Ping() error {
	return nil
}

// Test

func TestRepositoryUpdateNombre(t *testing.T) {
	// -------------------------------------------------------
	// Se inicializan los datos a usar en el test (input y output)
	// -------------------------------------------------------

	myMockStore := MockStore{}
	myRepository := NewRepository(&myMockStore)
	resultadoEsperado := domain.Product{
		Id:            1,
		Nombre:        "After Update",
		Color:         "rojo",
		Precio:        10,
		Stock:         10,
		Codigo:        "123",
		Publicado:     true,
		FechaCreacion: "2020-01-01",
	}

	// -------------------------------------------------------
	// Se ejecuta el test
	// -------------------------------------------------------

	resultadoObtenido, err := myRepository.UpdateNombre(1, "After Update")

	// -------------------------------------------------------
	// Se compara el resultado obtenido con el resultado esperado
	// -------------------------------------------------------

	assert.Nil(t, err)
	assert.Equal(t, resultadoEsperado, resultadoObtenido)
	assert.True(t, myMockStore.ReadWasCalled)
}

func TestRepositoryUpdatePrecio(t *testing.T) {
	// -------------------------------------------------------
	// Se inicializan los datos a usar en el test (input y output)
	// -------------------------------------------------------

	myMockStore := MockStore{}
	myRepository := NewRepository(&myMockStore)
	resultadoEsperado := domain.Product{
		Id:            1,
		Nombre:        "Before Update",
		Color:         "rojo",
		Precio:        200,
		Stock:         10,
		Codigo:        "123",
		Publicado:     true,
		FechaCreacion: "2020-01-01",
	}

	// -------------------------------------------------------
	// Se ejecuta el test
	// -------------------------------------------------------

	resultadoObtenido, err := myRepository.UpdatePrecio(1, 200)

	// -------------------------------------------------------
	// Se compara el resultado obtenido con el resultado esperado
	// -------------------------------------------------------

	assert.Nil(t, err)
	assert.Equal(t, resultadoEsperado, resultadoObtenido)
	assert.True(t, myMockStore.ReadWasCalled)
}
