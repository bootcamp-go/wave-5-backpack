package users

import (
	"github.com/stretchr/testify/assert"
	"goweb/internal/domain"
	"testing"
)

type MockStore struct {
	Data     []domain.User
	isCalled bool
}

func (fs *MockStore) Ping() error {
	return nil
}

func (fs *MockStore) Read(data interface{}) error {
	fs.isCalled = true
	users := data.(*[]domain.User)
	*users = fs.Data
	return nil
}

func (fs *MockStore) Write(data interface{}) error {
	return nil
}

func TestUpdateName(t *testing.T) {
	user := []domain.User{
		{
			Id:            3,
			Nombre:        "Daniela",
			Apellido:      "Update Before",
			Email:         "bedoya@gmail.com",
			Edad:          20,
			Altura:        1.61,
			Activo:        true,
			FechaCreacion: "2021-10-02T04:44:12 +05:00",
		},
	}

	mock := MockStore{Data: user}

	repo := NewRepository(&mock)
	userUpdated, err := repo.Patch(3, "After Update", 21)
	assert.Nil(t, err)

	assert.True(t, true, mock.isCalled)
	assert.Equal(t, 3, userUpdated.Id)
	assert.Equal(t, "After Update", userUpdated.Apellido)
}
