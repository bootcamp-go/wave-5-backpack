package usuarios

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/bootcamp-go/wave-5-backpack/goweb/internal/domain"
	"github.com/stretchr/testify/assert"
)

type MockIntegration struct {
	dataMock      []domain.Usuarios
	errWrite      string
	errRead       string
	readWasCalled bool
}

func (ms *MockIntegration) Validate() error {
	return nil
}

func (ms *MockIntegration) Write(data interface{}) error {
	if ms.errWrite != "" {
		return fmt.Errorf(ms.errWrite)
	}
	a := data.([]domain.Usuarios)
	ms.dataMock = append(ms.dataMock, a...)
	return nil
}

func (ms *MockIntegration) Read(data interface{}) error {
	if ms.errRead != "" {
		return fmt.Errorf(ms.errRead)
	}
	user := data.(*[]domain.Usuarios)
	*user = ms.dataMock
	ms.readWasCalled = true
	return nil
}

func TestIntegrationUpdate(t *testing.T) {
	// arrange
	database := []domain.Usuarios{
		{
			Id: 1, Nombre: "Yvo", Apellido: "Pintos", Altura: 3, FechaCreacion: "1992",
		},
		{
			Id: 2, Nombre: "Titan", Apellido: "Pintos", Altura: 1, FechaCreacion: "2019",
		},
	}

	newUser := domain.Usuarios{
		Id: 1, Nombre: "New", Apellido: "New", Email: "yvop", Edad: 10, Altura: 3, Activo: true, FechaCreacion: "1992",
	}

	mockInte := MockIntegration{
		dataMock: database,
		errWrite: "",
		errRead:  "",
	}

	// act
	repo := NewRepository(&mockInte)
	service := NewService(repo)
	user, err := service.Update(
		context.TODO(),
		newUser.Id,
		newUser.Nombre,
		newUser.Apellido,
		newUser.Email,
		newUser.Edad,
		newUser.Altura,
		newUser.Activo,
		newUser.FechaCreacion,
	)

	// assert
	assert.Nil(t, err)
	assert.Equal(t, mockInte.dataMock[0], user)
	assert.Equal(t, mockInte.dataMock[0].Id, user.Id)
	assert.True(t, mockInte.readWasCalled)
}

func TestIntegrationUpdateFailRead(t *testing.T) {
	// arrange
	expectedError := errors.New("error al leer la bd")
	mockInte := MockIntegration{
		dataMock: nil,
		errWrite: "",
		errRead:  "error al leer la bd",
	}
	// act
	repo := NewRepository(&mockInte)
	service := NewService(repo)
	user, err := service.Update(context.TODO(), 1, "", "", "", 0, 0, true, "")
	// assert
	assert.Equal(t, expectedError, err)
	assert.Equal(t, user.Id, 0)
}

func TestIntegrationUpdateFailWrite(t *testing.T) {
	// arrange
	expectedError := errors.New("error al escribir la bd")
	database := []domain.Usuarios{
		{
			Id: 1, Nombre: "Yvo", Apellido: "Pintos", Altura: 3, FechaCreacion: "1992",
		},
	}
	mockInte := MockIntegration{
		dataMock: database,
		errWrite: "error al escribir la bd",
		errRead:  "",
	}
	// act
	repo := NewRepository(&mockInte)
	service := NewService(repo)
	user, err := service.Update(context.TODO(), 1, "F", "F", "F", 3, 2, true, "F")
	// assert
	assert.Equal(t, expectedError, err)
	assert.Equal(t, user.Id, 0)
}

func TestIntegrationDelete(t *testing.T) {
	// arrange
	database := []domain.Usuarios{
		{
			Id: 1, Nombre: "Yvo", Apellido: "Pintos", Altura: 3, FechaCreacion: "1992",
		},
		{
			Id: 2, Nombre: "Titan", Apellido: "Pintos", Altura: 1, FechaCreacion: "2019",
		},
	}

	mockInteg := MockIntegration{
		dataMock: database,
		errWrite: "",
		errRead:  "",
	}

	// act
	repo := NewRepository(&mockInteg)
	service := NewService(repo)
	err := service.Delete(1)

	// assert
	assert.Nil(t, err)
	assert.NotEqual(t, mockInteg.dataMock[0].Id, 1)
}

func TestIntegrationDeleteError(t *testing.T) {
	// arrange
	database := []domain.Usuarios{
		{
			Id: 1, Nombre: "Yvo", Apellido: "Pintos", Altura: 3, FechaCreacion: "1992",
		},
		{
			Id: 2, Nombre: "Titan", Apellido: "Pintos", Altura: 1, FechaCreacion: "2019",
		},
	}

	mockInteg := MockIntegration{
		dataMock: database,
		errWrite: "",
		errRead:  "",
	}

	userNotFound := fmt.Errorf("usuario 33 not found")

	// act
	repo := NewRepository(&mockInteg)
	service := NewService(repo)
	errNotFound := service.Delete(33)

	// assert
	assert.ErrorContains(t, errNotFound, userNotFound.Error())
}
