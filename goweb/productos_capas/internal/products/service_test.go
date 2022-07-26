package products

import (
	"fmt"
	"goweb/productos_capas/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServiceIntegrationGetAll(t *testing.T) {
	//arrange
	database := []domain.Product{
		{Id: 1, Nombre: "Before Update", Color: "Negro", Precio: 100, Stock: 10, Codigo: "000", Publicado: true, FechaCreacion: "02-11-1999"},
		{Id: 2, Nombre: "Televisor", Color: "Azul", Precio: 200, Stock: 5, Codigo: "111", Publicado: false, FechaCreacion: "12-04-1989"},
	}
	mockStorage := MockStorage{
		dataMock: database,
		errWrite: "",
		errRead:  "",
	}

	//act
	repository := NewRepository(&mockStorage)
	service := NewService(repository)
	resultado, err := service.GetAll("", "", 0, 0, "", false, "")

	//assert
	assert.Equal(t, mockStorage.dataMock, resultado, "deben ser iguales")
	assert.Equal(t, nil, err, "deben ser iguales")
}

func TestServiceIntegrationGetAllFail(t *testing.T) {
	//arrange
	mockStorage := MockStorage{
		dataMock: nil,
		errWrite: "",
		errRead:  "cant read database",
	}
	expectedError := fmt.Errorf("cant read database")

	//act
	repository := NewRepository(&mockStorage)
	service := NewService(repository)
	result, err := service.GetAll("", "", 0, 0, "", false, "")

	//assert
	assert.Equal(t, expectedError, err, "deben ser iguales")
	assert.Nil(t, result)
}

func TestServiceIntegrationStore(t *testing.T) {
	//arrange
	newProduct := domain.Product{Id: 1, Nombre: "Before Update", Color: "Negro", Precio: 100, Stock: 10, Codigo: "000", Publicado: true, FechaCreacion: "02-11-1999"}
	mockStorage := MockStorage{
		dataMock: nil,
		errWrite: "",
		errRead:  "",
	}

	//act
	repository := NewRepository(&mockStorage)
	service := NewService(repository)
	result, err := service.Store(newProduct.Nombre, newProduct.Color, newProduct.Precio, newProduct.Stock, newProduct.Codigo, newProduct.Publicado, newProduct.FechaCreacion)

	//assert
	assert.Nil(t, err)
	assert.Equal(t, mockStorage.dataMock[0], result)
	assert.Equal(t, mockStorage.dataMock[0].Id, 1)
}

func TestServiceIntegrationStoreFail(t *testing.T) {
	//arrange
	newProduct := domain.Product{Id: 1, Nombre: "Before Update", Color: "Negro", Precio: 100, Stock: 10, Codigo: "000", Publicado: true, FechaCreacion: "02-11-1999"}
	mockStorage := MockStorage{
		dataMock: nil,
		errWrite: "cant write database",
		errRead:  "",
	}
	writeErr := fmt.Errorf("cant write database")
	expectedError := fmt.Errorf("cant write database, error: %w", writeErr)

	//act
	repository := NewRepository(&mockStorage)
	service := NewService(repository)
	result, err := service.Store(newProduct.Nombre, newProduct.Color, newProduct.Precio, newProduct.Stock, newProduct.Codigo, newProduct.Publicado, newProduct.FechaCreacion)

	//assert
	assert.Equal(t, expectedError, err)
	assert.Equal(t, domain.Product{}, result)
}

func TestServiceIntegrationUpdate(t *testing.T) {
	//arrange
	beforeUpdate := domain.Product{Id: 1, Nombre: "Before Update", Color: "Negro", Precio: 100, Stock: 10, Codigo: "000", Publicado: true, FechaCreacion: "02-11-1999"}
	afterUpdate := domain.Product{Id: 1, Nombre: "After Update", Color: "Blaco", Precio: 200, Stock: 20, Codigo: "111", Publicado: false, FechaCreacion: "03-12-2000"}

	mockStorage := MockStorage{
		dataMock: []domain.Product{beforeUpdate},
		errWrite: "",
		errRead:  "",
	}

	//act
	repository := NewRepository(&mockStorage)
	service := NewService(repository)
	result, err := service.Update(beforeUpdate.Id, afterUpdate.Nombre, afterUpdate.Color, afterUpdate.Precio, afterUpdate.Stock, afterUpdate.Codigo, afterUpdate.Publicado, afterUpdate.FechaCreacion)

	//assert
	assert.Equal(t, mockStorage.dataMock[0], result)
	assert.Equal(t, mockStorage.dataMock[0].Id, 1)
	assert.Nil(t, err)
	assert.True(t, mockStorage.ReadWasCalled)
}

func TestServiceIntegrationUpdateFail(t *testing.T) {
	//arrange
	beforeUpdate := domain.Product{Id: 1, Nombre: "Before Update", Color: "Negro", Precio: 100, Stock: 10, Codigo: "000", Publicado: true, FechaCreacion: "02-11-1999"}
	afterUpdate := domain.Product{Id: 1, Nombre: "After Update", Color: "Blaco", Precio: 200, Stock: 20, Codigo: "111", Publicado: false, FechaCreacion: "03-12-2000"}

	mockStorage := MockStorage{
		dataMock: []domain.Product{beforeUpdate},
		errWrite: "cant write database",
		errRead:  "",
	}
	writeErr := fmt.Errorf("cant write database")
	expectedError := fmt.Errorf("cant write database, error: %w", writeErr)

	//act
	repository := NewRepository(&mockStorage)
	service := NewService(repository)
	result, err := service.Update(beforeUpdate.Id, afterUpdate.Nombre, afterUpdate.Color, afterUpdate.Precio, afterUpdate.Stock, afterUpdate.Codigo, afterUpdate.Publicado, afterUpdate.FechaCreacion)

	//assert
	assert.Equal(t, expectedError, err)
	assert.Equal(t, domain.Product{}, result)
	assert.True(t, mockStorage.ReadWasCalled)
}

func TestServiceIntegrationDelete(t *testing.T) {
	//arrange
	existingProduct := domain.Product{Id: 1, Nombre: "Delete", Color: "Negro", Precio: 100, Stock: 10, Codigo: "000", Publicado: true, FechaCreacion: "02-11-1999"}
	mockStorage := MockStorage{
		dataMock: []domain.Product{existingProduct},
		errWrite: "",
		errRead:  "",
	}

	//act
	repository := NewRepository(&mockStorage)
	service := NewService(repository)
	_, err := service.Delete(1)

	//assert
	assert.Nil(t, err)
	assert.Equal(t, mockStorage.dataMock, []domain.Product{})
	assert.True(t, mockStorage.ReadWasCalled)
}

func TestServiceIntegrationDeleteFail(t *testing.T) {
	//arrange
	existingProduct := domain.Product{Id: 1, Nombre: "Delete", Color: "Negro", Precio: 100, Stock: 10, Codigo: "000", Publicado: true, FechaCreacion: "02-11-1999"}
	mockStorage := MockStorage{
		dataMock: []domain.Product{existingProduct},
		errWrite: "",
		errRead:  "",
	}
	expectedError := fmt.Errorf("product %d not found", 2)

	//act
	repository := NewRepository(&mockStorage)
	service := NewService(repository)
	result, err := service.Delete(2)

	//assert
	assert.Equal(t, expectedError, err)
	assert.Equal(t, domain.Product{}, result)
	assert.True(t, mockStorage.ReadWasCalled)
}
