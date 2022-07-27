/*
	----------------------------------------------------------------
	------------- Ejercicio 1 - Test Unitario GetAll() -------------
	----------------------------------------------------------------

	Generar un Stub del Store cuya función “Read” retorne dos productos con las especificaciones que deseen. Comprobar que GetAll() retorne la información exactamente igual a la esperada. Para esto:
	1. Dentro de la carpeta /internal/(producto/usuario/transacción), crear un archivo service_test.go con el test diseñado.
*/

package transactions

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"goweb/internal/domain"
)

type StubService struct {}

func(ss StubService) GetAll() ([]domain.Transaction, error) {
	return []domain.Transaction{
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
	}, nil
}
func (ss StubService) Delete(_ int) error{
	return nil
}
func (ss StubService) Store(_ int, _, _ string, _ int, _, _, _ string) (domain.Transaction, error) {
	return domain.Transaction{}, nil
}
func(ss StubService) Update(_ int, _, _ string, _ int, _, _, _ string) (domain.Transaction, error) {
	return domain.Transaction{}, nil
}
func(ss StubService) UpdateAmount(_, _ int) (domain.Transaction, error) {
	return domain.Transaction{}, nil
}
func (ss StubService) LastID() (int, error){
	return 0, nil
}


func TestServiceGetAll(t *testing.T) {
	//Arrange
	myStubRepository := StubService{}
	motor := NewService(&myStubRepository)
	resultExpected := [2]domain.Transaction{
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