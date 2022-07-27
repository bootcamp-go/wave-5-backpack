package users

import (
	"testing"

	"github.com/JofeGonzalezMeLi/goweb/cmd/internal/domain"
	"github.com/stretchr/testify/assert"
)

type StubStore struct{}

type MockStore struct {
	ReadInvoked bool
	Data        []domain.User
}

func (st *StubStore) Read(data interface{}) error {
	u := data.(*[]domain.User)
	p1 := domain.User{
		Id:             1,
		Edad:           55,
		Nombre:         "Domingo",
		Apellido:       "Croce",
		Email:          "dc@meli.com",
		Fecha_creacion: "16/03/2022",
		Altura:         170,
		Activo:         false,
	}
	p2 := domain.User{
		Id:             2,
		Edad:           59,
		Nombre:         "Nea",
		Apellido:       "Licet",
		Email:          "nl@meli.com",
		Fecha_creacion: "08/06/2022",
		Altura:         168,
		Activo:         true,
	}
	*u = append(*u, p1)
	*u = append(*u, p2)
	return nil
}
func (st *StubStore) Write(data interface{}) error {
	return nil
}

func (mt *MockStore) Read(data interface{}) error {
	mt.ReadInvoked = true
	u := data.(*[]domain.User)
	*u = mt.Data
	return nil
}
func (mt *MockStore) Write(data interface{}) error {
	return nil
}
func TestGetAll(t *testing.T) {
	stub := &StubStore{}
	repo := NewRepository(stub)
	expected := []domain.User{
		{
			Id:             1,
			Edad:           55,
			Nombre:         "Domingo",
			Apellido:       "Croce",
			Email:          "dc@meli.com",
			Fecha_creacion: "16/03/2022",
			Altura:         170,
			Activo:         false,
		},
		{
			Id:             2,
			Edad:           59,
			Nombre:         "Nea",
			Apellido:       "Licet",
			Email:          "nl@meli.com",
			Fecha_creacion: "08/06/2022",
			Altura:         168,
			Activo:         true,
		},
	}

	a, err := repo.GetAll()
	assert.Nil(t, err)
	assert.Equal(t, expected, a)
}

func TestUpdate(t *testing.T) {
	id, edad, apellido := 1, 23, "Gonzalez"
	users := []domain.User{{Id: 1, Edad: 22, Nombre: "Domingo", Apellido: "Crocce", Email: "dc@meli.com", Fecha_creacion: "16/03/2022", Altura: 170, Activo: false}}
	mockStore := MockStore{Data: users}
	r := NewRepository(&mockStore)
	pu, err := r.UpdateLastNameAndAge(id, edad, apellido)
	assert.Nil(t, err)
	assert.Equal(t, id, pu.Id)
	assert.Equal(t, edad, pu.Edad)
	assert.Equal(t, apellido, pu.Apellido)
	assert.True(t, true, mockStore.ReadInvoked)
}
