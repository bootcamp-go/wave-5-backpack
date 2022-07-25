package transactions

import (
	"encoding/json"
	"goweb/internal/domain"

	"testing"

	"github.com/stretchr/testify/assert"
)

// Ejercicio 1

type StubStore struct{}

func (s *StubStore) Write(data interface{}) error {
	return nil
}

func (s *StubStore) Read(data interface{}) error {

	newData := []domain.Transaction{
		{Id: 1, TranCode: "ABC1234", Currency: "USD", Amount: 200.0, Transmitter: "MERCADOPAGO", Reciever: "jose", TranDate: "05-07-22"},
		{Id: 2, TranCode: "asdfghj", Currency: "clp", Amount: 1500, Transmitter: "meli", Reciever: "frp", TranDate: "12-02-22"},
	}

	file, err := json.Marshal(newData)
	if err != nil {
		return err
	}
	return json.Unmarshal(file, &data)
}

func (s *StubStore) Ping() error {
	return nil
}

func TestGetAll(t *testing.T) {

	myStubStore := StubStore{}
	repositorio := NewRepository(&myStubStore)

	listaEsperada := []domain.Transaction{
		{Id: 1, TranCode: "ABC1234", Currency: "USD", Amount: 200.0, Transmitter: "MERCADOPAGO", Reciever: "jose", TranDate: "05-07-22"},
		{Id: 2, TranCode: "asdfghj", Currency: "clp", Amount: 1500, Transmitter: "meli", Reciever: "frp", TranDate: "12-02-22"},
	}

	resultado, _ := repositorio.GetAll()

	assert.Equal(t, listaEsperada, resultado)

}

type MockStore struct {
	readteInvocado bool
}

// Ejercicio 2

func (s *MockStore) Write(data interface{}) error {
	return nil
}

func (s *MockStore) Read(data interface{}) error {

	transactionsBeforeUpdate := data.(*[]domain.Transaction)
	*transactionsBeforeUpdate = []domain.Transaction{
		{Id: 1, TranCode: "ABC1234", Currency: "USD", Amount: 200.0, Transmitter: "MERCADOPAGO", Reciever: "jose", TranDate: "05-07-22"},
		{Id: 2, TranCode: "asdfghj", Currency: "clp", Amount: 1500, Transmitter: "meli", Reciever: "frp", TranDate: "12-02-22"},
	}

	s.readteInvocado = true
	return nil
}

func (s *MockStore) Ping() error {
	return nil
}

func TestUpdate(t *testing.T) {

	myMockStore := MockStore{}
	repositorio := NewRepository(&myMockStore)

	transactionsAfterUpdate := []domain.Transaction{
		{
			Id:          1,
			TranCode:    "ABC1234",
			Currency:    "USD",
			Amount:      200.0,
			Transmitter: "MERCADOPAGO",
			Reciever:    "jose",
			TranDate:    "05-07-22"},
		{
			Id:          2,
			TranCode:    "XYZ1234",
			Currency:    "clp",
			Amount:      250.0,
			Transmitter: "meli",
			Reciever:    "frp",
			TranDate:    "12-02-22"},
	}

	resultado, err := repositorio.UpdateCodeAmount(2, "XYZ1234", 250.0)

	assert.Nil(t, err)
	assert.True(t, myMockStore.readteInvocado)
	assert.Equal(t, transactionsAfterUpdate[1], resultado)

}
