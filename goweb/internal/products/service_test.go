package products

import (
	"fmt"
	"goweb/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

// ========================================================
// --------------------- MockStorage ----------------------
// ========================================================

type MockStorage struct {
	dataMock       []domain.Product
	ReadWasCalled  bool
	WriteWasCalled bool
	errWrite       string
	errRead        string
}

func (m *MockStorage) Write(data interface{}) error {
	if m.errWrite != "" {
		return fmt.Errorf(m.errWrite)
	}
	productos := data.(*[]domain.Product)
	//m.dataMock = append(m.dataMock, *productos...)
	m.dataMock = *productos

	m.WriteWasCalled = true

	return nil
}

func (m *MockStorage) Read(data interface{}) error {
	if m.errRead != "" {
		return fmt.Errorf(m.errRead)
	}
	productos := data.(*[]domain.Product)
	*productos = m.dataMock

	m.ReadWasCalled = true

	return nil
}

func (m *MockStorage) Ping() error {
	return nil
}

// ========================================================
// ------------------------- Tests ------------------------
// ========================================================

func TestUpdate(t *testing.T) {

	// -------------------------------------------------------
	// Se inicializan los datos a usar en el test (input y output)
	// -------------------------------------------------------

	// Update values
	updateId := 1
	updateName := "After Update"
	updateColor := "negro"
	updatePrecio := 200.0
	updateStock := 2
	updateCodigo := "123"
	updatePublicado := true
	updateFechaCreacion := "2020-01-01"
	// Database
	database := []domain.Product{
		{Id: 1, Nombre: "Before Update", Color: "rojo", Precio: 10, Stock: 10, Codigo: "123", Publicado: true, FechaCreacion: "2020-01-01"},
		{Id: 2, Nombre: "Before Update", Color: "verde", Precio: 10, Stock: 10, Codigo: "1234", Publicado: true, FechaCreacion: "2019-01-01"},
	}
	mockStorage := MockStorage{
		dataMock: database,
		errWrite: "",
		errRead:  "",
	}
	// Capas
	repository := NewRepositoryJsonCorrDB(&mockStorage)
	service := NewService(repository)

	// -------------------------------------------------------
	// Se ejecuta el test
	// -------------------------------------------------------

	result, err := service.Update(updateId, updateName, updateColor, updatePrecio, updateStock, updateCodigo, updatePublicado, updateFechaCreacion)

	// -------------------------------------------------------
	// Se compara el resultado obtenido con el resultado esperado
	// -------------------------------------------------------

	assert.Nil(t, err)
	assert.Equal(t, updateId, result.Id)
	assert.Equal(t, updateName, result.Nombre)
	assert.Equal(t, updateColor, result.Color)
	assert.Equal(t, updatePrecio, result.Precio)
	assert.Equal(t, updateStock, result.Stock)
	assert.Equal(t, updateCodigo, result.Codigo)
	assert.Equal(t, updatePublicado, result.Publicado)
	assert.Equal(t, updateFechaCreacion, result.FechaCreacion)
	assert.True(t, mockStorage.ReadWasCalled)

}

func TestDelete(t *testing.T) {

	// -------------------------------------------------------
	// Se inicializan los datos a usar en el test (input y output)
	// -------------------------------------------------------

	// Delete id
	deleteId := 1
	// Database
	database := []domain.Product{
		{Id: 1, Nombre: "Before Delete", Color: "rojo", Precio: 10, Stock: 10, Codigo: "123", Publicado: true, FechaCreacion: "2020-01-01"},
		{Id: 2, Nombre: "Before Delete", Color: "verde", Precio: 10, Stock: 10, Codigo: "1234", Publicado: true, FechaCreacion: "2019-01-01"},
	}
	mockStorage := MockStorage{
		dataMock: database,
		errWrite: "",
		errRead:  "",
	}
	// Capas
	repository := NewRepositoryJsonCorrDB(&mockStorage)
	service := NewService(repository)

	// -------------------------------------------------------
	// Se ejecuta el test
	// -------------------------------------------------------

	err := service.Delete(deleteId)

	// -------------------------------------------------------
	// Se compara el resultado obtenido con el resultado esperado
	// -------------------------------------------------------

	assert.Nil(t, err)
	assert.Equal(t, 1, len(mockStorage.dataMock))
	assert.True(t, mockStorage.ReadWasCalled)
	assert.True(t, mockStorage.WriteWasCalled)
}

func TestDeleteFail(t *testing.T) {

	// -------------------------------------------------------
	// Se inicializan los datos a usar en el test (input y output)
	// -------------------------------------------------------

	// Delete id
	deleteId := 3
	// Database
	database := []domain.Product{
		{Id: 1, Nombre: "Before Delete", Color: "rojo", Precio: 10, Stock: 10, Codigo: "123", Publicado: true, FechaCreacion: "2020-01-01"},
		{Id: 2, Nombre: "Before Delete", Color: "verde", Precio: 10, Stock: 10, Codigo: "1234", Publicado: true, FechaCreacion: "2019-01-01"},
	}
	mockStorage := MockStorage{
		dataMock: database,
		errWrite: "",
		errRead:  "",
	}
	// Error
	ErrProductNotFound := fmt.Errorf(ProductNotFound, deleteId)
	expectedError := fmt.Errorf("error eliminando el producto con el id %d: %w", deleteId, ErrProductNotFound)
	// Capas
	repository := NewRepositoryJsonCorrDB(&mockStorage)
	service := NewService(repository)

	// -------------------------------------------------------
	// Se ejecuta el test
	// -------------------------------------------------------

	err := service.Delete(deleteId)

	// -------------------------------------------------------
	// Se compara el resultado obtenido con el resultado esperado
	// -------------------------------------------------------

	assert.Equal(t, expectedError, err)
	assert.Equal(t, 2, len(mockStorage.dataMock))
}
