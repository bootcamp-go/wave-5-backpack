package products

import (
	"goweb/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Stub Struct

type StubStore struct{}

func (s *StubStore) Read(data interface{}) error {
	products := data.(*[]domain.Product)
	*products = append(*products, domain.Product{Id: 1, Nombre: "producto 1", Color: "rojo", Precio: 10, Stock: 10, Codigo: "123", Publicado: true, FechaCreacion: "2020-01-01"})
	*products = append(*products, domain.Product{Id: 2, Nombre: "producto 2", Color: "rojo", Precio: 10, Stock: 10, Codigo: "123", Publicado: true, FechaCreacion: "2020-01-01"})
	return nil
}

func (s *StubStore) Write(data interface{}) error {
	return nil
}

func (s *StubStore) Ping() error {
	return nil
}

// Test

func TestRepositoryGetAll(t *testing.T) {
	// -------------------------------------------------------
	// Se inicializan los datos a usar en el test (input y output)
	// -------------------------------------------------------

	myStubStore := StubStore{}
	myRepository := NewRepository(&myStubStore)
	resultadoEsperado := []domain.Product{
		{Id: 1, Nombre: "producto 1", Color: "rojo", Precio: 10, Stock: 10, Codigo: "123", Publicado: true, FechaCreacion: "2020-01-01"},
		{Id: 2, Nombre: "producto 2", Color: "rojo", Precio: 10, Stock: 10, Codigo: "123", Publicado: true, FechaCreacion: "2020-01-01"},
	}

	// -------------------------------------------------------
	// Se ejecuta el test
	// -------------------------------------------------------

	resultadoObtenido, err := myRepository.GetAll()

	// -------------------------------------------------------
	// Se compara el resultado obtenido con el resultado esperado
	// -------------------------------------------------------

	assert.Nil(t, err)
	assert.Equal(t, resultadoEsperado, resultadoObtenido)
}
