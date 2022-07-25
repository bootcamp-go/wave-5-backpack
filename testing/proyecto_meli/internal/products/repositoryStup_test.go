package products

import (
	"proyecto_meli/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

type StupStore struct{}

func (d StupStore) Read(data interface{}) error {
	a := data.(*[]domain.Product)
	*a = []domain.Product{
		{Id: 1, Nombre: "licuadora", Color: "rojo", Precio: 1000, Stock: 10, Codigo: "elec01", Publicado: true, FechaCreacion: "25-07-2022"},
		{Id: 2, Nombre: "cafetera", Color: "azul", Precio: 1500, Stock: 2, Codigo: "elec02", Publicado: false, FechaCreacion: "25-07-2022"},
	}
	return nil
}

func (d StupStore) Write(data interface{}) error {
	return nil
}

func (d StupStore) Ping() error {
	return nil
}

func TestGetAll(t *testing.T) {
	//arrange

	myStupStore := StupStore{}
	repositoryProduct := NewRepository(myStupStore)
	respuestaEsperada := []domain.Product{
		{Id: 1, Nombre: "licuadora", Color: "rojo", Precio: 1000, Stock: 10, Codigo: "elec01", Publicado: true, FechaCreacion: "25-07-2022"},
		{Id: 2, Nombre: "cafetera", Color: "azul", Precio: 1500, Stock: 2, Codigo: "elec02", Publicado: false, FechaCreacion: "25-07-2022"},
	}

	//act
	resultado, err := repositoryProduct.GetAll()

	//assert
	assert.Equal(t, respuestaEsperada, resultado)
	assert.Nil(t, err)

}
