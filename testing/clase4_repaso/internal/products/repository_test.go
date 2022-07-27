package products

import (
	"clase4_repaso/internal/domain"
	"clase4_repaso/test/mocks"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {
	// arrange
	database := []domain.Product{
		{
			ID:    1,
			Name:  "CellPhone",
			Type:  "Tech",
			Count: 3,
			Price: 250,
		}, {
			ID:    2,
			Name:  "Notebook",
			Type:  "Tech",
			Count: 10,
			Price: 1750.5,
		}}
	mockStorage := mocks.MockStorage{
		DataMock: database,
		ErrWrite: "",
		ErrRead:  "",
	}
	// act
	repo := NewRepository(&mockStorage)
	result, err := repo.GetAll()
	// assert
	assert.Nil(t, err)
	assert.Equal(t, mockStorage.DataMock, result)
}

func TestStore(t *testing.T) {
	// arrange
	database := []domain.Product{
		{
			ID:    1,
			Name:  "CellPhone",
			Type:  "Tech",
			Count: 3,
			Price: 250,
		}, {
			ID:    2,
			Name:  "Notebook",
			Type:  "Tech",
			Count: 10,
			Price: 1750.5,
		}}
	mockStorage := mocks.MockStorage{
		DataMock: database,
		ErrWrite: "",
		ErrRead:  "",
	}

	newProduct := domain.Product{
		ID:    3,
		Name:  "Tablet",
		Type:  "Tech",
		Count: 5,
		Price: 1050.99,
	}
	// act
	repo := NewRepository(&mockStorage)
	result, err := repo.Store(
		newProduct.ID,
		newProduct.Name,
		newProduct.Type,
		newProduct.Count,
		newProduct.Price)
	// assert
	fmt.Printf("\n%+v", mockStorage.DataMock)
	assert.Nil(t, err)
	assert.Equal(t, newProduct, result)
	assert.Equal(t, mockStorage.DataMock[len(mockStorage.DataMock)-1], newProduct)
}
