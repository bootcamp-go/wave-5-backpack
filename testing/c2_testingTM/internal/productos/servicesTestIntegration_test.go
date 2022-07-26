package productos

import (
	"testing"

	"github.com/bootcamp-go/wave-5-backpack/tree/ospina_christian/testing/c2_testingTM/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestServiceIntegrationUpdate(t *testing.T) {
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
		{
			Id:            2,
			Nombre:        "Alcaparras",
			Color:         "Verde",
			Precio:        1230,
			Stock:         12,
			Codigo:        "@323",
			Publicado:     true,
			FechaCreación: "12/08/2022",
		},
	}
	mockStorage := MockStorage{
		dataMock:   database,
		errWrite:   "",
		errRead:    "",
		readUpdate: false,
	}

	//act
	repo := NewRepository(&mockStorage)
	service := NewService(repo)
	result, err := service.Update(
		database[0].Id,
		database[0].Nombre,
		database[0].Color,
		database[0].Precio,
		database[0].Stock,
		database[0].Codigo,
		database[0].Publicado,
		database[0].FechaCreación)
	if err != nil {
		mockStorage.readUpdate = false
	}
	mockStorage.readUpdate = true
	//assert
	assert.Nil(t, err)
	assert.True(t, mockStorage.readUpdate)
	assert.Equal(t, mockStorage.dataMock[0], result)
	assert.Equal(t, mockStorage.dataMock[0].Id, 1)
}

func TestServiceIntegrationDelete(t *testing.T) {
	//arrange
	//writeErr := fmt.Errorf("cant read database")
	//expectError := fmt.Errorf("error deleting el producto %w", writeErr)
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
		{
			Id:            2,
			Nombre:        "Alcaparras",
			Color:         "Verde",
			Precio:        1230,
			Stock:         12,
			Codigo:        "@323",
			Publicado:     true,
			FechaCreación: "12/08/2022",
		},
	}
	mockStorage := MockStorage{
		dataMock:   database,
		errWrite:   "",
		errRead:    "",
		readUpdate: false,
	}

	//act
	repo := NewRepository(&mockStorage)
	service := NewService(repo)
	err := service.Delete(1)
	pro, _ := service.GetForId(1)

	//assert
	assert.Nil(t, err)
	// assert.Nil(t, errq)
	assert.Equal(t, domain.Productos{}, pro)
}
