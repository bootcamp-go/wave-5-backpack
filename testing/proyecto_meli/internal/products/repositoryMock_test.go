package products

import (
	"proyecto_meli/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockStore struct {
	ReadWasCalled bool
}

func (d *MockStore) Read(data interface{}) error {
	d.ReadWasCalled = true
	a := data.(*[]domain.Product)
	*a = []domain.Product{
		{Id: 1, Nombre: "Before Update", Color: "rojo", Precio: 1000, Stock: 10, Codigo: "elec01", Publicado: true, FechaCreacion: "25-07-2022"},
	}
	return nil
}

func (d MockStore) Write(data interface{}) error {
	return nil
}

func (d MockStore) Ping() error {
	return nil
}

func Test_Update_Name_Price(t *testing.T) {
	//arrange

	myMockStore := MockStore{}
	repositoryProduct := NewRepository(&myMockStore)
	respuestaEsperada := domain.Product{
		Id: 1, Nombre: "After Update", Color: "rojo", Precio: 0, Stock: 10, Codigo: "elec01", Publicado: true, FechaCreacion: "25-07-2022",
	}

	//act
	resultado, err := repositoryProduct.Update_Name_Price(1, "After Update", 0)

	//assert
	assert.Equal(t, respuestaEsperada, resultado)
	assert.Nil(t, err)
	assert.True(t, myMockStore.ReadWasCalled)

}
