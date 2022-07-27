package products

import (
	"clase4_parte1/internal/domain"
	"testing"

	"github.com/stretchr/testify/assert"
)

type StoreMock struct {
	Data []domain.Product
	Err  error
}

func (s StoreMock) Read(data interface{}) error {
	if s.Err != nil {
		return s.Err
	}
	ref := data.(*[]domain.Product)
	*ref = s.Data
	return nil
}

func (s *StoreMock) Write(data interface{}) error {
	if s.Err != nil {
		return s.Err
	}
	s.Data = data.([]domain.Product)
	return nil
}

func TestGetAll(t *testing.T) {
	// Initializing input/output
	input := []domain.Product{
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
		},
	}
	mockStore := StoreMock{
		Data: input,
	}
	myRepo := NewRepository(&mockStore)
	// Test Execution
	resp, _ := myRepo.GetAll()
	// Validation
	assert.Equal(t, input, resp)
}

func TestDelete(t *testing.T) {
	// Initializing input/output
	input := []domain.Product{
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
		},
	}
	lenExpected := len(input) - 1
	mockStore := StoreMock{
		Data: input,
	}
	myRepo := NewRepository(&mockStore)
	// Test Execution
	err := myRepo.Delete(1)
	// Validation
	assert.Nil(t, err)
	assert.Equal(t, lenExpected, len(mockStore.Data))
}

func TestUpdateName(t *testing.T) {
	// Initializing input/output
	input := []domain.Product{
		{
			ID:    1,
			Name:  "CellPhone",
			Type:  "Tech",
			Count: 3,
			Price: 250,
		},
	}
	newName := "Laptop"
	mockStore := StoreMock{
		Data: input,
	}
	myRepo := NewRepository(&mockStore)
	// Test Execution
	product, err := myRepo.UpdateName(1, newName)
	// Validation
	assert.Nil(t, err)
	assert.Equal(t, newName, product.Name)
}
