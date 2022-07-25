package transactions

import (
	"testing"

	"github.com/bootcamp-go/wave-5-backpack/tree/lopez_cristian/goweb/internal/models"
	"github.com/stretchr/testify/assert"
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
			Monto:    1000.5,
			Cod:      "aaa111",
			Moneda:   "ARS",
			Emisor:   "Mercado Pago",
			Receptor: "BBVA",
			Fecha:    "2020-25-07",
		},
		{
			ID:       1,
			Monto:    1000.5,
			Cod:      "aaa111",
			Moneda:   "ARS",
			Emisor:   "Mercado Pago",
			Receptor: "BBVA",
			Fecha:    "2020-25-07",
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
	repoStub := &StubRepository{}

	service := NewService(repoStub)

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
			ID:       1,
			Monto:    1000.5,
			Cod:      "aaa111",
			Moneda:   "ARS",
			Emisor:   "Mercado Pago",
			Receptor: "BBVA",
			Fecha:    "2020-25-07",
		},
	}

	// Act
	transactions, err := service.GetAll()

	// Assert
	assert.Equal(t, transactionsExpected, transactions)
	assert.Nil(t, err)
}
