package transactions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// stub test
type StubStore struct {
}

var tr = []Transaction{
	{
		Id:                1,
		CodigoTransaccion: 100,
		Moneda:            "COP",
		Monto:             34.3,
		Emisor:            "Buyer",
		Receptor:          "MeLi",
		FechaTransaccion:  "26-07-2022",
	},
	{
		Id:                2,
		CodigoTransaccion: 100,
		Moneda:            "USD",
		Monto:             34.3,
		Emisor:            "Buyer",
		Receptor:          "MePa",
		FechaTransaccion:  "26-07-2022",
	}}

func (fs *StubStore) Write(data interface{}) error {
	return nil
}

func (fs *StubStore) Read(data interface{}) error {
	a := data.(*[]Transaction) // type assertion , like pointer of data(type []Transaction) i.(Type)
	*a = tr

	return nil
}

func (fs *StubStore) Ping() error {
	return nil
}

//
func TestRead(t *testing.T) {
	//arrange : put the instantiated methods for convert the tested function in an operational one
	stubDB := StubStore{}
	repo := NewRepository(&stubDB)
	expexted := []Transaction{
		{
			Id:                1,
			CodigoTransaccion: 100,
			Moneda:            "COP",
			Monto:             34.3,
			Emisor:            "Buyer",
			Receptor:          "MeLi",
			FechaTransaccion:  "26-07-2022",
		},
		{
			Id:                2,
			CodigoTransaccion: 100,
			Moneda:            "USD",
			Monto:             34.3,
			Emisor:            "Buyer",
			Receptor:          "MePa",
			FechaTransaccion:  "26-07-2022",
		}}

	//act : call the function or method to test his functionality
	a, err := repo.GetAll()

	//asser : evaluate performance of function at issue
	assert.Nil(t, err)
	assert.Equal(t, expexted, a)
}
