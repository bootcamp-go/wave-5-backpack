package users

import (
	"clase2_2/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

type StubStore struct{}

func (f *StubStore) Read(data interface{}) error {
	users := []domain.User{{Id: 1, Name: "nombre1", LastName: "apellido1", Mail: "mail1", Years: 25, Tall: 1, Enable: true, CreateDate: "0/0/0"},
		{Id: 2, Name: "nombre2", LastName: "apellido2", Mail: "mail2", Years: 20, Tall: 2, Enable: false, CreateDate: "0/0/0"}}
	(*data.(*[]domain.User)) = users
	return nil
}
func (f *StubStore) Write(data interface{}) error {

	return nil
}

//Ejercicio 1
func TestGetAll(t *testing.T) {
	db := &StubStore{}
	repository := NewRepository(db)
	esperado := []domain.User{{Id: 1, Name: "nombre1", LastName: "apellido1", Mail: "mail1", Years: 25, Tall: 1, Enable: true, CreateDate: "0/0/0"},
		{Id: 2, Name: "nombre2", LastName: "apellido2", Mail: "mail2", Years: 20, Tall: 2, Enable: false, CreateDate: "0/0/0"}}

	resultado, err := repository.GetAll()

	assert.Nil(t, err)
	assert.Equal(t, esperado, resultado, "No son iguales")
}
