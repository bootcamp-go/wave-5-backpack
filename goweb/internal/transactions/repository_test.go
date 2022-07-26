package transactions

import (
	"fmt"
	"testing"

	"github.com/bootcamp-go/wave-5-backpack/tree/lopez_cristian/goweb/internal/models"
	"github.com/stretchr/testify/assert"
)

type MockStorage struct {
	Data        []models.Transaction
	ReadCalled  bool
	WriteCalled bool
	ErrorRead   string
	ErrorWrite  string
}

func (s *MockStorage) Read(data interface{}) error {
	s.ReadCalled = true

	p := data.(*[]models.Transaction)
	*p = s.Data

	return nil
}

func (s *MockStorage) Write(data interface{}) error {
	s.WriteCalled = true

	p := data.([]models.Transaction)
	s.Data = p

	return nil
}

func TestUpdateStorage(t *testing.T) {
	//Arrange
	tr := models.Transaction{
		ID:       1,
		Monto:    1000.5,
		Cod:      "aaa111",
		Moneda:   "ARS",
		Emisor:   "Mercado Pago",
		Receptor: "BBVA",
		Fecha:    "2020-25-07",
	}

	expected := models.Transaction{
		ID:       1,
		Monto:    500,
		Cod:      "aaa112",
		Moneda:   "ARS",
		Emisor:   "BBVA",
		Receptor: "Mercado Pago",
		Fecha:    "2020-25-07",
	}

	//Act
	mockStorage := MockStorage{Data: []models.Transaction{tr}}
	repository := repository{&mockStorage}
	transaction, err := repository.Update(1, 500, "aaa112", "ARS", "BBVA", "Mercado Pago")

	//Assert
	assert.Equal(t, expected, transaction)
	assert.Nil(t, err)
}

func TestDeleteStorage(t *testing.T) {
	//Arrange
	data := []models.Transaction{
		{
			ID:       1,
			Monto:    1000.5,
			Cod:      "aaa111",
			Moneda:   "ARS",
			Emisor:   "Mercado Pago",
			Receptor: "BBVA",
			Fecha:    "2020-25-07",
		},
		{
			ID:       2,
			Monto:    500,
			Cod:      "aaa112",
			Moneda:   "ARS",
			Emisor:   "BBVA",
			Receptor: "Mercado Pago",
			Fecha:    "2020-25-07",
		},
	}

	//Act
	storage := MockStorage{Data: data}
	repo := NewRepository(&storage)
	id, err := repo.Delete(2)

	//Assert
	assert.True(t, storage.ReadCalled)
	assert.True(t, storage.WriteCalled)
	assert.Equal(t, 2, id)
	assert.Len(t, storage.Data, 1)
	assert.Nil(t, err)
}

func TestDeleteStorageNotFound(t *testing.T) {
	//Arrange
	storage := MockStorage{}
	errExpected := fmt.Errorf("error: ID %v no existe\n", 1)

	//Act
	repo := NewRepository(&storage)
	id, err := repo.Delete(1)

	//Assert
	if assert.Errorf(t, errExpected, "1") {
		assert.Equal(t, errExpected, err)
	}
	assert.Equal(t, 0, id)
	assert.Len(t, storage.Data, 0)
	assert.True(t, storage.ReadCalled)
	assert.False(t, storage.WriteCalled)
}

type StubStorage struct{}

func (s *StubStorage) Read(data interface{}) error {
	p := data.(*[]models.Transaction)
	*p = []models.Transaction{
		{
			ID:       1,
			Monto:    1000.5,
			Cod:      "aaa111",
			Moneda:   "ARS",
			Emisor:   "Mercado Pago",
			Receptor: "BBVA",
			Fecha:    "2020-25-07",
		},
		{
			ID:       2,
			Monto:    1000.5,
			Cod:      "aaa111",
			Moneda:   "ARS",
			Emisor:   "Mercado Pago",
			Receptor: "BBVA",
			Fecha:    "2020-25-07",
		},
	}

	return nil
}

func (s *StubStorage) Write(data interface{}) error {
	return nil
}

func TestRepositoryGetAll(t *testing.T) {
	//Arrange
	storage := &StubStorage{}
	repository := NewRepository(storage)
	transactionsExpected := []models.Transaction{
		{
			ID:       1,
			Monto:    1000.5,
			Cod:      "aaa111",
			Moneda:   "ARS",
			Emisor:   "Mercado Pago",
			Receptor: "BBVA",
			Fecha:    "2020-25-07",
		},
		{
			ID:       2,
			Monto:    1000.5,
			Cod:      "aaa111",
			Moneda:   "ARS",
			Emisor:   "Mercado Pago",
			Receptor: "BBVA",
			Fecha:    "2020-25-07",
		},
	}
	//Act
	transactions, err := repository.GetAll()

	//Assert
	assert.Equal(t, transactionsExpected, transactions)
	assert.Nil(t, err)
}
