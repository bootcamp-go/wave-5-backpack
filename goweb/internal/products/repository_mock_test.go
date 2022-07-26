package products

import (
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

	update, err := repo.ParcialUpdate(1, "Update After", 0)

	assert.Nil(t, err)
	assert.Equal(t,"Update After", update.Name)
	assert.True(t, mock.ReadStore)
}

