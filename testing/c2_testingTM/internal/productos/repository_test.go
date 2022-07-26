package productos

import (
	"fmt"
	"testing"

	"github.com/bootcamp-go/wave-5-backpack/tree/ospina_christian/testing/c2_testingTM/internal/domain"
	"github.com/stretchr/testify/assert"
)

type MockStorage struct {
	dataMock   []domain.Productos
	errWrite   string
	errRead    string
	readUpdate bool
}

func (m *MockStorage) Read(data interface{}) error {
	if m.errRead != "" {
		return fmt.Errorf(m.errRead)
	}
	a := data.(*[]domain.Productos)
	*a = m.dataMock
	return nil
}

func (m *MockStorage) Write(data interface{}) error {
	if m.errWrite != "" {
		return fmt.Errorf(m.errWrite)
	}
	a := data.([]domain.Productos)
	m.dataMock = append(m.dataMock, a[len(a)-1])
	return nil
}

func (m *MockStorage) Ping() error {
	if m.errRead != "" {
		return fmt.Errorf(m.errRead)
	}
	return nil
}

func TestGetAll(t *testing.T) {
	//arrange
	database := []domain.Productos{
		{
			Id:            1,
			Nombre:        "Esparragos",
			Color:         "Verde",
			Precio:        12300,
			Stock:         12,
			Codigo:        "@123",
			Publicado:     true,
			FechaCreación: "12/08/2022",
		},
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
