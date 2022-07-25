package users

import (
	"ejercicioTT/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

type StubStore struct {
}

func (fs *StubStore) Read(data interface{}) error {
	all := data.(*[]domain.Usuarios)
	*all = []domain.Usuarios{
		{Id: 2, Nombre: "Marcela", Apellido: "Monroy", Email: "marcela@hotmail.com", Edad: 27, Altura: 1.67},
		{Id: 3, Nombre: "Marcelo", Apellido: "Moncada", Email: "marcelo@hotmail.com", Edad: 20, Altura: 1.82},
	}
	return nil
}

func (fs *StubStore) Write(data interface{}) error {
	return nil
}

func (fs *StubStore) Ping() error {
	return nil
}

func TestGetAll(t *testing.T) {
	stub := StubStore{}
	repo := NewRepository(&stub)
	expected := []domain.Usuarios{
		{Id: 2, Nombre: "Marcela", Apellido: "Monroy", Email: "marcela@hotmail.com", Edad: 27, Altura: 1.67},
		{Id: 3, Nombre: "Marcelo", Apellido: "Moncada", Email: "marcelo@hotmail.com", Edad: 20, Altura: 1.82},
	}
	all, err := repo.GetAll()
	assert.Nil(t, err)
	assert.Equal(t, expected, all, "No coincide la información de usuario esperada con la obtenida")
}
