package transactions

import (
	"goweb/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

type StubStore struct {
	data []domain.Transaction
}

func (fs *StubStore) Read(data interface{}) error {
	a := data.(*[]domain.Transaction)
	*a = fs.data
	return nil
}

func (fs *StubStore) Write(data interface{}) error {
	a := data.([]domain.Transaction)
	fs.data = a
	return nil
}

func TestGetAll(t *testing.T) {
	expected := []domain.Transaction{
		{
			Id:        1,
			Code:      "aaf",
			Currency:  "COP",
			Amount:    300,
			Issuer:    "kevin",
			Recipient: "Daniel",
			Date:      "20 june",
		},
		{
			Id:        2,
			Code:      "b1f",
			Currency:  "COP",
			Amount:    200,
			Issuer:    "Jose",
			Recipient: "Prieto",
			Date:      "21 june",
		},
	}

	stub := StubStore{data: expected}
	repo := NewRepository(&stub)

	tr, err := repo.GetAll()

	assert.Nil(t, err)
	assert.Equal(t, expected, tr)
}

func TestUpdtate2(t *testing.T) {
	data := []domain.Transaction{
		{
			Id:        1,
			Code:      "aaf",
			Currency:  "COP",
			Amount:    300,
			Issuer:    "kevin",
			Recipient: "Daniel",
			Date:      "20 june",
		},
	}
	expected1 := domain.Transaction{
		Id:        1,
		Code:      "bbb",
		Currency:  "COP",
		Amount:    1,
		Issuer:    "kevin",
		Recipient: "Daniel",
		Date:      "20 june",
	}
	expected2 := []domain.Transaction{
		{
			Id:        1,
			Code:      "bbb",
			Currency:  "COP",
			Amount:    1,
			Issuer:    "kevin",
			Recipient: "Daniel",
			Date:      "20 june",
		},
	}

	stub := StubStore{data: data}
	repo := NewRepository(&stub)

	tr, err := repo.Update2(1, "bbb", 1.0)
	assert.Nil(t, err)
	assert.Equal(t, expected1, tr)

	trs, err := repo.GetAll()
	assert.Nil(t, err)
	assert.Equal(t, trs, expected2)
}
