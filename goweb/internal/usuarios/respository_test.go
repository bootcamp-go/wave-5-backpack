package usuarios

import (
	"testing"

	"github.com/bootcamp-go/wave-5-backpack/goweb/internal/domain"
	"github.com/stretchr/testify/assert"
)

type StubStore struct {
}

func (ss *StubStore) Validate() error {
	return nil
}

func (ss *StubStore) Write(data interface{}) error {
	return nil
}

func (ss *StubStore) Read(data interface{}) error {
	a := data.(*[]domain.Usuarios) //PUNTERO DE LISTA DE USUARIOS
	*a = []domain.Usuarios{        //ACA LLENO ESOS VALORES DEL PUNTERO, por eso lo desreferencio
		{Id: 1, Nombre: "Yvo", Apellido: "Pintos", Altura: 3, FechaCreacion: "1992"},
		{Id: 2, Nombre: "Pedro", Apellido: "Juan", Altura: 3, FechaCreacion: "1232"},
	}
	return nil

}

func TestGetAllRepo(t *testing.T) {
	myStubStore := StubStore{}
	repo := NewRepository(&myStubStore) //Probando el repository, yo le paso datos dummy a lo que quiero probar
	expected := []domain.Usuarios{
		{Id: 1, Nombre: "Yvo", Apellido: "Pintos", Altura: 3, FechaCreacion: "1992"},
		{Id: 2, Nombre: "Pedro", Apellido: "Juan", Altura: 3, FechaCreacion: "1232"},
	}

	user, err := repo.GetAll()

	assert.Equal(t, user, expected)
	assert.Nil(t, err)
}
