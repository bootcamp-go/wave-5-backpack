package transactions

import (
	"goweb/internals/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

type StubRepository struct{}

func (sr *StubRepository) GetAll() ([]domain.Transaction, error) {
	return []domain.Transaction{
		{
			Id:       1,
			Codigo:   "SAKMDLN89392",
			Moneda:   "ARG",
			Monto:    1000,
			Emisor:   "Cristian",
			Receptor: "Julian",
			Fecha:    "25-07-2022",
		},
		{
			Id:       2,
			Codigo:   "SASADLN89392",
			Moneda:   "MEX",
			Monto:    2000,
			Emisor:   "Julián",
			Receptor: "Cristian",
			Fecha:    "25-07-2022",
		},
	}, nil
}

func (sr *StubRepository) Store(id int, codigo string, moneda string, monto int, emisor string, receptor string) (domain.Transaction, error) {
	return domain.Transaction{}, nil
}
func (sr *StubRepository) LastID() (int, error) {
	return 0, nil
}
func (sr *StubRepository) Update(id int, codigo string, moneda string, monto int, emisor string, receptor string) (domain.Transaction, error) {
	return domain.Transaction{}, nil
}
func (sr *StubRepository) Delete(id int) error {
	return nil
}

func TestGetAll(t *testing.T) {
	//Arrange
	stub := StubRepository{}
	service := NewService(&stub)
	TransactionExpected := []domain.Transaction{
		{
			Id:       1,
			Codigo:   "SAKMDLN89392",
			Moneda:   "ARG",
			Monto:    1000,
			Emisor:   "Cristian",
			Receptor: "Julian",
			Fecha:    "25-07-2022",
		},
		{
			Id:       2,
			Codigo:   "SASADLN89392",
			Moneda:   "MEX",
			Monto:    2000,
			Emisor:   "Julián",
			Receptor: "Cristian",
			Fecha:    "25-07-2022",
		},
	}
	//Act
	transaction, err := service.GetAll()
	//Assert
	assert.Equal(t, TransactionExpected, transaction)
	assert.Nil(t, err)
}
