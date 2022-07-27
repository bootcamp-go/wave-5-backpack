package users

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/bootcamp-go/wave-5-backpack/goweb/internal/domain"
	"github.com/stretchr/testify/assert"
)

type MockIntegrationStorage struct {
	readInvoked bool
	userData    []domain.ModelUser
	errWrite    string
	errRead     string
}

func (mis *MockIntegrationStorage) Read(data interface{}) error {
	if mis.errRead != "" {
		return fmt.Errorf(mis.errRead)
	}
	mis.readInvoked = true
	user := data.(*[]domain.ModelUser)
	*user = mis.userData
	return nil
}

func (mis *MockIntegrationStorage) Write(data interface{}) error {
	if mis.errWrite != "" {
		return fmt.Errorf(mis.errWrite)
	}
	user := data.(*[]domain.ModelUser)
	//mis.userData = append(mis.userData, *user...)
	mis.userData = *user
	return nil
}

func (mis *MockIntegrationStorage) Ping() error {
	return nil
}

func TestIntegrationUpdate(t *testing.T) {
	// Arrange
	fecha, _ := time.Parse("2006-01-02", "2022-07-25")
	users := []domain.ModelUser{
		{Id: 1, Nombre: "Juan", Apellido: "Perez", Email: "juan.perez@gmail.com", Edad: 22, Altura: 1.60, Activo: true, FechaCreacion: fecha, Borrado: false},
		{Id: 2, Nombre: "Norma", Apellido: "Carrasco", Email: "norma.carrasco@gmail.com", Edad: 28, Altura: 1.56, Activo: false, FechaCreacion: fecha, Borrado: false},
	}
	mock := MockIntegrationStorage{userData: users, errWrite: "", errRead: ""}
	newUser := domain.ModelUser{Id: 2, Nombre: "Norma", Apellido: "Carrasco", Email: "norma.carrasco@gmail.com", Edad: 28, Altura: 1.56, Activo: false, FechaCreacion: fecha, Borrado: false}

	// Act
	repo := NewRepository(&mock)
	serv := NewService(repo)
	updated, err := serv.Update(newUser.Id, newUser.Nombre, newUser.Apellido, newUser.Email, newUser.Edad, newUser.Altura)

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, mock.userData[1], updated)
	assert.True(t, mock.readInvoked)
}

func TestIntegrationUpdateFailRead(t *testing.T) {
	// Arrange
	errorEsperado := errors.New("no se puede leer la db, error: error al leer la bd")
	mock := MockIntegrationStorage{
		userData: nil,
		errWrite: "",
		errRead:  "error al leer la bd",
	}

	// Act
	repo := NewRepository(&mock)
	service := NewService(repo)
	updated, err := service.Update(1, "", "", "", 0, 0)

	// Assert
	assert.Equal(t, errorEsperado, err)
	assert.Equal(t, updated.Id, 0)
}

func TestIntegrationUpdateFailWrite(t *testing.T) {
	// Arrange
	expectedError := errors.New("no se puede escribir en la db, error: error al escribir la bd")
	fecha, _ := time.Parse("2006-01-02", "2022-07-25")
	database := []domain.ModelUser{
		{Id: 1, Nombre: "Juan", Apellido: "Perez", Email: "juan.perez@gmail.com", Edad: 22, Altura: 1.60, Activo: true, FechaCreacion: fecha, Borrado: false},
		{Id: 2, Nombre: "Norma", Apellido: "Carrasco", Email: "norma.carrasco@gmail.com", Edad: 28, Altura: 1.56, Activo: false, FechaCreacion: fecha, Borrado: false},
	}
	mock := MockIntegrationStorage{
		userData: database,
		errWrite: "error al escribir la bd",
		errRead:  "",
	}

	// Act
	repo := NewRepository(&mock)
	service := NewService(repo)
	user, err := service.Update(1, "Update", "UpdateF", "UpdateF", 25, 1.72)

	// Assert
	assert.Equal(t, expectedError, err)
	assert.Equal(t, user.Id, 0)
}

func TestIntegrationDelete(t *testing.T) {
	// Arrange
	fecha, _ := time.Parse("2006-01-02", "2022-07-25")
	database := []domain.ModelUser{
		{Id: 1, Nombre: "Juan", Apellido: "Perez", Email: "juan.perez@gmail.com", Edad: 22, Altura: 1.60, Activo: true, FechaCreacion: fecha, Borrado: false},
		{Id: 2, Nombre: "Norma", Apellido: "Carrasco", Email: "norma.carrasco@gmail.com", Edad: 28, Altura: 1.56, Activo: false, FechaCreacion: fecha, Borrado: false},
	}

	mock := MockIntegrationStorage{
		userData: database,
		errWrite: "",
		errRead:  "",
	}

	// Act
	repo := NewRepository(&mock)
	service := NewService(repo)
	err := service.Delete(1)

	// Assert
	assert.Nil(t, err)
	assert.True(t, mock.userData[0].Borrado)
}

func TestIntegrationDeleteError(t *testing.T) {
	// Arrange
	fecha, _ := time.Parse("2006-01-02", "2022-07-25")
	database := []domain.ModelUser{
		{Id: 1, Nombre: "Juan", Apellido: "Perez", Email: "juan.perez@gmail.com", Edad: 22, Altura: 1.60, Activo: true, FechaCreacion: fecha, Borrado: false},
		{Id: 2, Nombre: "Norma", Apellido: "Carrasco", Email: "norma.carrasco@gmail.com", Edad: 28, Altura: 1.56, Activo: false, FechaCreacion: fecha, Borrado: false},
	}

	mockInteg := MockIntegrationStorage{
		userData: database,
		errWrite: "",
		errRead:  "",
	}

	errorEsperado := fmt.Errorf("usuario 123 no econtrado")

	// Act
	repo := NewRepository(&mockInteg)
	service := NewService(repo)
	err := service.Delete(123)

	// Assert
	assert.ErrorContains(t, err, errorEsperado.Error())
}
