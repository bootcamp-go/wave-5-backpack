package transactions

import (
	"errors"
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

	delIdExist := 1
	delIdNoExist := 4

	expectedError := errors.New("id not found")
	//act delete correct

	_, err := serv.Delete(delIdExist)
	tr, _ := serv.GetAll()
	for _, v := range tr {
		if v.Id == delIdExist {
			err = errors.New("id still exist")
		}
	}

	//assert delete correct
	fmt.Println(tr)
	assert.True(t, mockDB.MockReadUsed)
	assert.Nil(t, err, "have an error: %w", err)

	//act delete wrong

	_, err = serv.Delete(delIdNoExist)
	//assert delete correct
	fmt.Println(tr)
	assert.EqualError(t, expectedError, err.Error(), "have an error %w", err)
}
