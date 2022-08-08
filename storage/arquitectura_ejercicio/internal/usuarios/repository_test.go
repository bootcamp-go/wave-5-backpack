package usuarios

import (
	"fmt"
	"testing"

	"github.com/anesquivel/wave-5-backpack/storage/arquitectura_ejercicio/internal/domain"
	"github.com/stretchr/testify/assert"
)

type MockStorage struct {
	dataMock       []domain.Usuario
	errRead        string
	errWrite       string
	hadCalledRead  bool
	hadCalledWrite bool
}

func (m *MockStorage) Read(data interface{}) error {
	if m.errRead != "" {
		return fmt.Errorf(m.errRead)
	}

	user := data.(*[]domain.Usuario)
	*user = m.dataMock

	m.hadCalledRead = true
	return nil
}

func (m *MockStorage) Write(data interface{}) error {

	if m.errWrite != "" {
		return fmt.Errorf(m.errWrite)
	}

	a := data.([]domain.Usuario)
	m.dataMock = append(m.dataMock, a...)

	m.hadCalledWrite = true
	return nil

}

func (m *MockStorage) Ping() error {
	if m.errRead != "" {
		return fmt.Errorf(m.errRead)
	}
	return nil
}

func TestGetall(t *testing.T) {
	// arrange
	database := []domain.Usuario{
		{
			Id:          1,
			Names:       "Andy",
			LastName:    "Esquivel",
			Age:         23,
			Estatura:    1.52,
			Email:       "andy@gmail.com",
			IsActivo:    true,
			DateCreated: "25-07-2022",
		},
		{
			Id:          2,
			Names:       "Gabriela",
			LastName:    "Rueda",
			Age:         23,
			Estatura:    1.52,
			Email:       "gr@gmail.com",
			IsActivo:    true,
			DateCreated: "26-07-2022",
		},
	}

	mockStorage := MockStorage{
		dataMock: database,
		errRead:  "",
		errWrite: "",
	}

	// act
	repo := NewRepository(&mockStorage)
	res, err := repo.GetAll()

	// assert

	assert.Nil(t, err)
	assert.Equal(t, mockStorage.dataMock, res)
}
