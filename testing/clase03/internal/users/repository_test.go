
package users

import (
	"goweb/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

// 1) creo lo que sería la pkg/store de mi test
type MockStore struct {
	ReadWasCalled bool
	dataMock []domain.User
	errWrite string
	errRead string
}

// así quedaría para el test
func (m *MockStore) Read(data interface{}) error{
	
	BeforeUpdate := data.(*[]domain.User)//ACA ESTOY RECIBIENDO DESDE REPOSITORY UN PUNTERO DE LISTA DE USUARIOS
	*BeforeUpdate = []domain.User{//ACA LLENO ESOS VALORES DEL PUNTERO, por eso lo desreferencio
		{Id: 1, Name: "nombre1", LastName: "apellido1", Email: "mail1@mail.com", Age: 22, Height:1.83, Active: true, CreatedAt: "25/07/2022"},
		{Id: 2, Name: "nombre2", LastName: "apellido2", Email: "mail2@mail.com", Age: 23, Height:1.60, Active: true, CreatedAt: "25/07/2022"},
	}

	m.ReadWasCalled = true
	return nil
}

// van los otros métodos
func (m *MockStore) Write(data interface{}) error{
	return nil
}

func (m *MockStore) Ping() error{
	return nil
}

// ahora va el testeo

func TestGetAllUsers(t *testing.T){
	mock := MockStore{}
    repo := NewRepository(&mock)
	resultadoEsperado := []domain.User{
		{Id: 1, Name: "nombre1", LastName: "apellido1", Email: "mail1@mail.com", Age: 22, Height:1.83, Active: true, CreatedAt: "25/07/2022"},
		{Id: 2, Name: "nombre2", LastName: "apellido2", Email: "mail2@mail.com", Age: 23, Height:1.60, Active: true, CreatedAt: "25/07/2022"},
	}

	a, err := repo.GetAllUsers()

	assert.Nil(t,err)
	assert.Equal(t, resultadoEsperado,a)

}

func TestUpdateTotal(t *testing.T){
	mock := MockStore{
		dataMock: []domain.User{
			{Id: 1, Name: "beforeNombre", LastName: "beforeApellido1", Email: "mail1@mail.com", Age: 22, Height:1.83, Active: true, CreatedAt: "25/07/2022"},
			{Id: 2, Name: "nombre2", LastName: "apellido2", Email: "mail2@mail.com", Age: 23, Height:1.60, Active: true, CreatedAt: "25/07/2022"},
		},
	}
    repo := NewRepository(&mock) //Probando el repository, yo le paso datos dummy a lo que quiero probar

	resultadoEsperado := domain.User{
		Id: 1, Name: "AfterNombre",
		LastName: "AfterApellido",
		Email: "mail1@mail.com",
		Age: 22,
		Height:1.83,
		Active: true,
		CreatedAt: "25/07/2022"}

	afterUpdate, err := repo.UpdateTotal(1, "AfterNombre", "AfterApellido", "mail1@mail.com", 22, 1.83, true, "25/07/2022")

	assert.Equal(t, resultadoEsperado,afterUpdate)
	assert.True(t, mock.ReadWasCalled)
	assert.Nil(t,err)
	
}


