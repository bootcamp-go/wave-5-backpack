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
	updateColor := "gris"
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
	repository := NewRepository(&mockStorage)
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

func TestUpdateFail(t *testing.T) {

	// -------------------------------------------------------
	// Se inicializan los datos a usar en el test (input y output)
	// -------------------------------------------------------

	// Update values
	updateId := 3
	updateName := "After Update"
	updateColor := "gris"
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
	// Error
	ErrProductNotFound := fmt.Errorf(ProductNotFound, updateId)
	expectedError := fmt.Errorf("error actualizando el producto con el id %d: %w", updateId, ErrProductNotFound)
	// Capas
	repository := NewRepository(&mockStorage)
	service := NewService(repository)

	// -------------------------------------------------------
	// Se ejecuta el test
	// -------------------------------------------------------

	_, err := service.Update(updateId, updateName, updateColor, updatePrecio, updateStock, updateCodigo, updatePublicado, updateFechaCreacion)

	// -------------------------------------------------------
	// Se compara el resultado obtenido con el resultado esperado
	// -------------------------------------------------------

	assert.Equal(t, expectedError, err)
	assert.Equal(t, 2, len(mockStorage.dataMock))
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
	repository := NewRepository(&mockStorage)
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
	repository := NewRepository(&mockStorage)
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

func TestGetById(t *testing.T) {

	// -------------------------------------------------------
	// Se inicializan los datos a usar en el test (input y output)
	// -------------------------------------------------------

	// Get id
	getId := 2
	resultadoEsperado := domain.Product{
		Id:            2,
		Nombre:        "producto 2",
		Color:         "verde",
		Precio:        10,
		Stock:         10,
		Codigo:        "1234",
		Publicado:     true,
		FechaCreacion: "2019-01-01",
	}

	// Database
	database := []domain.Product{
		{Id: 1, Nombre: "producto 1", Color: "rojo", Precio: 10, Stock: 10, Codigo: "123", Publicado: true, FechaCreacion: "2020-01-01"},
		{Id: 2, Nombre: "producto 2", Color: "verde", Precio: 10, Stock: 10, Codigo: "1234", Publicado: true, FechaCreacion: "2019-01-01"},
	}
	mockStorage := MockStorage{
		dataMock: database,
		errWrite: "",
		errRead:  "",
	}
	// Capas
	repository := NewRepository(&mockStorage)
	service := NewService(repository)

	// -------------------------------------------------------
	// Se ejecuta el test
	// -------------------------------------------------------

	result, err := service.GetById(getId)

	// -------------------------------------------------------
	// Se compara el resultado obtenido con el resultado esperado
	// -------------------------------------------------------

	assert.Nil(t, err)
	assert.Equal(t, resultadoEsperado, result)
	assert.True(t, mockStorage.ReadWasCalled)
}

func TestGetByIdFail(t *testing.T) {

	// -------------------------------------------------------
	// Se inicializan los datos a usar en el test (input y output)
	// -------------------------------------------------------

	// Get id
	getId := 3
	// Database
	database := []domain.Product{
		{Id: 1, Nombre: "producto 1", Color: "rojo", Precio: 10, Stock: 10, Codigo: "123", Publicado: true, FechaCreacion: "2020-01-01"},
		{Id: 2, Nombre: "producto 2", Color: "verde", Precio: 10, Stock: 10, Codigo: "1234", Publicado: true, FechaCreacion: "2019-01-01"},
	}
	mockStorage := MockStorage{
		dataMock: database,
		errWrite: "",
		errRead:  "",
	}
	// Error
	ErrProductNotFound := fmt.Errorf(ProductNotFound, getId)
	expectedError := fmt.Errorf("error obteniendo el producto con el id %d: %w", getId, ErrProductNotFound)
	// Capas
	repository := NewRepository(&mockStorage)
	service := NewService(repository)

	// -------------------------------------------------------
	// Se ejecuta el test
	// -------------------------------------------------------

	_, err := service.GetById(getId)

	// -------------------------------------------------------
	// Se compara el resultado obtenido con el resultado esperado
	// -------------------------------------------------------

	assert.Equal(t, expectedError, err)
	assert.True(t, mockStorage.ReadWasCalled)
}

func TestStore(t *testing.T) {

	// -------------------------------------------------------
	// Se inicializan los datos a usar en el test (input y output)
	// -------------------------------------------------------

	// Store data
	newProduct := domain.Product{
		Id:            1,
		Nombre:        "producto 1",
		Color:         "rojo",
		Precio:        10,
		Stock:         10,
		Codigo:        "123",
		Publicado:     true,
		FechaCreacion: "2020-01-01",
	}
	// Database
	mockStorage := MockStorage{
		dataMock: nil,
		errWrite: "",
		errRead:  "",
	}
	// Capas
	repository := NewRepository(&mockStorage)
	service := NewService(repository)

	// -------------------------------------------------------
	// Se ejecuta el test
	// -------------------------------------------------------

	resultProduct, err := service.Store(
		newProduct.Nombre,
		newProduct.Color,
		newProduct.Precio,
		newProduct.Stock,
		newProduct.Codigo,
		newProduct.Publicado,
		newProduct.FechaCreacion,
	)

	// -------------------------------------------------------
	// Se compara el resultado obtenido con el resultado esperado
	// -------------------------------------------------------

	assert.Nil(t, err)
	assert.Equal(t, newProduct, resultProduct)
	assert.Equal(t, newProduct.Id, resultProduct.Id)
	assert.Equal(t, 1, len(mockStorage.dataMock))
	assert.True(t, mockStorage.WriteWasCalled)
}

func TestGetAll(t *testing.T) {

	// -------------------------------------------------------
	// Se inicializan los datos a usar en el test (input y output)
	// -------------------------------------------------------

	// Database
	database := []domain.Product{
		{Id: 1, Nombre: "producto 1", Color: "rojo", Precio: 10, Stock: 10, Codigo: "123", Publicado: true, FechaCreacion: "2020-01-01"},
		{Id: 2, Nombre: "producto 2", Color: "verde", Precio: 10, Stock: 10, Codigo: "1234", Publicado: true, FechaCreacion: "2019-01-01"},
	}
	mockStorage := MockStorage{
		dataMock: database,
		errWrite: "",
		errRead:  "",
	}
	// Capas
	repository := NewRepository(&mockStorage)
	service := NewService(repository)

	// -------------------------------------------------------
	// Se ejecuta el test
	// -------------------------------------------------------

	result, err := service.GetAll()

	// -------------------------------------------------------
	// Se compara el resultado obtenido con el resultado esperado
	// -------------------------------------------------------

	assert.Nil(t, err)
	assert.Equal(t, database, result)
	assert.True(t, mockStorage.ReadWasCalled)
}

func TestUpdateNombre(t *testing.T) {

	// -------------------------------------------------------
	// Se inicializan los datos a usar en el test (input y output)
	// -------------------------------------------------------

	// Update data
	updateId := 2
	updateNombre := "producto modificado"
	// Database
	database := []domain.Product{
		{Id: 1, Nombre: "producto 1", Color: "rojo", Precio: 10, Stock: 10, Codigo: "123", Publicado: true, FechaCreacion: "2020-01-01"},
		{Id: 2, Nombre: "producto 2", Color: "verde", Precio: 10, Stock: 10, Codigo: "1234", Publicado: true, FechaCreacion: "2019-01-01"},
	}
	mockStorage := MockStorage{
		dataMock: database,
		errWrite: "",
		errRead:  "",
	}
	// Capas
	repository := NewRepository(&mockStorage)
	service := NewService(repository)

	// -------------------------------------------------------
	// Se ejecuta el test
	// -------------------------------------------------------

	result, err := service.UpdateNombre(updateId, updateNombre)

	// -------------------------------------------------------
	// Se compara el resultado obtenido con el resultado esperado
	// -------------------------------------------------------

	assert.Nil(t, err)
	assert.Equal(t, updateNombre, result.Nombre)
	assert.True(t, mockStorage.WriteWasCalled)
}

func TestUpdatePrecio(t *testing.T) {

	// -------------------------------------------------------
	// Se inicializan los datos a usar en el test (input y output)
	// -------------------------------------------------------

	// Update data
	updateId := 2
	updatePrecio := 2100.0
	// Database
	database := []domain.Product{
		{Id: 1, Nombre: "producto 1", Color: "rojo", Precio: 10, Stock: 10, Codigo: "123", Publicado: true, FechaCreacion: "2020-01-01"},
		{Id: 2, Nombre: "producto 2", Color: "verde", Precio: 10, Stock: 10, Codigo: "1234", Publicado: true, FechaCreacion: "2019-01-01"},
	}
	mockStorage := MockStorage{
		dataMock: database,
		errWrite: "",
		errRead:  "",
	}
	// Capas
	repository := NewRepository(&mockStorage)
	service := NewService(repository)

	// -------------------------------------------------------
	// Se ejecuta el test
	// -------------------------------------------------------

	result, err := service.UpdatePrecio(updateId, updatePrecio)

	// -------------------------------------------------------
	// Se compara el resultado obtenido con el resultado esperado
	// -------------------------------------------------------

	assert.Nil(t, err)
	assert.Equal(t, updatePrecio, result.Precio)
	assert.True(t, mockStorage.WriteWasCalled)
}
