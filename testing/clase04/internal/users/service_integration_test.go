package users

import (
	//"fmt"
	"fmt"
	"goweb/internal/domain"
	"testing"

	//"errors"
	"github.com/stretchr/testify/assert"
)


func TestIntegrationUpdateTotal(t *testing.T){
	// arrange
	dataBase := []domain.User{
		{Id: 1, Name: "nombre1", LastName: "apellido1", Email: "mail1@mail.com", Age: 22, Height:1.83, Active: true, CreatedAt: "25/07/2022"},
		{Id: 2, Name: "nombre2", LastName: "apellido2", Email: "mail2@mail.com", Age: 23, Height:1.60, Active: true, CreatedAt: "25/07/2022"},
	}

	mockStorage := MockStore{
		dataMock: dataBase,
		errRead: "",
		errWrite: "",
		ReadWasCalled: false,
	}

	// act
    repo := NewRepository(&mockStorage)
	service := NewService(repo)
	result, err := service.UpdateTotal(1, "nombre1_actualizado", "apellido1_actualizado",  "mail1_actualizado@mail.com", 23, 1.85, true, "25/07/2022")

	//assert
	assert.Nil(t, err)
	assert.Equal(t, mockStorage.dataMock[0], result)
	assert.Equal(t, mockStorage.dataMock[0].Id, 1)
	assert.True(t, mockStorage.ReadWasCalled)

}

func TestIntegrationUpdateTotalFailRead(t *testing.T){
	// arrange

	mockStorage := MockStore{
		dataMock: nil,
		errRead: "cant read database",
		errWrite: "",
		ReadWasCalled: false,
	}
	readErr := fmt.Errorf("cant read database")
	expectedError := fmt.Errorf("error updating user: %w", readErr)

	// act
    repo := NewRepository(&mockStorage)
	service := NewService(repo)
	result, err := service.UpdateTotal(1, "nombre1_actualizado", "apellido1_actualizado",  "mail1_actualizado@mail.com", 23, 1.85, true, "25/07/2022")

	//assert
	assert.Equal(t, domain.User{}, result)
	assert.ErrorContains(t, err, expectedError.Error())
	assert.False(t, mockStorage.ReadWasCalled)

}

func TestIntegrationUpdateTotalFailWrite(t *testing.T){
	// arrange
	dataBase := []domain.User{
		{Id: 1, Name: "nombre1", LastName: "apellido1", Email: "mail1@mail.com", Age: 22, Height:1.83, Active: true, CreatedAt: "25/07/2022"},
		{Id: 2, Name: "nombre2", LastName: "apellido2", Email: "mail2@mail.com", Age: 23, Height:1.60, Active: true, CreatedAt: "25/07/2022"},
	}

	mockStorage := MockStore{
		dataMock: dataBase,
		errRead: "",
		errWrite: "cant write database",
		ReadWasCalled: false,
	}
	writeErr := fmt.Errorf("cant write database")
	expectedError := fmt.Errorf("error updating user: %w", writeErr)

	// act
    repo := NewRepository(&mockStorage)
	service := NewService(repo)
	result, err := service.UpdateTotal(1, "nombre1_actualizado", "apellido1_actualizado",  "mail1_actualizado@mail.com", 23, 1.85, true, "25/07/2022")

	//assert
	assert.True(t, mockStorage.ReadWasCalled)
	assert.Equal(t, domain.User{}, result)
	assert.ErrorContains(t, err, expectedError.Error())

}

func TestIntegrationDelete(t *testing.T){
	// arrange
	dataBase := []domain.User{
		{Id: 1, Name: "nombre1", LastName: "apellido1", Email: "mail1@mail.com", Age: 22, Height:1.83, Active: true, CreatedAt: "25/07/2022"},
		{Id: 2, Name: "nombre2", LastName: "apellido2", Email: "mail2@mail.com", Age: 23, Height:1.60, Active: true, CreatedAt: "25/07/2022"},
	}

	mockStorage := MockStore{
		dataMock: dataBase,
		errRead: "",
		errWrite: "",
		ReadWasCalled: false,
	}
	
	// act
    repo := NewRepository(&mockStorage)
	service := NewService(repo)
	
	err := service.Delete(1)

	//assert
	assert.True(t, mockStorage.ReadWasCalled)
	assert.Nil(t, err)
	assert.NotEqual(t, mockStorage.dataMock[0].Id, 1)

}

func TestIntegrationDeleteFail(t *testing.T){
	// arrange
	dataBase := []domain.User{
		{Id: 1, Name: "nombre1", LastName: "apellido1", Email: "mail1@mail.com", Age: 22, Height:1.83, Active: true, CreatedAt: "25/07/2022"},
		{Id: 2, Name: "nombre2", LastName: "apellido2", Email: "mail2@mail.com", Age: 23, Height:1.60, Active: true, CreatedAt: "25/07/2022"},
	}

	mockStorage := MockStore{
		dataMock: dataBase,
		errRead: "",
		errWrite: "",
		ReadWasCalled: false,
	}
	deleteError := fmt.Errorf("user 33 not found")
	expectedError := fmt.Errorf("error deleting user: %w", deleteError)
	
	// act
    repo := NewRepository(&mockStorage)
	service := NewService(repo)
	
	err := service.Delete(33)

	//assert
	assert.True(t, mockStorage.ReadWasCalled)
	assert.Equal(t, expectedError, err)
}