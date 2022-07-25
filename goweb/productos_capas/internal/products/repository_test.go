package products

import (
	"encoding/json"
	"goweb/productos_capas/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Store struct {
	ReadWasCalled bool
}

var products = []domain.Product{
	{Id: 1, Nombre: "Before Update", Color: "Negro", Precio: 100, Stock: 10, Codigo: "000", Publicado: true, FechaCreacion: "02-11-1999"},
	{Id: 2, Nombre: "Televisor", Color: "Azul", Precio: 200, Stock: 5, Codigo: "111", Publicado: false, FechaCreacion: "12-04-1989"},
}

func (s *Store) Write(data interface{}) error {
	return nil
}

func (s *Store) Read(data interface{}) error {
	s.ReadWasCalled = true
	byteData, _ := json.Marshal(products)
	return json.Unmarshal(byteData, data)
}

func (s *Store) Ping() error {
	return nil
}

func TestStub(t *testing.T) {
	myStore := Store{}
	repository := NewRepository(&myStore)

	resultado, err := repository.GetAll("", "", 0, 0, "", false, "")
	assert.Equal(t, products, resultado, "deben ser iguales")
	assert.Equal(t, nil, err, "deben ser iguales")
}

func TestMock(t *testing.T) {
	myStore := Store{}
	repository := NewRepository(&myStore)

	nombreEsperado := "After Update"
	precioEsperado := 50
	resultado, err := repository.UpdateNamePrice(1, nombreEsperado, precioEsperado)

	assert.Equal(t, true, myStore.ReadWasCalled)
	assert.Equal(t, nombreEsperado, resultado.Nombre)
	assert.Equal(t, precioEsperado, resultado.Precio)
	assert.Equal(t, nil, err)
}
