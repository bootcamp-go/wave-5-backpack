package usuarios

import (
	"Clase4-1/internal/domain"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockStore struct {
	ReadWasCalled bool
	dataMock      []domain.Usuario
	errWrite      string
	errRead       string
}

func (ms *MockStore) Read(data interface{}) error {
	ms.ReadWasCalled = true
	if ms.errRead != "" {
		return fmt.Errorf(ms.errRead)
	}
	a := data.(*[]domain.Usuario)
	*a = ms.dataMock
	return nil

	//	{Id: 1, Nombre: "prueba1", Apellido: "BeforeUpdate", Email: "prueba1Email", Edad: 25, Altura: 180, Activo: true, FechaDeCreacion: "29/10/2004"},
}
func (ms *MockStore) Write(data interface{}) error {
	if ms.errWrite != "" {
		return fmt.Errorf(ms.errWrite)
	}

	a := data.([]domain.Usuario)
	//fmt.Printf("Recibo en write%v", a)
	ms.dataMock = a
	return nil
}
func (fs *MockStore) Ping() error {
	return nil
}

func TestUpdateSurnameAndAge(t *testing.T) {
	//arrange
	database := []domain.Usuario{
		{Id: 1, Nombre: "prueba1", Apellido: "BeforeUpdate", Email: "prueba1Email", Edad: 25, Altura: 180, Activo: true, FechaDeCreacion: "29/10/2004"},
	}
	myMockStore := MockStore{
		dataMock: database,
		errWrite: "",
		errRead:  "",
	}
	repo := NewRepository(&myMockStore)
	resultadoEsperado := domain.Usuario{Id: 1, Nombre: "prueba1", Apellido: "AfterUpdate", Email: "prueba1Email", Edad: 45, Altura: 180, Activo: true, FechaDeCreacion: "29/10/2004"}

	//act
	resultado, err := repo.UpdateSurnameAndAge(1, "AfterUpdate", 45)

	//assert
	assert.Nil(t, err)
	assert.Equal(t, resultadoEsperado, resultado)
	assert.True(t, myMockStore.ReadWasCalled)

}
