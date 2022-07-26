package products

import (
	"clase3_parte1/internal/domain"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServiceIntegrationGetAll(t *testing.T) {
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
	mockStorage := MockStorage{
		dataMock: database,
		errWrite: "",
		errRead:  "",
	}
	// act
	repo := NewRepository(&mockStorage)
	service := NewService(repo)
	result, err := service.GetAll()
	// assert
	assert.Nil(t, err)
	assert.Equal(t, mockStorage.dataMock, result)
}

func TestServiceIntegrationGetAllFail(t *testing.T) {
	// arrange
	expectedError := errors.New("cant read database")
	mockStorage := MockStorage{
		dataMock: nil,
		errWrite: "",
		errRead:  "cant read database",
	}
	// act
	repo := NewRepository(&mockStorage)
	service := NewService(repo)
	result, err := service.GetAll()
	// assert
	assert.ErrorContains(t, err, expectedError.Error())
	assert.Nil(t, result)
}

func TestServiceIntegrationStore(t *testing.T) {
	// arrange
	newProduct := domain.Product{
		ID:    0,
		Name:  "Tablet",
		Type:  "Tech",
		Count: 5,
		Price: 1050.99,
	}
	mockStorage := MockStorage{
		dataMock: nil,
		errWrite: "",
		errRead:  "",
	}
	// act
	repo := NewRepository(&mockStorage)
	service := NewService(repo)
	result, err := service.Store(
		newProduct.Name,
		newProduct.Type,
		newProduct.Count,
		newProduct.Price)
	// assert
	assert.Nil(t, err)
	assert.Equal(t, mockStorage.dataMock[0], result)
	assert.Equal(t, mockStorage.dataMock[0].ID, 1)
}

func TestServiceIntegrationStoreFail(t *testing.T) {
	// arrange
	newProduct := domain.Product{
		ID:    0,
		Name:  "Tablet",
		Type:  "Tech",
		Count: 5,
		Price: 1050.99,
	}
	writeErr := fmt.Errorf("cant write database")
	expectedError := fmt.Errorf("error creating product: %w", writeErr)
	mockStorage := MockStorage{
		dataMock: nil,
		errWrite: "cant write database",
		errRead:  "",
	}
	// act
	repo := NewRepository(&mockStorage)
	service := NewService(repo)
	result, err := service.Store(
		newProduct.Name,
		newProduct.Type,
		newProduct.Count,
		newProduct.Price)
	// assert
	assert.ErrorContains(t, err, expectedError.Error())
	assert.Equal(t, domain.Product{}, result)
}
