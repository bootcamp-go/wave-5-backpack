package transactions

import (
	"ejer02-TT/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockGetAll struct{}

func (m *MockGetAll) GetAll() ([]domain.Transaction, error) {

	var tl []domain.Transaction

	t := domain.Transaction{TranCode: "tranCode", Currency: "moneda", Amount: 12.5, Transmitter: "transmitter", Reciever: "reciever", TranDate: "tranDate"}

	tl = append(tl, t)

	return tl, nil
}
func (m *MockGetAll) Store(id int, tranCode, currency string, amount float64, transmitter, receiver, tranDate string) (domain.Transaction, error) {

	t := domain.Transaction{
		Id:          id,
		TranCode:    tranCode,
		Currency:    currency,
		Amount:      amount,
		Transmitter: transmitter,
		Reciever:    receiver,
		TranDate:    tranCode,
	}

	return t, nil

}

func (m *MockGetAll) Update(id int, tranCode, currency string, amount float64, transmitter, receiver, tranDate string) (domain.Transaction, error) {
	var tl []domain.Transaction

	t := domain.Transaction{
		Id:          id,
		TranCode:    tranCode,
		Currency:    currency,
		Amount:      amount,
		Transmitter: transmitter,
		Reciever:    receiver,
		TranDate:    tranCode,
	}

	for i := range tl {
		if tl[i].Id == id {
			t.Id = id
			tl[i] = t

		}
	}

	return t, nil
}

func (m *MockGetAll) UpdateCodeAndAmount(id int, tranCode string, amount float64) (domain.Transaction, error) {
	var tl []domain.Transaction

	t := domain.Transaction{
		Id:       id,
		TranCode: tranCode,
		Amount:   amount,
		TranDate: tranCode,
	}

	for i := range tl {
		if tl[i].Id == id {
			t.Id = id
			tl[i] = t

		}
	}
	return t, nil
}

func (m *MockGetAll) Delete(id int) error {

	return nil
}

func (m *MockGetAll) LastID() (int, error) {
	var ps []domain.Transaction

	if len(ps) == 0 {
		return 0, nil
	}

	return ps[len(ps)-1].Id, nil
}

func TestGetAllService(t *testing.T) {

	myMockDB := MockGetAll{}
	motor := NewService(&myMockDB)

	resultado, _ := motor.GetAll()

	var expected []domain.Transaction

	resultadoEsperado := domain.Transaction{TranCode: "tranCode", Currency: "moneda", Amount: 12.5, Transmitter: "transmitter", Reciever: "reciever", TranDate: "tranDate"}

	expected = append(expected, resultadoEsperado)

	assert.Equal(t, expected, resultado)

}
