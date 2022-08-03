package repositorio

import (
	"C3-Testing/internal/domain"

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

func TestGetAll(t *testing.T) {
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

	repo := NewRepository(&mockStorage)
	result, err := repo.GetAll()
	//assert
	assert.Nil(t, err)
	assert.Equal(t, mockStorage.dataMock, result)

}
