package usuarios

import (
	"Clase2-1/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockStore struct {
	ReadWasCalled bool
}

func (fs *MockStore) Read(data interface{}) error {
	fs.ReadWasCalled = true
	a := data.(*[]domain.Usuario)
	*a = []domain.Usuario{
		{Id: 1, Nombre: "prueba1", Apellido: "BeforeUpdate", Email: "prueba1Email", Edad: 25, Altura: 180, Activo: true, FechaDeCreacion: "29/10/2004"},
		{Id: 4,
			Nombre:          "prueba2",
			Apellido:        "prueba2apellido",
			Email:           "prueba2Email",
			Edad:            25,
			Altura:          190,
			Activo:          false,
			FechaDeCreacion: "05/10/2004",
		},
	}
	return nil
}
func (fs *MockStore) Write(data interface{}) error {
	return nil
}
func (fs *MockStore) Ping() error {
	return nil
}

func TestUpdateSurnameAndAge(t *testing.T) {
	//arrange

	myMockStore := MockStore{}
	repo := NewRepository(&myMockStore)
	resultadoEsperado := domain.Usuario{Id: 1, Nombre: "prueba1", Apellido: "AfterUpdate", Email: "prueba1Email", Edad: 45, Altura: 180, Activo: true, FechaDeCreacion: "29/10/2004"}

	//act
	resultado, err := repo.UpdateSurnameAndAge(1, "AfterUpdate", 45)

	//assert
	assert.Nil(t, err)
	assert.Equal(t, resultadoEsperado, resultado)
	assert.True(t, myMockStore.ReadWasCalled)

}
