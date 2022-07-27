package transacciones

import (
	"fmt"
	"goweb/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Mock Test
type MockRepository struct {
	DataMock []domain.Transaccion
	errWrite string
	errRead  string
	ReadCall bool
}

func (m *MockRepository) Store(id int, codigo_transaccion, moneda, emisor, receptor, fecha_transaccion string, monto float64) (domain.Transaccion, error) {
	return domain.Transaccion{}, nil
}

func (m *MockRepository) Update(id int, codigo_transaccion, moneda, emisor, receptor, fecha_transaccion string, monto float64) (domain.Transaccion, error) {
	m.GetByID(1)
	return domain.Transaccion{
		ID:                 id,
		Monto:              monto,
		Codigo_transaccion: codigo_transaccion,
		Moneda:             moneda,
		Emisor:             emisor,
		Receptor:           receptor,
		Fecha_transaccion:  fecha_transaccion}, nil

}

func (m *MockRepository) UpdateCTandMonto(id int, codigo_transaccion string, monto float64) (domain.Transaccion, error) {
	return domain.Transaccion{}, nil
}

func (m *MockRepository) GetAll() ([]domain.Transaccion, error) {
	return []domain.Transaccion{}, nil
}

func (m *MockRepository) GetByID(id int) (domain.Transaccion, error) {
	m.ReadCall = true
	return domain.Transaccion{}, nil
}

func (m *MockRepository) LastID() (int, error) {
	return 0, nil
}

func (m *MockRepository) Delete(id int) error {
	return nil
}
func (ms *MockRepository) Write(data interface{}) error {
	if ms.errWrite != "" {
		return fmt.Errorf(ms.errWrite)
	}
	a := data.([]domain.Transaccion)
	ms.DataMock = append(ms.DataMock, a...)
	return nil
}

func (ms *MockRepository) Read(data interface{}) error {
	if ms.errRead != "" {
		return fmt.Errorf(ms.errRead)
	}
	user := data.(*[]domain.Transaccion)
	*user = ms.DataMock
	ms.ReadCall = true
	return nil
}

func TestUpdate(t *testing.T) {
	// Arrange
	mock := MockRepository{}
	service := NewService(&mock)
	beforeUpdate := domain.Transaccion{
		ID:                 1,
		Monto:              500,
		Codigo_transaccion: "asd124",
		Moneda:             "USD",
		Emisor:             "Mercado Pago",
		Receptor:           "Santander",
		Fecha_transaccion:  "2022-07-25 12:30",
	}
	transactionExpected := domain.Transaccion{

		ID:                 1,
		Monto:              100,
		Codigo_transaccion: "asd124",
		Moneda:             "USD",
		Emisor:             "Mercado Pago",
		Receptor:           "Santander",
		Fecha_transaccion:  "2022-07-25 12:30",
	}

	// Act

	transaction, err := service.Update(beforeUpdate.ID, "asd124", "USD", "Mercado Pago", "Santander", "2022-07-25 12:30", 100)

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, transactionExpected, transaction)
	assert.True(t, mock.ReadCall)
}
