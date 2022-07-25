package transactions

import (
	"arquitectura/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

type repositoryStub struct {
	GetAllWasCalled bool
	BeforeUpdate    domain.Transaction
}

func (r *repositoryStub) GetAll() ([]domain.Transaction, error) {
	r.GetAllWasCalled = true
	t1 := domain.Transaction{
		Id:          1,
		TranCode:    "SKU1234",
		Currency:    "CLP",
		Amount:      500000,
		Transmitter: "cmonsalve",
		Reciever:    "jperez",
		TranDate:    "10-10-2021",
	}
	t2 := domain.Transaction{
		Id:          2,
		TranCode:    "SKU5678",
		Currency:    "USD",
		Amount:      500,
		Transmitter: "jperez",
		Reciever:    "ctorres",
		TranDate:    "20-07-2022",
	}
	list := []domain.Transaction{t1, t2}
	return list, nil
}

func (r *repositoryStub) Store(id int, tranCode, currency string, amount float64, transmitter, receiver, tranDate string) (domain.Transaction, error) {
	return r.BeforeUpdate, nil
}

func (r *repositoryStub) Update(id int, tranCode, currency string, amount float64, transmitter, receiver, tranDate string) (domain.Transaction, error) {
	return r.BeforeUpdate, nil
}

func (r *repositoryStub) UpdateTranCode(id int, tranCode string) (domain.Transaction, error) {
	r.GetAll() // Se llama al método GetAll para validar la variable GetAllWasCalled
	r.BeforeUpdate.TranCode = tranCode
	return r.BeforeUpdate, nil
}

func (r *repositoryStub) UpdateAmount(id int, amount float64) (domain.Transaction, error) {
	return r.BeforeUpdate, nil
}

func (r *repositoryStub) Delete(id int) error {
	return nil
}

func (r *repositoryStub) LastID() (int, error) {
	return 0, nil
}

func TestGetAll(t *testing.T) {
	t1 := domain.Transaction{
		Id:          1,
		TranCode:    "SKU1234",
		Currency:    "CLP",
		Amount:      500000,
		Transmitter: "cmonsalve",
		Reciever:    "jperez",
		TranDate:    "10-10-2021",
	}
	t2 := domain.Transaction{
		Id:          2,
		TranCode:    "SKU5678",
		Currency:    "USD",
		Amount:      500,
		Transmitter: "jperez",
		Reciever:    "ctorres",
		TranDate:    "20-07-2022",
	}
	stubRepo := repositoryStub{}
	service := NewService(&stubRepo)
	expectedResult := []domain.Transaction{t1, t2}

	result, err := service.GetAll()

	assert.NoError(t, err)
	assert.Equal(t, expectedResult, result)
}

func TestUpdateTranCode(t *testing.T) {
	before := domain.Transaction{
		Id:          1,
		TranCode:    "BEFORE",
		Currency:    "CLP",
		Amount:      500000,
		Transmitter: "cmonsalve",
		Reciever:    "jperez",
		TranDate:    "10-10-2021",
	}
	stubRepo := repositoryStub{false, before}
	service := NewService(&stubRepo)
	expectedResult := domain.Transaction{
		Id:          1,
		TranCode:    "AFTER",
		Currency:    "CLP",
		Amount:      500000,
		Transmitter: "cmonsalve",
		Reciever:    "jperez",
		TranDate:    "10-10-2021",
	}

	result, err := service.UpdateTranCode(1, "AFTER")

	assert.NoError(t, err)
	assert.Equal(t, expectedResult, result)
	assert.True(t, stubRepo.GetAllWasCalled)
}
