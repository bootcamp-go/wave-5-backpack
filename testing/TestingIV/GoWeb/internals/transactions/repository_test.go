package transactions

import (
	"GoWeb/internals/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

type StubStore struct{}

func (st *StubStore) Ping() error {
	return nil
}

func (st *StubStore) Write(data interface{}) error {
	return nil
}

func (st *StubStore) Read(data interface{}) error {
	p := data.(*[]domain.Transanction)
	*p = []domain.Transanction{
		{
			Id:       1,
			Code:     "QWE123",
			Coin:     "COP",
			Amount:   10000,
			Emisor:   "Juan David",
			Receptor: "MeLi",
			Date:     "06-06-2022",
		},
		{
			Id:       2,
			Code:     "ASD123",
			Coin:     "USD",
			Amount:   3000,
			Emisor:   "Sergio",
			Receptor: "MePa",
			Date:     "07-06-2022",
		},
	}
	return nil
}

func TestGetAllRepo(t *testing.T) {
	//arrange
	repo := StubStore{}
	get := NewRepository(&repo)
	esperado := []domain.Transanction{{
		Id:       1,
		Code:     "QWE123",
		Coin:     "COP",
		Amount:   10000,
		Emisor:   "Juan David",
		Receptor: "MeLi",
		Date:     "06-06-2022",
	},
		{
			Id:       2,
			Code:     "ASD123",
			Coin:     "USD",
			Amount:   3000,
			Emisor:   "Sergio",
			Receptor: "MePa",
			Date:     "07-06-2022",
		},
	}
	// act
	tran, err := get.GetAll()

	//assert
	assert.Equal(t, esperado, tran)
	assert.Nil(t, err)

}
