package products

import (
	"errors"
	"fmt"
	"goweb/internal/domain"
	"goweb/pkg/store"
	"testing"

	"github.com/stretchr/testify/assert"
)

var dbs = []domain.Product{
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

func TestServiceIntegrationGetAll(t *testing.T) {

mockStorage := store.MockStorage{
	DataMock: dbs,
	ErrWrite: "",
	ErrRead: "",
}

repo := NewRepository(&mockStorage)
service := NewService(repo)

response, err := service.GetAll()

assert.Equal(t, mockStorage.DataMock, response)
assert.Nil(t, err)
assert.True(t, mockStorage.ReadFile)
}

func TestServiceIntegrationGetAllFail(t *testing.T) {
	
	expectedError := errors.New("cannot read the file")
	mockStorage := store.MockStorage{
		DataMock: nil,
		ErrWrite: "",
		ErrRead:  "cannot read the file",
	}
	
	repo := NewRepository(&mockStorage)
	service := NewService(repo)
	result, err := service.GetAll()
	// assert
	assert.Equal(t, expectedError, err)
	assert.Nil(t, result)
 }

 func TestServiceIntegrationUpdate(t *testing.T) {	
	mockStorage := store.MockStorage{
		DataMock: dbs,
		ErrWrite: "",
		ErrRead: "",
	}
	
	repo := NewRepository(&mockStorage)
	service := NewService(repo)
	product := domain.Product{Id: 1,
							  Name: "Update After",
							  Color: "Blanco",
							  Price: 1000,
							  Stock: 25,
							  Code: "NXV321",
							  Publisher: false,
							  CreatedAt: "2006-01-02T15:04:05Z07:00"}

	response, err := service.Update(product.Id, product.Name, product.Color, product.Price, product.Stock, product.Code, product.Publisher)

	assert.Equal(t, product, response)
	assert.Nil(t, err)
	assert.True(t, mockStorage.ReadFile)
}

func TestServiceIntegrationDelete(t *testing.T) {
	mockStorage := store.MockStorage{
		DataMock: dbs,
		ErrWrite: "",
		ErrRead: "",
	}

	repo := NewRepository(&mockStorage)
	service := NewService(repo)
	
	err := service.Delete(1)

	assert.Nil(t, err)
	assert.True(t, mockStorage.ReadFile)
}

func TestServiceIntegrationDeleteFail(t *testing.T) {
	mockStorage := store.MockStorage{
		DataMock: dbs,
		ErrWrite: "",
		ErrRead: "",
	}
	
	id := 3
	expetErr := fmt.Errorf("producto %d no encontrado", id)
	repo := NewRepository(&mockStorage)
	service := NewService(repo)
	

	err := service.Delete(id)
	assert.NotNil(t, err)
	assert.ErrorContains(t, err, expetErr.Error())
}
 