package products

import (
	"encoding/json"
	"goweb/clase1_clase2/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Store struct {
	ReadWasCalled bool
}

var products = []domain.Product{
	{Id: 1, Nombre: "Before Update", Color: "Blanco", Precio: 600, Stock: 4, Codigo: "B453", Publicado: true, Fecha: "05-05-2022"},
	{Id: 2, Nombre: "Televisor", Color: "Verde", Precio: 100, Stock: 4, Codigo: "V479", Publicado: false, Fecha: "17-06-2022"},
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
	precioEsperado := 3
	resultado, err := repository.UpdateFields(1, nombreEsperado, precioEsperado)

	assert.Equal(t, true, myStore.ReadWasCalled)
	assert.Equal(t, nombreEsperado, resultado.Nombre)
	assert.Equal(t, precioEsperado, resultado.Precio)
	assert.Equal(t, nil, err)
}
