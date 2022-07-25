package transactions

import (
	"testing"

	"github.com/bootcamp-go/wave-5-backpack/tree/lopez_cristian/goweb/internal/models"
	"github.com/stretchr/testify/assert"
)

type MockRepository struct {
	ReadCalled bool
}

func (m *MockRepository) Store(monto float64, cod, moneda, emisor, receptor string) (models.Transaction, error) {
	return models.Transaction{}, nil
}

func (m *MockRepository) Update(id int, monto float64, cod, moneda, emisor, receptor string) (models.Transaction, error) {
	return models.Transaction{
		ID:       id,
		Monto:    monto,
		Cod:      cod,
		Moneda:   moneda,
		Emisor:   emisor,
		Receptor: receptor,
		Fecha:    "2022-07-25 12:00"}, nil
}

func (m *MockRepository) Patch(id int, monto float64, cod, moneda, emisor, receptor string) (models.Transaction, error) {
	return models.Transaction{}, nil
}

func (m *MockRepository) GetAll() ([]models.Transaction, error) {
	return []models.Transaction{}, nil
}

func (m *MockRepository) GetByID(id int) (models.Transaction, error) {
	m.ReadCalled = true
	return models.Transaction{}, nil
}

func (m *MockRepository) GetLastID() (int, error) {
	return 0, nil
}

func (m *MockRepository) Delete(id int) (int, error) {
	return 0, nil
}

func TestUpdate(t *testing.T) {
	//Arrange
	mock := MockRepository{}
	service := NewService(&mock)
	beforeUpdate := models.Transaction{
		ID:       1,
		Monto:    500.5,
		Cod:      "aaa111",
		Moneda:   "ARS",
		Emisor:   "Mercado Pago",
		Receptor: "BBVA",
		Fecha:    "2022-07-25 12:00",
	}
	transactionExpected := models.Transaction{
		ID:       1,
		Monto:    1000,
		Cod:      "aaa112",
		Moneda:   "USD",
		Emisor:   "BBVA",
		Receptor: "Mercado Pago",
		Fecha:    "2022-07-25 12:00",
	}

	// Act
	transaction, err := service.Update(beforeUpdate.ID, 1000, "aaa112", "USD", "BBVA", "Mercado Pago")

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, transactionExpected, transaction)
	assert.True(t, mock.ReadCalled)
}
