package transactions

import (
	"GoWeb/internals/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockStore struct {
	ChangeEmisor bool
}

func (m *MockStore) Ping() error {
	return nil
}

func (m *MockStore) Write(data interface{}) error {
	return nil
}

func (m *MockStore) Read(data interface{}) error {
	m.ChangeEmisor = true
	beforeUpdate := data.(*[]domain.Transanction)
	*beforeUpdate = []domain.Transanction{
		{
			Id:       1,
			Code:     "1",
			Coin:     "COP",
			Amount:   10000,
			Emisor:   "name1",
			Receptor: "MeLi",
			Date:     "25-07-2022",
		},
		{
			Id:       2,
			Code:     "2",
			Coin:     "COP",
			Amount:   10000,
			Emisor:   "name2",
			Receptor: "MeLi",
			Date:     "25-07-2022",
		},
	}
	return nil
}

func TestUpdateMock(t *testing.T) {
	//arrange
	myMockSoter := MockStore{}
	update := NewRepository(&myMockSoter)
	expect := domain.Transanction{
		Id:       1,
		Code:     "1",
		Coin:     "COP",
		Amount:   10000,
		Emisor:   "NewName",
		Receptor: "MeLi",
		Date:     "25-07-2022",
	}
	// act
	afterUpdate, err := update.Update(1, "1", "COP", 10000, "NewName", "MeLi", "25-07-2022")

	//assert

	assert.Equal(t, expect, afterUpdate)
	assert.True(t, myMockSoter.ChangeEmisor)
	assert.Nil(t, err)

}
