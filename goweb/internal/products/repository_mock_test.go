package products

import (
	"fmt"
	"goweb/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)




type MockStore struct {
	ReadStore bool
}

func (mock *MockStore) Read(data interface{}) error {
	 mock.ReadStore = true
	 p := data.(*[]domain.Product)
	 *p = []domain.Product{
		 {Id: 1, 
		 Name: "Update Before", 
		 Color: "azul",
		 Price: 1500, 
		 Stock: 100, 
		 Code: "AFN123", 
		 Publisher: true, 
		 CreatedAt: "2006-01-02T15:04:05Z07:00"},
		 {Id: 2, 
		 Name: "product2", 
		 Color: "blanco",
		 Price: 1200, 
		 Stock: 50, 
		 Code: "BFN123", 
		 Publisher: false, 
		 CreatedAt: "2006-01-02T15:04:05Z07:00"},
	 }
	return nil
}

func (mock *MockStore) Write(data interface{}) error {
	return nil
}

func (mock *MockStore) Ping() error  {
	//	err := os.OpenFile()
		return nil
}

func TestParcialUpdate(t *testing.T)  {

	
	mock := MockStore{}

	repo := NewRepository(&mock)
	id,name, price := 1, "aaaaa", 2888.00
	update, err := repo.ParcialUpdate(id, name, price)

	assert.Nil(t, err)
	if name != "" {
		assert.Equal(t,name, update.Name)
	}
	if price > 0 {
		assert.Equal(t, price, update.Price)
	}
	if name != "" && price > 0 {
		assert.Equal(t,name, update.Name)
		assert.Equal(t, price, update.Price)
	}
	
	assert.True(t, mock.ReadStore)
}

func TestParcialUpdateFail(t *testing.T) {
	mock := MockStore{}

	repo := NewRepository(&mock)
	id,name, price := 3, "aaaa", 12220.00
	_, err := repo.ParcialUpdate(id, name, price)
	errs := fmt.Errorf("producto %d no encontrado", id)
	assert.NotNil(t, err)
	assert.ErrorContains(t, err, errs.Error())

}

