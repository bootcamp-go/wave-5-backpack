package transactions

import (
	"GoWeb/internals/domain"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockStore struct {
	ReadWasCalled bool
	dataMock      []domain.Transanction
	errRead       string
	errWrite      string
}

func (m *MockStore) Ping() error {
	return nil
}

func (m *MockStore) Write(data interface{}) error {

	if m.errWrite != "" {
		return fmt.Errorf(m.errWrite)
	}
	a := data.([]domain.Transanction)
	m.dataMock = append(m.dataMock, a...)
	return nil
}

func (m *MockStore) Read(data interface{}) error {

	if m.errRead != "" {
		return fmt.Errorf(m.errWrite)
	}
	m.ReadWasCalled = true
	beforeUpdate := data.(*[]domain.Transanction)
	*beforeUpdate = m.dataMock
	return nil
}

func TestGetAllTransaction(t *testing.T) {
	//arrange
	data := []domain.Transanction{
		{Id: 1, Code: "QWE123", Coin: "COP", Amount: 1000, Emisor: "Juan", Receptor: "MeLi", Date: "07-27-2022"},
		{Id: 2, Code: "QWE123", Coin: "COP", Amount: 2000, Emisor: "David", Receptor: "MeLi", Date: "07-27-2022"},
	}

	mock := MockStore{
		dataMock: data,
	}
	repo := NewRepository(&mock)
	esperado := data

	// act
	tran, err := repo.GetAll()

	//assert
	assert.Nil(t, err)
	assert.Equal(t, esperado, tran)
}

func TestUpdate(t *testing.T) {

	//arrange
	data := []domain.Transanction{
		{Id: 1, Code: "QWE123", Coin: "COP", Amount: 1000, Emisor: "Juan", Receptor: "MeLi", Date: "07-27-2022"},
		{Id: 2, Code: "QWE123", Coin: "COP", Amount: 2000, Emisor: "David", Receptor: "MeLi", Date: "07-27-2022"},
	}

	mock := MockStore{
		dataMock: data,
	}
	repo := NewRepository(&mock)

	// act
	afterUpdate, err := repo.Update(1, "AAA", "USD", 3000, "NewName", "MeLi", "07-28-2022")

	//assert
	assert.Equal(t, mock.dataMock[0], afterUpdate)
	assert.True(t, mock.ReadWasCalled)
	assert.Nil(t, err)

}
