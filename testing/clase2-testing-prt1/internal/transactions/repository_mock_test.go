/*---------------------------------------------------------*

     Assignment:	Practica #1 - Testing
         Author:	Israel Fabela
	   Language:	go1.18.3 darwin/arm64
		  Topic:	Testing

	© Mercado Libre - IT Bootcamp 2022

----------------------------------------------------------*/
package transactions

import (
	"clase2-testing-prt1/internal/domain"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockBank struct {
	db             interface{}
	ReadBankCalled bool
}

func (m *MockBank) Read(data interface{}) error {
	r := reflect.ValueOf(data)
	r = reflect.Indirect(r)
	r.Set(reflect.ValueOf(m.db))
	m.ReadBankCalled = true
	return nil
}

func (m *MockBank) Write(data interface{}) error {
	m.db = data
	return nil
}

func (m *MockBank) Ping() error {
	return nil
}

func TestUpdate(t *testing.T) {

	//	Arrange
	testBeforeUpdate := []domain.Transaction{
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
			Receptor:          "Santander",
			Fecha:             "2020/07/08",
		},
	}

	testAfterUpdate := []domain.Transaction{
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
			Monto:             81274.75,
			Moneda:            "USD",
			Emisor:            "BBVA",
			Receptor:          "Santander",
			Fecha:             "2020/07/08",
		},
	}

	myMockBank := &MockBank{
		db:             testBeforeUpdate,
		ReadBankCalled: false,
	}
	repository := NewRepository(myMockBank)

	//	Act
	resultado, err := repository.UpdateOne(2, "cde456", 81274.75)

	//	Assert
	assert.Nil(t, err)
	assert.Equal(t, testAfterUpdate, myMockBank.db)
	assert.Equal(t, testAfterUpdate[1], resultado, "Deben coincidir")
	assert.True(t, myMockBank.ReadBankCalled)
}
