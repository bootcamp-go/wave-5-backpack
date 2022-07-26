package products

import (
	"goweb/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

type StubTest struct{}

func (st StubTest) Read(data interface{}) error {
	p := data.(*[]domain.Product)
	*p = []domain.Product{
		{Id: 1, 
		Name: "product1", 
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

func (st StubTest) Ping() error  {
//	err := os.OpenFile()
	return nil
}

func (st StubTest) Write(data interface{}) error {
	return nil
}



func TestStubGetAll(t *testing.T) {
	myStub := StubTest{}
	repo := NewRepository(myStub)
	esperado := []domain.Product{
		{Id: 1, 
		Name: "product1", 
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
	resultado, err := repo.GetAll()
	assert.Nil(t, err)
	assert.Equal(t, esperado, resultado)

}