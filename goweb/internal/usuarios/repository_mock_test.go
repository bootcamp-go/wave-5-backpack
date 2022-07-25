package usuarios

import (
	"testing"

	"github.com/bootcamp-go/wave-5-backpack/goweb/internal/domain"
	"github.com/stretchr/testify/assert"
)

type MockStore struct {
	readWasCalled bool
}

func (ms *MockStore) Validate() error {
	return nil
}

func (ms *MockStore) Write(data interface{}) error {
	return nil
}

func (ms *MockStore) Read(data interface{}) error {
	BeforeUpd := data.(*[]domain.Usuarios) //ACA ESTOY RECIBIENDO DESDE REPOSITORY UN PUNTERO DE LISTA DE USUARIOS
	*BeforeUpd = []domain.Usuarios{        //ACA LLENO ESOS VALORES DEL PUNTERO, por eso lo desreferencio
		{Id: 1, Nombre: "Yvo", Apellido: "Pintos", Altura: 3, FechaCreacion: "1992"},
		{Id: 2, Nombre: "Pedro", Apellido: "Juan", Altura: 3, FechaCreacion: "1232"},
	}
	ms.readWasCalled = true
	return nil

}

func TestUpdateNAL(t *testing.T) {
	myMockStore := MockStore{}
	repo := NewRepository(&myMockStore) //Probando el repository, yo le paso datos dummy a lo que quiero probar
	expected := domain.Usuarios{Id: 1, Nombre: "Nuevo", Apellido: "Nuevo", Altura: 3, FechaCreacion: "1992"}

	user, err := repo.UpdateNameAndLastName(1, "Nuevo", "Nuevo")
	assert.True(t, myMockStore.readWasCalled)
	assert.Equal(t, user, expected)
	assert.Nil(t, err)
}
