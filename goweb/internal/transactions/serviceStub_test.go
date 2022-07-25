package transactions

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/bootcamp-go/wave-5-backpack/tree/lopez_cristian/goweb/internal/models"
)

// Stub Test
type StubRepository struct{}

func (s *StubRepository) Store(monto float64, cod, moneda, emisor, receptor string) (models.Transaction, error) {
	return models.Transaction{}, nil
}

func (s *StubRepository) Update(id int, monto float64, cod, moneda, emisor, receptor string) (models.Transaction, error) {
	return models.Transaction{}, nil
}

func (s *StubRepository) Patch(id int, monto float64, cod, moneda, emisor, receptor string) (models.Transaction, error) {
	return models.Transaction{}, nil
}

func (s *StubRepository) GetAll() ([]models.Transaction, error) {
	return []models.Transaction{
		{
			ID:       1,
			Monto:    500.5,
			Cod:      "aaa111",
			Moneda:   "ARS",
			Emisor:   "Mercado Pago",
			Receptor: "BBVA",
			Fecha:    "2022-07-25 12:00",
		},
		{
			ID:       2,
			Monto:    1000,
			Cod:      "aaa112",
			Moneda:   "ARS",
			Emisor:   "BBVA",
			Receptor: "Mercado Pago",
			Fecha:    "2022-07-25 12:00",
		},
	}, nil
}

func (s *StubRepository) GetByID(id int) (models.Transaction, error) {
	return models.Transaction{}, nil
}

func (s *StubRepository) GetLastID() (int, error) {
	return 0, nil
}

func (s *StubRepository) Delete(id int) (int, error) {
	return 0, nil
}

func TestGetAll(t *testing.T) {
	// Arrange
	service := NewService(&StubRepository{})
	transactionsExpected := []models.Transaction{
		{
			ID:       1,
			Monto:    500.5,
			Cod:      "aaa111",
			Moneda:   "ARS",
			Emisor:   "Mercado Pago",
			Receptor: "BBVA",
			Fecha:    "2022-07-25 12:00",
		},
		{
			ID:       2,
			Monto:    1000,
			Cod:      "aaa112",
			Moneda:   "ARS",
			Emisor:   "BBVA",
			Receptor: "Mercado Pago",
			Fecha:    "2022-07-25 12:00",
		},
	}

	// Act
	transactions, err := service.GetAll()

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, transactionsExpected, transactions)
}
