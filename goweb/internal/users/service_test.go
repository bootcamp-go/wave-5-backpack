package users

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"goweb/internal/domain"
	"testing"
)

type mockService struct {
	data     []domain.User
	Err      error
	isCalled bool
}

func (m *mockService) Read(data interface{}) error {
	if m.Err != nil {
		return m.Err
	}
	user := data.(*[]domain.User)
	*user = m.data
	m.isCalled = true
	return nil
}

func (m *mockService) Write(data interface{}) error {
	if m.Err != nil {
		return m.Err
	}
	user := data.([]domain.User)
	m.data = append(m.data, user...)
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
		data: users,
		Err:  errors.New("Error: "),
	}

	repo := NewRepository(&mock)
	service := NewService(repo)
	_, err := service.Update(6, "Update", "Update", "bedoya@gmail.com", 21, 1.61, true, "2021-10-02T04:44:12 +05:00")

	expectError := fmt.Errorf("error al actualizar el usuario %d", 6)
	assert.EqualError(t, err, expectError.Error())
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
	}

	mock := mockService{
		data: users,
		Err:  nil,
	}

	repo := NewRepository(&mock)
	service := NewService(repo)
	err := service.Delete(1)

	assert.Nil(t, err.Error())
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
		Err:  nil,
	}

	repo := NewRepository(&mock)
	service := NewService(repo)
	err := service.Delete(8)

	assert.ErrorContains(t, err, "error al eliminar el usuario 8")
}
