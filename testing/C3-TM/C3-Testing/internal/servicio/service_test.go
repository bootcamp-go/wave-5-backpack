package servicio

import (
	"C3-Testing/internal/domain"
	"C3-Testing/internal/repositorio"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockStorage struct {
	dataMock []domain.User
	errWrite string
	errRead  string
}

func (m *MockStorage) Read(data interface{}) error {
	if m.errRead != "" {
		return fmt.Errorf(m.errRead)
	}
	a := data.(*[]domain.User)
	*a = m.dataMock
	return nil
}

func (m *MockStorage) Open(data interface{}) error {
	if m.errRead != "" {
		return fmt.Errorf(m.errRead)
	}
	return nil
}

func (m *MockStorage) Write(data interface{}) error {
	if m.errRead != "" {
		return fmt.Errorf(m.errRead)
	}
	a := data.([]domain.User)
	m.dataMock = append(m.dataMock, a[len(a)-1])
	return nil
}

func TestServiceIntegrationGetAll(t *testing.T) {
	//arrange
	database := []domain.User{
		{
			Id:        1,
			FirstName: "Roko",
			LastName:  "Moonrock",
			Email:     "roko@dogmail.com",
			Age:       4,
			Height:    1.20,
			Activo:    true,
			CreatedAt: "11/08/1996",
		},
		{
			Id:        2,
			FirstName: "Luna",
			LastName:  "Moonrock",
			Email:     "roko@dogmail.com",
			Age:       4,
			Height:    1.20,
			Activo:    true,
			CreatedAt: "11/08/1996",
		},
	}

	mockStorage := MockStorage{
		dataMock: database,
		errWrite: "",
		errRead:  "",
	}

	//act
	repo := repositorio.NewRepository(&mockStorage)
	service := NewService(repo)
	result, err := service.GetAll()
	//assert
	assert.Nil(t, err)
	assert.Equal(t, mockStorage.dataMock, result)
}

func TestGetAllFail(t *testing.T) {
	//arrange
	expectedError := errors.New("cant read database")
	mockStorage := MockStorage{
		dataMock: nil,
		errWrite: "cant read database",
		errRead:  "",
	}
	//act
	repo := repositorio.NewRepository(&mockStorage)
	service := NewService(repo)
	result, err := service.GetAll()
	//assert
	assert.ErrorContains(t, err, expectedError.Error())
	assert.Nil(t, result)
}
