package users

import (
	"goweb/internal/domain"

	"testing"

	"github.com/stretchr/testify/assert"
)

// 1) creo lo que sería la pkg/store de mi test
type StubStore struct {}

// 2) ahora tendría que escribir los métodos para esa store

func (fs *StubStore) Read(data interface{}) error{
	a := data.(*[]domain.User)
	*a = []domain.User{
		{Id: 1, Name: "nombre1", LastName: "apellido1", Email: "mail1@mail.com", Age: 22, Height:1.83, Active: true, CreatedAt: "25/07/2022"},
		{Id: 2, Name: "nombre2", LastName: "apellido2", Email: "mail2@mail.com", Age: 23, Height:1.60, Active: true, CreatedAt: "25/07/2022"},
	}
	return nil
}

// van los otros métodos
func (fs *StubStore) Write(data interface{}) error{
	return nil
}
func (fs *StubStore) Ping() error{
	return nil
}

// ahora va el testeo
func TestGetAllUsers(t *testing.T){
	stub := StubStore{}
    repo := NewRepository(&stub)
	resultadoEsperado := []domain.User{
		{Id: 1, Name: "nombre1", LastName: "apellido1", Email: "mail1@mail.com", Age: 22, Height:1.83, Active: true, CreatedAt: "25/07/2022"},
		{Id: 2, Name: "nombre2", LastName: "apellido2", Email: "mail2@mail.com", Age: 23, Height:1.60, Active: true, CreatedAt: "25/07/2022"},
	}

	a, err := repo.GetAllUsers()

	assert.Nil(t,err)
	assert.Equal(t, resultadoEsperado,a)

}


