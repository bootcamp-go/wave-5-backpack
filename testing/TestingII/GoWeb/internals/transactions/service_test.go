package transactions

import (
	"GoWeb/internals/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

type StubDB struct{}

func (s *StubDB) GetAll() ([]domain.Transanction, error) {
	return []domain.Transanction{{
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
		}}, nil
}

func (s StubDB) lastID() (int, error) {
	return 0, nil
}

func (s StubDB) Store(id int, code, coin string, amount float64, emisor, receptor, date string) (domain.Transanction, error) {
	return domain.Transanction{}, nil
}

func (s StubDB) Update(id int, code, coin string, amount float64, emisor, receptor, date string) (domain.Transanction, error) {
	return domain.Transanction{}, nil
}

func (s StubDB) Delete(id int) error {
	return nil
}

func (s StubDB) UpdateCode(id int, code string, amount float64) (domain.Transanction, error) {
	return domain.Transanction{}, nil
}

func TestGetAll(t *testing.T) {
	//arrange
	myStubDB := &StubDB{}
	Get := NewService(myStubDB)
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
	tran, _ := Get.GetAll()

	assert.Equal(t, esperado, tran)

}
