package transaction

import (
	"proyecto-web/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

type StubRepo struct {
}

func TestReadAll(t *testing.T) {
	// arrange
	stub := StubRepo{}
	service := NewService(&stub)

	// act
	result := service.GetAll()

	t1 := result[0]
	t2 := result[1]

	// assert
	assert.Equal(t, len(result), 2)
	assert.Equal(t, t1.CodigoTransaccion, "A1")
	assert.Equal(t, t2.CodigoTransaccion, "A2")
}

func (s *StubRepo) GetAll() []domain.Transaction {
	return []domain.Transaction{
		{
			Id:                0,
			CodigoTransaccion: "A1",
			Moneda:            "PESOS",
			Monto:             5.0,
			Emisor:            "ARCOR",
			Receptor:          "AFIP",
			FechaTransaccion:  "12-01-2022",
		},
		{
			Id:                1,
			CodigoTransaccion: "A2",
			Moneda:            "DOLARES",
			Monto:             20.0,
			Emisor:            "TOYOTA",
			Receptor:          "AFIP",
			FechaTransaccion:  "22-06-2022",
		},
	}
}

func (s *StubRepo) Create(id int, codigoTransaccion string, moneda string, monto float64, emisor, receptor, fecha string) domain.Transaction {
	return domain.Transaction{}
}

func (s *StubRepo) GetById(id int) (domain.Transaction, error) {
	return domain.Transaction{}, nil
}

func (s *StubRepo) Update(id int, codigoTransaccion string, moneda string, monto float64, emisor string, receptor string, fecha string) (domain.Transaction, error) {
	return domain.Transaction{}, nil
}

func (s *StubRepo) UpdateParcial(id int, codigoTransaccion string, monto float64) (domain.Transaction, error) {
	return domain.Transaction{}, nil
}

func (s *StubRepo) Delete(id int) error {
	return nil
}
