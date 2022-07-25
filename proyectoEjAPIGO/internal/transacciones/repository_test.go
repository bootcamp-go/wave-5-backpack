package transacciones

import (
	"goweb/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Stub Test
type StubRepository struct{}

func (s *StubRepository) Store(id int, codigo_transaccion, moneda, emisor, receptor, fecha_transaccion string, monto float64) (domain.Transaccion, error) {
	return domain.Transaccion{}, nil
}

func (s *StubRepository) Update(id int, codigo_transaccion, moneda, emisor, receptor, fecha_transaccion string, monto float64) (domain.Transaccion, error) {
	return domain.Transaccion{}, nil
}

func (s *StubRepository) UpdateCTandMonto(id int, codigo_transaccion string, monto float64) (domain.Transaccion, error) {
	return domain.Transaccion{}, nil
}

func (s *StubRepository) GetAll() ([]domain.Transaccion, error) {
	return []domain.Transaccion{
		{
			ID:                 1,
			Monto:              500.5,
			Codigo_transaccion: "aaa111",
			Moneda:             "ARS",
			Emisor:             "Mercado Pago",
			Receptor:           "BBVA",
			Fecha_transaccion:  "2022-07-25 12:00",
		},
		{
			ID:                 2,
			Monto:              1000,
			Codigo_transaccion: "aaa112",
			Moneda:             "ARS",
			Emisor:             "BBVA",
			Receptor:           "Mercado Pago",
			Fecha_transaccion:  "2022-07-25 12:00",
		},
	}, nil
}

func (s *StubRepository) GetByID(id int) (domain.Transaccion, error) {
	return domain.Transaccion{}, nil
}

func (s *StubRepository) LastID() (int, error) {
	return 0, nil
}

func (s *StubRepository) Delete(id int) error {
	return nil
}

func TestGetAll(t *testing.T) {
	// Arrange
	service := NewService(&StubRepository{})
	transactionsExpected := []domain.Transaccion{
		{
			ID:                 1,
			Monto:              500,
			Codigo_transaccion: "asd124",
			Moneda:             "USD",
			Emisor:             "Mercado Pago",
			Receptor:           "Santander",
			Fecha_transaccion:  "2022-07-25 12:30",
		},
		{
			ID:                 2,
			Monto:              1500,
			Codigo_transaccion: "asd123",
			Moneda:             "USD",
			Emisor:             "BBVA",
			Receptor:           "Mercado Pago",
			Fecha_transaccion:  "2022-07-25 12:35",
		},
	}

	// Act
	transactions, err := service.GetAll()

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, transactionsExpected, transactions)
}
