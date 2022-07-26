package products

import (
	"fmt"
	"goweb/clase1_clase2/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServiceIntegrationGetAll(t *testing.T) {
	// arrange
	database := []domain.Product{
		{
			Id:        1,
			Nombre:    "Nevera",
			Color:     "Blanco",
			Precio:    600,
			Stock:     4,
			Codigo:    "B453",
			Publicado: true,
			Fecha:     "05-05-2022",
		},
		{
			Id:        1,
			Nombre:    "Televisor",
			Color:     "Verde",
			Precio:    400,
			Stock:     9,
			Codigo:    "T543",
			Publicado: true,
			Fecha:     "02-06-2022",
		},
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
	expectedError := fmt.Errorf("cant read database")
	mockStorage := MockStorage{
		dataMock: nil,
		errWrite: "",
		errRead:  "cant read database",
	}
	// act
	repo := NewRepository(&mockStorage)
	service := NewService(repo)
	result, err := service.GetAll("", "", 0, 0, "", false, "")
	// assert
	assert.Equal(t, expectedError, err)
	assert.Nil(t, result)
}

func TestServiceIntegrationStore(t *testing.T) {
	// arrange
	newProduct := domain.Product{
		Id:        1,
		Nombre:    "Nevera",
		Color:     "Blanco",
		Precio:    600,
		Stock:     4,
		Codigo:    "B453",
		Publicado: true,
		Fecha:     "05-05-2022",
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
		newProduct.Nombre,
		newProduct.Color,
		newProduct.Precio,
		newProduct.Stock,
		newProduct.Codigo,
		newProduct.Publicado,
		newProduct.Fecha)
	// assert
	assert.Nil(t, err)
	assert.Equal(t, mockStorage.dataMock[0], result)
	assert.Equal(t, mockStorage.dataMock[0].Id, 1)
}

func TestServiceIntegrationStoreFail(t *testing.T) {
	// arrange
	newProduct := domain.Product{
		Id:        1,
		Nombre:    "Nevera",
		Color:     "Blanco",
		Precio:    600,
		Stock:     4,
		Codigo:    "B453",
		Publicado: true,
		Fecha:     "05-05-2022",
	}
	writeErr := fmt.Errorf("cant write database")
	expectedError := fmt.Errorf("cant write database, error: %w", writeErr)
	mockStorage := MockStorage{
		dataMock: nil,
		errWrite: "cant write database",
		errRead:  "",
	}
	// act
	repo := NewRepository(&mockStorage)
	service := NewService(repo)
	result, err := service.Store(
		newProduct.Nombre,
		newProduct.Color,
		newProduct.Precio,
		newProduct.Stock,
		newProduct.Codigo,
		newProduct.Publicado,
		newProduct.Fecha)
	// assert
	assert.Equal(t, expectedError, err)
	assert.Equal(t, domain.Product{}, result)
}

func TestServiceIntegrationUpdate(t *testing.T) {
	// arrange
	updateProduct := domain.Product{
		Id:        1,
		Nombre:    "Lavadora",
		Color:     "Blanco",
		Precio:    600,
		Stock:     4,
		Codigo:    "B453",
		Publicado: true,
		Fecha:     "05-05-2022",
	}
	mockStorage := MockStorage{
		dataMock: []domain.Product{updateProduct},
		errWrite: "",
		errRead:  "",
	}
	// act
	repo := NewRepository(&mockStorage)
	service := NewService(repo)
	result, err := service.Update(
		updateProduct.Id,
		updateProduct.Nombre,
		updateProduct.Color,
		updateProduct.Precio,
		updateProduct.Stock,
		updateProduct.Codigo,
		updateProduct.Publicado,
		updateProduct.Fecha)
	// assert
	assert.Nil(t, err)
	assert.True(t, mockStorage.ReadWasCalled)
	assert.Equal(t, mockStorage.dataMock[0], result)
	assert.Equal(t, mockStorage.dataMock[0].Id, 1)
}

func TestServiceIntegrationUpdateFail(t *testing.T) {
	// arrange
	updateProduct := domain.Product{
		Id:        1,
		Nombre:    "Lavadora",
		Color:     "Blanco",
		Precio:    600,
		Stock:     4,
		Codigo:    "B453",
		Publicado: true,
		Fecha:     "05-05-2022",
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
		updateProduct.Id,
		updateProduct.Nombre,
		updateProduct.Color,
		updateProduct.Precio,
		updateProduct.Stock,
		updateProduct.Codigo,
		updateProduct.Publicado,
		updateProduct.Fecha)
	// assert
	assert.True(t, mockStorage.ReadWasCalled)
	assert.Equal(t, expectedError, err)
	assert.Equal(t, domain.Product{}, result)
}

func TestServiceIntegrationDelete(t *testing.T) {
	// arrange
	deleteProduct := domain.Product{
		Id:        1,
		Nombre:    "Lavadora",
		Color:     "Blanco",
		Precio:    600,
		Stock:     4,
		Codigo:    "B453",
		Publicado: true,
		Fecha:     "05-05-2022",
	}
	mockStorage := MockStorage{
		dataMock: []domain.Product{deleteProduct},
		errWrite: "",
		errRead:  "",
	}
	// act
	repo := NewRepository(&mockStorage)
	service := NewService(repo)
	_, err := service.Delete(deleteProduct.Id)
	// assert
	assert.Nil(t, err)
	assert.Equal(t, mockStorage.dataMock, []domain.Product{})
}

func TestServiceIntegrationDeleteFail(t *testing.T) {
	// arrange
	deleteProduct := domain.Product{
		Id:        1,
		Nombre:    "Lavadora",
		Color:     "Blanco",
		Precio:    600,
		Stock:     4,
		Codigo:    "B453",
		Publicado: true,
		Fecha:     "05-05-2022",
	}
	expectedError := fmt.Errorf("product %d not found", 2)
	mockStorage := MockStorage{
		dataMock: []domain.Product{deleteProduct},
		errWrite: "",
		errRead:  "",
	}
	// act
	repo := NewRepository(&mockStorage)
	service := NewService(repo)
	result, err := service.Delete(2)
	// assert
	assert.Equal(t, expectedError, err)
	assert.Equal(t, domain.Product{}, result)
}
