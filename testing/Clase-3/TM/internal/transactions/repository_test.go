package transactions

import (
	"fmt"
	"goweb/internal/domain"

	"testing"

	"github.com/stretchr/testify/assert"
)

type MockStore struct {
	dataMock   []domain.Transaction
	errWrite   string
	errRead    string
	readCalled bool
}

func (m *MockStore) Read(data interface{}) error {
	if m.errRead != "" {
		return fmt.Errorf(m.errRead)
	}
	m.readCalled = true
	a := data.(*[]domain.Transaction)
	*a = m.dataMock

	fmt.Println(a)
	return nil
}

func (m *MockStore) Write(data interface{}) error {
	if m.errWrite != "" {
		return fmt.Errorf(m.errWrite)
	}
	a := data.([]domain.Transaction)
	m.dataMock = append(m.dataMock, a...)

	return nil

}

func (s *MockStore) Ping() error {
	return nil
}

// Ejercicio 1

func TestUpdate(t *testing.T) {
	// arrange

	database := []domain.Transaction{
		{
			Id:          1,
			TranCode:    "ABC1234",
			Currency:    "USD",
			Amount:      200.0,
			Transmitter: "MERCADOPAGO",
			Reciever:    "JOSE",
			TranDate:    "05-07-22",
		},
		{
			Id:          2,
			TranCode:    "XYZ1234",
			Currency:    "CLP",
			Amount:      1500000.0,
			Transmitter: "MELI",
			Reciever:    "JUAN",
			TranDate:    "12-02-22",
		},
		{
			Id:          3,
			TranCode:    "BBB456",
			Currency:    "USD",
			Amount:      500.00,
			Transmitter: "BANCO",
			Reciever:    "JUAN",
			TranDate:    "26-07-22",
		},
	}

	mockStorage := MockStore{
		dataMock: database,
	}

	// act
	repo := NewRepository(&mockStorage)
	result, err := repo.Update(3, "BBB456", "USD", 500.00, "BANCO", "JUAN", "26-07-22")

	// assert
	assert.Nil(t, err)
	assert.True(t, mockStorage.readCalled)
	assert.Equal(t, mockStorage.dataMock[2], result)

}

//Ejercicio 2

func TestDelete(t *testing.T) {
	// arrange

	database := []domain.Transaction{
		{
			Id:          1,
			TranCode:    "ABC1234",
			Currency:    "USD",
			Amount:      200.0,
			Transmitter: "MERCADOPAGO",
			Reciever:    "JOSE",
			TranDate:    "05-07-22",
		},
		{
			Id:          2,
			TranCode:    "XYZ1234",
			Currency:    "CLP",
			Amount:      1500000.0,
			Transmitter: "MELI",
			Reciever:    "JUAN",
			TranDate:    "12-02-22",
		},
		{
			Id:          3,
			TranCode:    "BBB456",
			Currency:    "USD",
			Amount:      500.00,
			Transmitter: "BANCO",
			Reciever:    "JUAN",
			TranDate:    "26-07-22",
		},
	}

	mockStorage := MockStore{
		dataMock: database,
	}

	// act
	repo := NewRepository(&mockStorage)
	err := repo.Delete(3)

	// assert
	assert.Nil(t, err)

}
