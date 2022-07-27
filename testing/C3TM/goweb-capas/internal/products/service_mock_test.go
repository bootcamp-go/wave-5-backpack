package products

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServiceIntegrationUpdate(t *testing.T) {
	// arrange
	updateUser := Product{
		ID:    1,
		Name:  "Tablet",
		Type:  "Tecnologia",
		Count: 20,
		Price: 500750.0,
	}
	mockStorage := MockStorage{
		dataMock: []Product{updateUser},
		errWrite: "",
		errRead:  "",
	}
	// act
	repo := NewRepository(&mockStorage)
	service := NewService(repo)
	result, err := service.Update(
		updateUser.ID,
		updateUser.Name,
		updateUser.Type,
		updateUser.Count,
		updateUser.Price)
	// assert
	assert.Nil(t, err)
	assert.True(t, mockStorage.ReadWasCalled)
	assert.Equal(t, mockStorage.dataMock[0], result)
	assert.Equal(t, mockStorage.dataMock[0].ID, 1)
}

func TestServiceIntegrationUpdateFail(t *testing.T) {
	// arrange
	updateUser := Product{
		ID:    1,
		Name:  "Tablet",
		Type:  "Tecnologia",
		Count: 20,
		Price: 500750.0,
	}
	expectedError := fmt.Errorf("cant read database")
	mockStorage := MockStorage{
		dataMock: nil,
		errWrite: "",
		errRead:  "cant read database",
	}
	// act
	repo := NewRepository(&mockStorage)
	service := NewService(repo)
	result, err := service.Update(
		updateUser.ID,
		updateUser.Name,
		updateUser.Type,
		updateUser.Count,
		updateUser.Price)
	// assert
	assert.True(t, mockStorage.ReadWasCalled)
	assert.Equal(t, expectedError, err)
	assert.Equal(t, Product{}, result)
}

func TestServiceIntegrationDelete(t *testing.T) {
	// arrange
	deleteUser := Product{
		ID:    1,
		Name:  "Tablet",
		Type:  "Tecnologia",
		Count: 20,
		Price: 500750.0,
	}
	mockStorage := MockStorage{
		dataMock: []Product{deleteUser},
		errWrite: "",
		errRead:  "",
	}
	// act
	repo := NewRepository(&mockStorage)
	service := NewService(repo)
	err := service.Delete(deleteUser.ID)
	// assert
	assert.Nil(t, err)
	assert.Equal(t, mockStorage.dataMock, []Product{})
}

func TestServiceIntegrationDeleteFail(t *testing.T) {
	// arrange
	deleteUser := Product{
		ID:    1,
		Name:  "Tablet",
		Type:  "Tecnologia",
		Count: 20,
		Price: 500750.0,
	}
	expectedError := fmt.Errorf("product %d not found", 2)
	mockStorage := MockStorage{
		dataMock: []Product{deleteUser},
		errWrite: "",
		errRead:  "",
	}
	// act
	repo := NewRepository(&mockStorage)
	service := NewService(repo)
	err := service.Delete(2)
	// assert
	assert.Equal(t, expectedError, err)
}
