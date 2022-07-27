package users

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"goweb/internal/domain"
	"testing"
)

type mockService struct {
	data     []domain.User
	ErrRead  error
	ErrWrite error
	isCalled bool
}

func (m *mockService) Read(data interface{}) error {
	if m.ErrRead != nil {
		return m.ErrRead
	}
	user := data.(*[]domain.User)
	*user = m.data
	m.isCalled = true
	return nil
}

func (m *mockService) Write(data interface{}) error {
	if m.ErrWrite != nil {
		return m.ErrWrite
	}
	user := data.([]domain.User)
	m.data = append(m.data, user[len(user)-1])
	return nil
}

func (m mockService) Ping() error {
	return nil
}

func TestUpdate(t *testing.T) {
	users := []domain.User{
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
	}

	mock := mockService{
		data: users,
	}

	expectUser := domain.User{
		Id:            3,
		Nombre:        "Update",
		Apellido:      "Update",
		Email:         "bedoya@gmail.com",
		Edad:          21,
		Altura:        1.61,
		Activo:        true,
		FechaCreacion: "2021-10-02T04:44:12 +05:00",
	}

	repo := NewRepository(&mock)
	service := NewService(repo)
	userUpdated, err := service.Update(3, "Update", "Update", "bedoya@gmail.com", 21, 1.61, true, "2021-10-02T04:44:12 +05:00")

	assert.Nil(t, err)
	assert.Equal(t, expectUser, userUpdated)
	assert.True(t, true, mock.isCalled)
	assert.Equal(t, mock.data[0].Id, 3)
}

func TestUpdateFailedRead(t *testing.T) {
	users := []domain.User{
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
	}

	mock := mockService{
		data:    users,
		ErrRead: fmt.Errorf("error al leer el archivo"),
	}

	repo := NewRepository(&mock)
	service := NewService(repo)
	_, err := service.Update(6, "Update", "Update", "bedoya@gmail.com", 21, 1.61, true, "2021-10-02T04:44:12 +05:00")

	assert.EqualError(t, err, mock.ErrRead.Error())
}

func TestDelete(t *testing.T) {
	users := []domain.User{
		{
			Id:            1,
			Nombre:        "Daniela",
			Apellido:      "Bedoya",
			Email:         "bedoya@gmail.com",
			Edad:          20,
			Altura:        1.61,
			Activo:        true,
			FechaCreacion: "2021-10-02T04:44:12 +05:00",
		},
		{
			Id:            2,
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
	}

	mock := mockService{
		data: users,
	}

	repo := NewRepository(&mock)
	service := NewService(repo)
	err := service.Delete(1)

	assert.Nil(t, err)
}

func TestDeleteNotFound(t *testing.T) {
	users := []domain.User{
		{
			Id:            1,
			Nombre:        "Daniela",
			Apellido:      "Bedoya",
			Email:         "bedoya@gmail.com",
			Edad:          20,
			Altura:        1.61,
			Activo:        true,
			FechaCreacion: "2021-10-02T04:44:12 +05:00",
		},
		{
			Id:            2,
			Nombre:        "Daniela",
			Apellido:      "Bedoya",
			Email:         "bedoya@gmail.com",
			Edad:          20,
			Altura:        1.61,
			Activo:        true,
			FechaCreacion: "2021-10-02T04:44:12 +05:00",
		},
	}

	mock := mockService{
		data: users,
	}

	repo := NewRepository(&mock)
	service := NewService(repo)
	err := service.Delete(8)

	assert.ErrorContains(t, err, "error al eliminar el usuario 8")
}
