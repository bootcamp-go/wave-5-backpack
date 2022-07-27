package users

import (
	"clase2_2/internal/domain"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockStorage struct {
	dataMock     []domain.User
	errWrite     string
	errRead      string
	isReadCalled bool
}

func (m *MockStorage) Read(data interface{}) error {
	if m.errRead != "" {
		return fmt.Errorf(m.errRead)
	}
	a := data.(*[]domain.User)
	*a = m.dataMock
	m.isReadCalled = true
	return nil
}

func (m *MockStorage) Write(data interface{}) error {
	if m.errWrite != "" {
		return fmt.Errorf(m.errWrite)
	}
	a := data.([]domain.User)
	m.dataMock = append(m.dataMock, a[len(a)-1])
	return nil
}

// ejercicio1
func TestIntegrationUpdate(t *testing.T) {
	dataDB := []domain.User{
		{Id: 1, Name: "nombre1", LastName: "apellido1", Mail: "mail1", Years: 25, Tall: 1, Enable: true, CreateDate: "0/0/0"},
		{Id: 2, Name: "nombre2", LastName: "apellido2", Mail: "mail2", Years: 20, Tall: 2, Enable: false, CreateDate: "10/10/10"},
	}
	mockStorage := MockStorage{
		dataMock:     dataDB,
		errWrite:     "",
		errRead:      "",
		isReadCalled: false,
	}
	esperado := domain.User{Id: 1, Name: "nombre10", LastName: "apellido10", Mail: "mail10", Years: 20, Tall: 10, Enable: true, CreateDate: "10/10/10"}

	repo := NewRepository(&mockStorage)
	serv := NewService(repo)

	result, err := serv.UpdateUser("nombre10", "apellido10", "mail10", "10/10/10", 20, 1, 10, true)

	assert.Nil(t, err)
	assert.True(t, mockStorage.isReadCalled, "no se a ejecutado el método")
	assert.Equal(t, esperado, result, "No son iguales")
}

func TestIntegrationDelete(t *testing.T) {
	dataDB := []domain.User{
		{Id: 1, Name: "nombre1", LastName: "apellido1", Mail: "mail1", Years: 25, Tall: 1, Enable: true, CreateDate: "0/0/0"},
		{Id: 2, Name: "nombre2", LastName: "apellido2", Mail: "mail2", Years: 20, Tall: 2, Enable: false, CreateDate: "10/10/10"},
	}
	mockStorage := MockStorage{
		dataMock:     dataDB,
		errWrite:     "",
		errRead:      "",
		isReadCalled: false,
	}

	repo := NewRepository(&mockStorage)
	serv := NewService(repo)

	err := serv.Delete(1)

	assert.True(t, mockStorage.isReadCalled)
	assert.Nil(t, err)

}
func TestIntegrationDeleteFail(t *testing.T) {
	dataDB := []domain.User{
		{Id: 1, Name: "nombre1", LastName: "apellido1", Mail: "mail1", Years: 25, Tall: 1, Enable: true, CreateDate: "0/0/0"},
		{Id: 2, Name: "nombre2", LastName: "apellido2", Mail: "mail2", Years: 20, Tall: 2, Enable: false, CreateDate: "10/10/10"},
	}
	mockStorage := MockStorage{
		dataMock:     dataDB,
		errWrite:     "",
		errRead:      "",
		isReadCalled: false,
	}
	errorEsperado := errors.New("no se encontró el producto de id 2")

	repo := NewRepository(&mockStorage)
	serv := NewService(repo)

	err := serv.Delete(1)

	assert.Equal(t, err, errorEsperado)

}
