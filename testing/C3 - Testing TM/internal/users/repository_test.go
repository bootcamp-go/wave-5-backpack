package users

import (
	"C3/internal/domain"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockStorage struct {
	ReadWasCalled bool
	dataMock      []domain.Usuarios
	errWrite      string
	errRead       string
}

func (m *MockStorage) Write(data interface{}) error {
	if m.errWrite != "" {
		return fmt.Errorf(m.errWrite)
	}
	a := data.([]domain.Usuarios)
	m.dataMock = a
	return nil
}

func (m *MockStorage) Read(data interface{}) error {
	m.ReadWasCalled = true
	if m.errRead != "" {
		return fmt.Errorf(m.errRead)
	}
	a := data.(*[]domain.Usuarios)
	*a = m.dataMock
	return nil
}

func (m *MockStorage) Ping() error {
	return nil
}

func TestMock(t *testing.T) {
	//arrange
	database := []domain.Usuarios{
		{Id: 2, Nombre: "Marcela", Apellido: "Monroy", Email: "marcela@hotmail.com", Edad: 27, Altura: 1.67},
		{Id: 3, Nombre: "Marcelo", Apellido: "Moncada", Email: "marcelo@hotmail.com", Edad: 20, Altura: 1.82},
	}
	mockStorage := MockStorage{
		dataMock: database,
		errWrite: "",
		errRead:  "",
	}
	expLast := "Martinez"
	expAge := 28
	//act
	repo := NewRepository(&mockStorage)
	result, err := repo.UpdateLastAge(2, expLast, expAge)

	assert.Equal(t, true, mockStorage.ReadWasCalled)
	assert.Equal(t, expLast, result.Apellido)
	assert.Equal(t, expAge, result.Edad)
	assert.Equal(t, nil, err)
}

func TestGetAll(t *testing.T) {
	//arrange
	database := []domain.Usuarios{
		{Id: 2, Nombre: "Marcela", Apellido: "Monroy", Email: "marcela@hotmail.com", Edad: 27, Altura: 1.67},
		{Id: 3, Nombre: "Marcelo", Apellido: "Moncada", Email: "marcelo@hotmail.com", Edad: 20, Altura: 1.82},
	}
	mockStorage := MockStorage{
		dataMock: database,
		errWrite: "",
		errRead:  "",
	}

	//act
	repo := NewRepository(&mockStorage)
	result, err := repo.GetAll()

	//assert
	assert.Nil(t, err)
	assert.Equal(t, mockStorage.dataMock, result)
}
