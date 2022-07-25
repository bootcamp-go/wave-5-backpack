package transactions

import (
	"ejer02-TT/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

type StubStore struct {
	ReadCalled bool
}

func (fs *StubStore) Read(data interface{}) error {
	fs.ReadCalled = true
	a := data.(*[]domain.Transaction)
	*a = []domain.Transaction{
		{Id: 1, TranCode: "tranCode", Currency: "moneda", Amount: 12.5, Transmitter: "transmitter", Reciever: "reciever", TranDate: "tranDate"},
	}
	return nil
}

func (fs *StubStore) Write(data interface{}) error {
	return nil
}

func (fs *StubStore) Ping() error {
	return nil
}

func TestGetAll(t *testing.T) {
	stub := StubStore{}
	repo := NewRepository(&stub)
	expected := []domain.Transaction{
		{Id: 1, TranCode: "tranCode", Currency: "moneda", Amount: 12.5, Transmitter: "transmitter", Reciever: "reciever", TranDate: "tranDate"},
	}

	a, err := repo.GetAll()

	assert.Nil(t, err)
	assert.Equal(t, expected, a)
}

func TestUpdate(t *testing.T) {
	stub := StubStore{}
	repo := NewRepository(&stub)
	expected := domain.Transaction{Id: 1, TranCode: "cambio", Currency: "moneda", Amount: 10.5, Transmitter: "transmitter", Reciever: "reciever", TranDate: "tranDate"}

	a, err := repo.UpdateCodeAndAmount(1, "cambio", 10.5)

	assert.Nil(t, err)
	assert.Equal(t, expected, a)
	assert.True(t, stub.ReadCalled)
}
