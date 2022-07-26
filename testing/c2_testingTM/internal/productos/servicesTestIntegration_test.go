package productos

import (
	"errors"
	"fmt"
	"testing"

	"github.com/bootcamp-go/wave-5-backpack/tree/ospina_christian/testing/c2_testingTM/internal/domain"
	"github.com/stretchr/testify/assert"
)

//Integration Test (Update) services / Repository
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

//Integration Test (Delete) services / Repository
func TestServiceIntegrationDelete(t *testing.T) {
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
	err := service.Delete(1)
	pro, _ := service.GetForId(1)

	//assert
	assert.Nil(t, err)
	assert.Equal(t, domain.Productos{}, pro)
}

func TestServiceIntegrationDeleteFail(t *testing.T) {
	//arrange
	errorAdd := errors.New("cant read database")
	expectedError := fmt.Errorf("error deleting el producto %w", errorAdd)
	mockStorage := MockStorage{
		dataMock:   nil,
		errWrite:   "",
		errRead:    "cant read database",
		readUpdate: false,
	}

	//act
	repo := NewRepository(&mockStorage)
	service := NewService(repo)
	err := service.Delete(1)

	//assert
	assert.Equal(t, expectedError, err)
}

// ------------------**------------------

//Integration Test (Store) services / Repository
func TestServiceIntegrationStore(t *testing.T) {
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
	prod, err := service.Store("Papaya", "Verde", 2100, 23, "as1", true, "26/07/2022")
	pro, _ := service.GetForId(3)

	//assert
	assert.Nil(t, err)
	assert.Equal(t, pro, prod)
}

//Integration Test (GetAll) services / Repository
func TestServiceIntegrationGetAll(t *testing.T) {
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
	prod, err := service.GetAll()

	//assert
	assert.Nil(t, err)
	assert.Equal(t, database, prod)
}

func TestServiceIntegrationGetAllFail(t *testing.T) {
	//arrange
	expectedError := errors.New("cant read database")
	mockStorage := MockStorage{
		dataMock:   nil,
		errWrite:   "",
		errRead:    "cant read database",
		readUpdate: false,
	}

	//act
	repo := NewRepository(&mockStorage)
	service := NewService(repo)
	result, err := service.GetAll()

	//assert
	assert.Equal(t, expectedError, err)
	assert.Nil(t, result)
}

// ------------------**------------------

//Integration Test (GetId) services / Repository
func TestServiceIntegrationGetId(t *testing.T) {
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
	prod, err := service.GetForId(2)

	//assert
	assert.Nil(t, err)
	assert.Equal(t, database[1], prod)
}

func TestServiceIntegrationGetIdFail(t *testing.T) {
	//arrange
	errorAdd := errors.New("cant read database")
	expectedError := fmt.Errorf("error al obtener el producto %w", errorAdd)
	mockStorage := MockStorage{
		dataMock:   nil,
		errWrite:   "",
		errRead:    "cant read database",
		readUpdate: false,
	}

	//act
	repo := NewRepository(&mockStorage)
	service := NewService(repo)
	prod, err := service.GetForId(1)

	//assert
	assert.Equal(t, expectedError, err)
	assert.Equal(t, domain.Productos{}, prod)
}

// ------------------**------------------

//Integration Test (UpdatePrecio) services / Repository
func TestServiceIntegrationUpdatePrecio(t *testing.T) {
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
	prod, err := service.UpdatePrecio(2, 5000)

	//assert
	assert.Nil(t, err)
	assert.Equal(t, database[1].Precio, prod.Precio)
}
