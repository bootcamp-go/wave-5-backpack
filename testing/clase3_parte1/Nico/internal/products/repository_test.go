package products

import (
	"clase3_parte1/internal/domain"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockStorage struct {
	dataMock []domain.Product
	errWrite string
	errRead  string
}

func (m *MockStorage) Read(data interface{}) error {
	if m.errRead != "" {
		return fmt.Errorf(m.errRead)
	}
	a := data.(*[]domain.Product)
	*a = m.dataMock
	return nil
}

func (m *MockStorage) Write(data interface{}) error {
	if m.errWrite != "" {
		return fmt.Errorf(m.errWrite)
	}
	a := data.([]domain.Product)
	m.dataMock = append(m.dataMock, a[len(a)-1])
	return nil
}

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
	mockStorage := MockStorage{
		dataMock: database,
		errWrite: "",
		errRead:  "",
	}
	// act
	repo := NewRepository(&mockStorage)
	result, err := repo.GetAll()
	// assert
	assert.Nil(t, err)
	assert.Equal(t, mockStorage.dataMock, result)
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
	mockStorage := MockStorage{
		dataMock: database,
		errWrite: "",
		errRead:  "",
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
	fmt.Printf("\n%+v", mockStorage.dataMock)
	assert.Nil(t, err)
	assert.Equal(t, newProduct, result)
	assert.Equal(t, mockStorage.dataMock[len(mockStorage.dataMock)-1], newProduct)
}
