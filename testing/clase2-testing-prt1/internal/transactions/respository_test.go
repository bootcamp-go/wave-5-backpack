/*---------------------------------------------------------*

     Assignment:	Practica #1 - Testing
         Author:	Israel Fabela
	   Language:	go1.18.3 darwin/arm64
		  Topic:	Testing

	© Mercado Libre - IT Bootcamp 2022

----------------------------------------------------------*/
package transactions

import (
	"testing"

	"clase2-testing-prt1/internal/domain"

	"github.com/stretchr/testify/assert"
)

type StubBank struct{}

func (fs *StubBank) Read(data interface{}) error {

	//	Arrange
	hardCoding := data.(*[]domain.Transaction)
	*hardCoding = []domain.Transaction{
		{
			Id:                1,
			CodigoTransaccion: "abc123",
			Monto:             1100.5,
			Moneda:            "MXN",
			Emisor:            "HSBC",
			Receptor:          "Banorte",
			Fecha:             "2022/05/16",
		},
		{
			Id:                2,
			CodigoTransaccion: "cde456",
			Monto:             2500.5,
			Moneda:            "USD",
			Emisor:            "BBVA",
			Receptor:          "",
			Fecha:             "2020/07/08",
		},
	}

	return nil
}

func (fs *StubBank) Write(data interface{}) error {
	return nil
}

func (fs *StubBank) Ping() error {
	return nil
}

func (fs *StubBank) Update(id int, codeTra string, coin string, monto float64, emisor string,
	receptor string, fecha string) (domain.Transaction, error) {
	return domain.Transaction{}, nil
}

func (fs *StubBank) Patch(id int, codeTra string, monto float64) (domain.Transaction, error) {
	return domain.Transaction{}, nil
}

func (fs *StubBank) Delete(id int) error {
	return nil
}

func TestGetAll(t *testing.T) {

	//	Arrange
	stub := StubBank{}
	repo := NewRepository(&stub)
	expected := []domain.Transaction{
		{Id: 1,
			CodigoTransaccion: "abc123",
			Monto:             1100.5,
			Moneda:            "MXN",
			Emisor:            "HSBC",
			Receptor:          "Banorte",
			Fecha:             "2022/05/16",
		},
		{
			Id:                2,
			CodigoTransaccion: "cde456",
			Monto:             2500.5,
			Moneda:            "USD",
			Emisor:            "BBVA",
			Receptor:          "",
			Fecha:             "2020/07/08",
		},
	}

	//	Act
	a, err := repo.GetAll()

	//	Assert
	assert.Nil(t, err)
	assert.Equal(t, expected, a)
}
