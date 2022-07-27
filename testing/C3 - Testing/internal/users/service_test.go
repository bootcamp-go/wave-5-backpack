package users

import (
	"C3/internal/domain"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServiceIntegrationGetAll(t *testing.T) {
	// arrange
	database := []domain.Usuarios{
		{Id: 2, Nombre: "Marcela", Apellido: "Monroy", Email: "marcela@hotmail.com", Edad: 27, Altura: 1.67},
		{Id: 3, Nombre: "Marcelo", Apellido: "Moncada", Email: "marcelo@hotmail.com", Edad: 20, Altura: 1.82},
	}
	mockStorage := MockStorage{
		dataMock: database,
		errWrite: "",
		errRead:  "",
	}
	// act
	repo := NewRepository(&mockStorage)
	service := NewService(repo)
	result, err := service.GetAll("", "", 0, 0, "", false, "")
	// assert
	assert.Nil(t, err)
	assert.Equal(t, mockStorage.dataMock, result)
}

func TestServiceIntegrationGetAllFail(t *testing.T) {
	// arrange
	expectedError := fmt.Errorf("can't read database")
	mockStorage := MockStorage{
		dataMock: nil,
		errWrite: "",
		errRead:  "can't read database",
	}
	// act
	repo := NewRepository(&mockStorage)
	service := NewService(repo)
	result, err := service.GetAll()
	// assert
	assert.Equal(t, expectedError, err)
	assert.Nil(t, result)
}

func TestServiceIntegrationStore(t *testing.T) {
	// arrange
	newUser := []domain.Usuarios{
		{Id: 5, Nombre: "Marcos", Apellido: "Monroy", Email: "marcos@hotmail.com", Edad: 29, Altura: 1.89}
	}
	mockStorage := MockStorage{
		dataMock: nil,
		errWrite: "",
		errRead:  "",
	}
	// act
	repo := NewRepository(&mockStorage)
	service := NewService(repo)
	result, err := service.Store(
		newUser.Nombre,
		newUser.Apellido,
		newUser.Email,
		newUser.Edad,
		newUser.Altura)
	// assert
	assert.Nil(t, err)
	assert.Equal(t, mockStorage.dataMock[0], result)
	assert.Equal(t, mockStorage.dataMock[0].Id, 5)
}

func TestServiceIntegrationStoreFail(t *testing.T) {
	// arrange
	newUser := []domain.Usuarios{
		{Id: 5, Nombre: "Marcos", Apellido: "Monroy", Email: "marcos@hotmail.com", Edad: 29, Altura: 1.89}
	}
	writeErr := fmt.Errorf("can't write database")
	expectedError := fmt.Errorf("can't write database, error: %w", writeErr)
	mockStorage := MockStorage{
		dataMock: nil,
		errWrite: "can't write database",
		errRead:  "",
	}
	// act
	repo := NewRepository(&mockStorage)
	service := NewService(repo)
	result, err := service.Store(
		newUser.Nombre,
		newUser.Apellido,
		newUser.Email,
		newUser.Edad,
		newUser.Altura)
	// assert
	assert.Equal(t, expectedError, err)
	assert.Equal(t, domain.Usuarios{}, result)
}

func TestServiceIntegrationUpdate(t *testing.T) {
	// arrange
	updateUser := domain.Usuarios{
		{Id: 5, Nombre: "Marcos", Apellido: "Monroy", Email: "marcos@hotmail.com", Edad: 29, Altura: 1.89}
	}
	mockStorage := MockStorage{
		dataMock: []domain.Usuarios{updateUser},
		errWrite: "",
		errRead:  "",
	}
	// act
	repo := NewRepository(&mockStorage)
	service := NewService(repo)
	result, err := service.Update(
		updateUser.Id,
		updateUser.Nombre,
		updateUser.Apellido,
		updateUser.Email,
		updateUser.Edad,
		updateUser.Altura)
	// assert
	assert.Nil(t, err)
	assert.True(t, mockStorage.ReadWasCalled)
	assert.Equal(t, mockStorage.dataMock[0], result)
	assert.Equal(t, mockStorage.dataMock[0].Id, 1)
}

func TestServiceIntegrationUpdateFail(t *testing.T) {
	// arrange
	updateUser := domain.Usuarios{
		{Id: 5, Nombre: "Marcos", Apellido: "Monroy", Email: "marcos@hotmail.com", Edad: 29, Altura: 1.89}
	}
	expectedError := fmt.Errorf("cant read database")
	mockStorage := MockStorage{
		dataMock: nil,
		errWrite: "",
		errRead:  "cant read database",
	}
	// act
	repo := NewRepository(&mockStorage)
	service := NewService(repo)
	result, err := service.Update(
		updateUser.Id,
		updateUser.Nombre,
		updateUser.Apellido,
		updateUser.Email,
		updateUser.Edad,
		updateUser.Altura)
	// assert
	assert.True(t, mockStorage.ReadWasCalled)
	assert.Equal(t, expectedError, err)
	assert.Equal(t, domain.Usuarios{}, result)
}

func TestServiceIntegrationDelete(t *testing.T) {
	// arrange
	deleteUser := domain.Usuarios{
		{Id: 5, Nombre: "Marcos", Apellido: "Monroy", Email: "marcos@hotmail.com", Edad: 29, Altura: 1.89}
	}
	mockStorage := MockStorage{
		dataMock: []domain.Usuarios{deleteUser},
		errWrite: "",
		errRead:  "",
	}
	// act
	repo := NewRepository(&mockStorage)
	service := NewService(repo)
	_, err := service.Delete(deleteUser.Id)
	// assert
	assert.Nil(t, err)
	assert.Equal(t, mockStorage.dataMock, []domain.Usuarios{})
}

func TestServiceIntegrationDeleteFail(t *testing.T) {
	// arrange
	deleteUser := domain.Usuarios{
		{Id: 5, Nombre: "Marcos", Apellido: "Monroy", Email: "marcos@hotmail.com", Edad: 29, Altura: 1.89}
	}
	expectedError := fmt.Errorf("user %d not found", 2)
	mockStorage := MockStorage{
		dataMock: []domain.Usuarios{deleteUser},
		errWrite: "",
		errRead:  "",
	}
	// act
	repo := NewRepository(&mockStorage)
	service := NewService(repo)
	result, err := service.Delete(2)
	// assert
	assert.Equal(t, expectedError, err)
	assert.Equal(t, domain.Usuarios{}, result)
}
