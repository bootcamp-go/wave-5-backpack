package transactions

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdateService(t *testing.T) {
	//arrange
	mockDB := mockStore{}
	repo := NewRepository(&mockDB)
	serv := NewService(repo)
	afterUpdate := Transaction{
		Id:                1,
		CodigoTransaccion: 100,
		Moneda:            "USD",
		Monto:             33.3,
		Emisor:            "Buyer",
		Receptor:          "MePago",
		FechaTransaccion:  "27-07-2022"}
	//act
	tr, err := serv.Update(
		afterUpdate.Id,
		afterUpdate.CodigoTransaccion,
		afterUpdate.Moneda,
		afterUpdate.Monto,
		afterUpdate.Emisor,
		afterUpdate.Receptor,
		afterUpdate.FechaTransaccion)

	//assert
	assert.True(t, mockDB.MockReadUsed)
	assert.Nil(t, err, "have an error %w", err)
	assert.Equal(t, afterUpdate, tr)
}
func TestDeleteService(t *testing.T) {
	//arrange
	mockDB := mockStore{}
	repo := NewRepository(&mockDB)
	serv := NewService(repo)
	//act
	_, err := serv.Delete(1)
	tr, _ := serv.GetAll()

	//assert
	fmt.Println(tr)
	assert.True(t, mockDB.MockReadUsed)
	assert.Nil(t, err, "have an error %w", err)
}
