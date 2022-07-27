/*
	----------------------------------------------------------------
	------------- Ejercicio 1 - Test Unitario GetAll() -------------
	----------------------------------------------------------------

	Generar un Stub del Store cuya función “Read” retorne dos productos con las especificaciones que deseen. Comprobar que GetAll() retorne la información exactamente igual a la esperada. Para esto:
	1. Dentro de la carpeta /internal/(producto/usuario/transacción), crear un archivo repository_test.go con el test diseñado.
 */

package transactions

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"goweb/internal/domain"
)

type StrubStore struct {}

func (ss StrubStore) Read(data interface{}) error {
	transactions := data.(*[]domain.Transaction)
	*transactions = []domain.Transaction{
		{
			Id:             1016,
			CodTransaction: "string",
			Currency:       "string",
			Amount:         1111,
			Sender:         "string",
			Receiver:       "string",
			DateOrder:      "string",
		},
		{
			Id:             1020,
			CodTransaction: "string",
			Currency:       "string",
			Amount:         1111,
			Sender:         "string",
			Receiver:       "string",
			DateOrder:      "string",
		},
	}
	return nil
}

func (ss StrubStore) Write(_ interface{}) error {
	return nil
}
func (ss StrubStore) Ping() error {
	return nil
}


func TestRepositoryGetAll(t *testing.T){
	//Arrange
	myStubStore := StrubStore{}
	motor := NewRepository(myStubStore)
	resultExpected := []domain.Transaction{
		{
			Id: 1016,
			CodTransaction: "string",
			Currency: "string",
			Amount: 1111,
			Sender: "string",
			Receiver: "string",
			DateOrder: "string",
		},
		{
			Id: 1020,
			CodTransaction: "string",
			Currency: "string",
			Amount: 1111,
			Sender: "string",
			Receiver: "string",
			DateOrder: "string",
		},
	}

	//Act
	result, err := motor.GetAll()

	//Assert
	assert.Nil(t, err)
	assert.Equal(t, resultExpected, result, "information does not match")
}