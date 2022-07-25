package users

import (
	"ejercicioTT/internal/domain"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockStore struct {
	infodb interface{}
	read   bool
}

func (ms *MockStore) Read(data interface{}) error {
	r := reflect.ValueOf(data)
	r = reflect.Indirect(r)
	r.Set(reflect.ValueOf(ms.infodb))
	ms.read = true
	return nil
}

func (ms *MockStore) Write(data interface{}) error {
	ms.infodb = data
	return nil
}

func (ms *MockStore) Ping() error {
	return nil
}

func TestUpdate(t *testing.T) {
	testUpdateUsersNow := []domain.Usuarios{
		{Id: 2, Nombre: "Marcela", Apellido: "Monroy", Email: "marcela@hotmail.com", Edad: 27, Altura: 1.67},
		{Id: 3, Nombre: "Marcelo", Apellido: "Moncada", Email: "marcelo@hotmail.com", Edad: 20, Altura: 1.82},
	}

	testUpdateUsersThen := []domain.Usuarios{
		{Id: 2, Nombre: "Marcela", Apellido: "Martinez", Email: "marcela@hotmail.com", Edad: 28, Altura: 1.67},
		{Id: 3, Nombre: "Marcelo", Apellido: "Moncada", Email: "marcelo@hotmail.com", Edad: 20, Altura: 1.82},
	}

	info_update := &MockStore{
		infodb: testUpdateUsersNow,
		read:   false,
	}

	repository := NewRepository(info_update)

	result, _ := repository.UpdateLastAge(2, "Martinez", 28)
	assert.Equal(t, testUpdateUsersThen, info_update.infodb)
	assert.Equal(t, testUpdateUsersThen[0], result)
	assert.True(t, info_update.read)
}
