package transactions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockStore struct {
	MockReadUsed bool
}

// global variable emulate the information stored
var Tr = []Transaction{
	{
		Id:                1,
		CodigoTransaccion: 100,
		Moneda:            "COP",
		Monto:             34.3,
		Emisor:            "Buyer",
		Receptor:          "MeLi",
		FechaTransaccion:  "26-07-2022",
	},
}

func (fs *mockStore) Write(data interface{}) error {
	Tr = data.([]Transaction) // data of type interface should have and assertion type to be compatible with Tr type
	return nil
}

func (fs *mockStore) Read(data interface{}) error {
	a := data.(*[]Transaction) // type assertion , like pointer of data(type []Transaction) i.(Type)
	*a = Tr
	fs.MockReadUsed = true
	return nil
}

func (fs *mockStore) Ping() error {
	return nil
}

func TestUpdate(t *testing.T) {
	//arrange
	mockDB := mockStore{}
	repo := NewRepository(&mockDB)
	afterUpdate := Transaction{
		Id:                1,
		CodigoTransaccion: 100,
		Moneda:            "USD",
		Monto:             33.3,
		Emisor:            "Buyer",
		Receptor:          "MePago",
		FechaTransaccion:  "27-07-2022"}
	//act
	tr, err := repo.Update(
		afterUpdate.Id,
		afterUpdate.CodigoTransaccion,
		afterUpdate.Moneda,
		afterUpdate.Monto,
		afterUpdate.Emisor,
		afterUpdate.Receptor,
		afterUpdate.FechaTransaccion)

	//assert
	assert.True(t, mockDB.MockReadUsed)
	assert.Nil(t, err)
	assert.Equal(t, afterUpdate, tr)
}
