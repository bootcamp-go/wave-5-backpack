package repositorio_test

import (
	"C2-C3/internal/domain"
	"C2-C3/internal/repositorio"
	"testing"

	"github.com/stretchr/testify/assert"
)

type StubStore struct{}

func (fs *StubStore) Read(data interface{}) error {
	a := data.(*[]*domain.User)
	u1 := &domain.User{
		FirstName: "required",
		LastName:  "required",
		Email:     "required",
		Age:       5,
		Height:    1.83,
		Activo:    true,
		CreatedAt: "required",
	}
	u2 := &domain.User{
		FirstName: "required",
		LastName:  "required",
		Email:     "required",
		Age:       5,
		Height:    1.83,
		Activo:    true,
		CreatedAt: "required",
	}
	*a = append(*a, u1)
	*a = append(*a, u2)

	return nil
}
func (fs *StubStore) Write(data interface{}) error {
	return nil
}
func (fs *StubStore) Open(data interface{}) error {
	return nil
}

/*
	Ejercicio 1
	Generar un Stub del Store cuya función “Read” retorne dos productos con las especificaciones que deseen.
	Comprobar que GetAll() retorne la información exactamente igual a la esperada.
*/

func TestGetAll(t *testing.T) {
	stub := &StubStore{}
	repo := repositorio.NewRepository(stub)
	expected := []*domain.User{
		{
			FirstName: "required",
			LastName:  "required",
			Email:     "required",
			Age:       5,
			Height:    1.83,
			Activo:    true,
			CreatedAt: "required",
		},
		{
			FirstName: "required",
			LastName:  "required",
			Email:     "required",
			Age:       5,
			Height:    1.83,
			Activo:    true,
			CreatedAt: "required",
		},
	}

	a, err := repo.GetAll()
	assert.Nil(t, err)
	assert.Equal(t, expected, a)

}

/*
Ejercicio 2
*/
type MockStore struct {
	ReadInvoked bool
	Data        []*domain.User
}

func (fs *MockStore) Read(data interface{}) error {
	fs.ReadInvoked = true
	a := data.(*[]*domain.User)
	*a = fs.Data
	return nil
}

func (fs *MockStore) Write(data interface{}) error {
	return nil
}

/* Test with Mock */
func TestUpdateName(t *testing.T) {
	id, nombre := 1, "Update After"
	products := *[]*domain.User{{Id: 1, Nombre: "Update Before", Stock: 1, Precio: 12}} //aki le movi

	mock := MockStore{Data: products}

	r := repositorio.NewRepository(&mock)
	productUpdated, err := r.UpdateName(id, nombre)
	assert.Nil(t, err)

	assert.Equal(t, id, productUpdated.Id)
	assert.Equal(t, nombre, productUpdated.Nombre)
	assert.True(t, true, mock.ReadInvoked)
}
