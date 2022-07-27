package users

import (
	"github.com/stretchr/testify/assert"
	"goweb/internal/domain"
	"testing"
)

type StubStore struct{}

func (fs *StubStore) Ping() error {
	return nil
}

func (fs *StubStore) Read(data interface{}) error {
	user := data.(*[]domain.User)
	user1 := domain.User{
		Id:            3,
		Nombre:        "Daniela",
		Apellido:      "Bedoya",
		Email:         "bedoya@gmail.com",
		Edad:          20,
		Altura:        1.61,
		Activo:        true,
		FechaCreacion: "2021-10-02T04:44:12 +05:00",
	}
	user2 := domain.User{
		Id:            4,
		Nombre:        "Daniela",
		Apellido:      "Bedoya",
		Email:         "bedoya@gmail.com",
		Edad:          20,
		Altura:        1.61,
		Activo:        true,
		FechaCreacion: "2021-10-02T04:44:12 +05:00",
	}

	*user = append(*user, user1)
	*user = append(*user, user2)
	return nil
}

func (fs *StubStore) Write(data interface{}) error {
	return nil
}

func TestGetAll(t *testing.T) {
	stub := &StubStore{}
	repo := NewRepository(stub)
	/*expected := []domain.User{
		{
			Id:            3,
			Nombre:        "Daniela",
			Apellido:      "Bedoya",
			Email:         "bedoya@gmail.com",
			Edad:          20,
			Altura:        1.61,
			Activo:        true,
			FechaCreacion: "2021-10-02T04:44:12 +05:00",
		},
		{
			Id:            4,
			Nombre:        "Daniela",
			Apellido:      "Bedoya",
			Email:         "bedoya@gmail.com",
			Edad:          20,
			Altura:        1.61,
			Activo:        true,
			FechaCreacion: "2021-10-02T04:44:12 +05:00",
		},
		{
			Id:            3,
			Nombre:        "Daniela",
			Apellido:      "After Update",
			Email:         "bedoya@gmail.com",
			Edad:          21,
			Altura:        1.61,
			Activo:        true,
			FechaCreacion: "2021-10-02T04:44:12 +05:00",
		},
	}*/

	result, err := repo.GetAll()
	assert.Nil(t, err)
	//assert.Equal(t, expected, result)
	assert.Equal(t, 3, len(result))

}
