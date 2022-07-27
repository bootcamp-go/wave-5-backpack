package users

import (
	"clase2_2/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MuckStore struct {
	isReadCalled bool
}

func (f *MuckStore) Read(data interface{}) error {
	users := []domain.User{{Id: 1, Name: "nombre1", LastName: "apellido1", Mail: "mail1", Years: 25, Tall: 1, Enable: true, CreateDate: "0/0/0"},
		{Id: 2, Name: "nombre2", LastName: "apellido2", Mail: "mail2", Years: 20, Tall: 2, Enable: false, CreateDate: "0/0/0"}}
	(*data.(*[]domain.User)) = users
	f.isReadCalled = true
	return nil
}
func (f *MuckStore) Write(data interface{}) error {

	return nil
}

//Ejercicio 2
func TestMuckUpdate(t *testing.T) {
	db := &MuckStore{}
	repository := NewRepository(db)
	esperado := domain.User{Id: 1, Name: "nombre10", LastName: "apellido10", Mail: "mail10", Years: 20, Tall: 10, Enable: false, CreateDate: "1/1/1"}

	resultado, err := repository.UpdateUser("nombre10", "apellido10", "mail10", "1/1/1", 20, 1, 10, false)

	assert.Nil(t, err)
	assert.True(t, db.isReadCalled, "no se a ejecutado el método")
	assert.Equal(t, esperado, resultado, "No son iguales")
}
