package users

import (
	"clase2_2/internal/domain"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type MuckStore struct {
	dataMock     []domain.User
	errWrite     string
	errRead      string
	isReadCalled bool
}

func (f *MuckStore) Read(data interface{}) error {
	if f.errRead != "" {
		return fmt.Errorf(f.errRead)
	}
	a := data.(*[]domain.User)
	*a = f.dataMock

	f.isReadCalled = true
	return nil
}
func (f *MuckStore) Write(data interface{}) error {

	return nil
}

//Ejercicio 2
func TestMuckGetAll(t *testing.T) {
	db := &MuckStore{}
	repository := NewRepository(db)
	esperado := domain.User{Id: 1, Name: "nombre10", LastName: "apellido10", Mail: "mail10", Years: 20, Tall: 10, Enable: false, CreateDate: "1/1/1"}

	resultado, err := repository.UpdateUser("nombre10", "apellido10", "mail10", "1/1/1", 20, 1, 10, false)

	assert.Nil(t, err)
	assert.True(t, db.isReadCalled, "no se a ejecutado el método")
	assert.Equal(t, esperado, resultado, "No son iguales")
}
