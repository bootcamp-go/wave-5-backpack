package products

import (
	"goweb/internal/domain"
	"goweb/pkg/store"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {
db := []domain.Product{
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

mockStorage := store.MockStorage{
	DataMock: db,
	ErrWrite: "",
	ErrRead: "",
}

repo := NewRepository(&mockStorage)

response, err := repo.GetAll()

assert.Equal(t, mockStorage.DataMock, response)
assert.Nil(t, err)
assert.True(t, mockStorage.ReadFile)
}

func TestUpdate(t *testing.T) {
	db := []domain.Product{
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
	
	mockStorage := store.MockStorage{
		DataMock: db,
		ErrWrite: "",
		ErrRead: "",
	}
	
	repo := NewRepository(&mockStorage)
	product := domain.Product{Id: 1,
							  Name: "Update After",
							  Color: "Blanco",
							  Price: 1000,
							  Stock: 25,
							  Code: "NXV321",
							  Publisher: false,
							  CreatedAt: "2006-01-02T15:04:05Z07:00"}

	response, err := repo.Update(product.Id, product.Name, product.Color, product.Price, product.Stock, product.Code, product.Publisher)

	assert.Equal(t, product, response)
	assert.Nil(t, err)
	assert.True(t, mockStorage.ReadFile)
}

func TestDelete(t *testing.T) {
	db := []domain.Product{
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
	
	mockStorage := store.MockStorage{
		DataMock: db,
		ErrWrite: "",
		ErrRead: "",
	}

	repo := NewRepository(&mockStorage)

	err := repo.Delete(1)

	assert.Nil(t, err)
	assert.True(t, mockStorage.ReadFile)
}