package products

import (
	"clase4_parte1/internal/domain"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestIntegrationServiceGetAll(t *testing.T) {
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
	dbMock := StoreMock{Data: input}
	myRepo := NewRepository(&dbMock)
	myService := NewService(myRepo)

	result, err := myService.GetAll()

	assert.Equal(t, input, result)
	assert.Nil(t, err)
}

func TestServiceGetAllError(t *testing.T) {
	expectedError := errors.New("error for GetAll")
	dbMock := StoreMock{Err: expectedError}
	myRepo := NewRepository(&dbMock)
	myService := NewService(myRepo)

	result, err := myService.GetAll()

	assert.Equal(t, expectedError, err)
	assert.Nil(t, result)
}

func TestStore(t *testing.T) {
	testProduct := domain.Product{
		Name:  "CellPhone",
		Type:  "Tech",
		Count: 3,
		Price: 52.0,
	}
	dbMock := StoreMock{Data: []domain.Product{testProduct}}
	myRepo := NewRepository(&dbMock)
	myService := NewService(myRepo)
	result, _ := myService.Store(testProduct.Name, testProduct.Type,
		testProduct.Count, testProduct.Price)
	assert.Equal(t, testProduct.Name, result.Name)
	assert.Equal(t, testProduct.Type, result.Type)
	assert.Equal(t, testProduct.Price, result.Price)
	assert.Equal(t, 1, result.ID)
}

func TestStoreError(t *testing.T) {
	testProduct := domain.Product{
		Name:  "CellPhone",
		Type:  "Tech",
		Count: 3,
		Price: 52.0,
	}
	expectedError := errors.New("error for Storage")
	dbMock := StoreMock{
		Err: expectedError,
	}
	myRepo := NewRepository(&dbMock)
	myService := NewService(myRepo)
	result, err := myService.Store(testProduct.Name, testProduct.Type,
		testProduct.Count, testProduct.Price)
	assert.Equal(t, expectedError, err)
	assert.Equal(t, domain.Product{}, result)
}

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) GetAll() ([]domain.Product, error) {
	args := m.Called()
	return args.Get(0).([]domain.Product), args.Error(1)
}

func (m *MockRepository) Store(id int, nombre, tipo string, cantidad int, precio float64) (domain.Product, error) {
	args := m.Called(id, nombre, tipo, cantidad, precio)
	return args.Get(0).(domain.Product), args.Error(1)
}

func (m *MockRepository) LastID() (int, error) {
	args := m.Called()
	return args.Get(0).(int), args.Error(1)

}

func (m *MockRepository) Update(id int, name, productType string, count int, price float64) (domain.Product, error) {
	args := m.Called(id, name, productType, count, price)
	return args.Get(0).(domain.Product), args.Error(1)
}

func (m *MockRepository) UpdateName(id int, name string) (domain.Product, error) {
	args := m.Called(id, name)
	return args.Get(0).(domain.Product), args.Error(1)
}

func (m *MockRepository) Delete(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestServiceGetAll(t *testing.T) {
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
	repo := new(MockRepository)
	repo.On("GetAll").Return(input, nil)
	myService := NewService(repo)

	productResult, err := myService.GetAll()

	assert.Nil(t, err)
	assert.Equal(t, input, productResult)

}

func TestServiceCreate(t *testing.T) {
	input := domain.Product{
		ID:    1,
		Name:  "CellPhone",
		Type:  "Tech",
		Count: 3,
		Price: 250,
	}

	repo := new(MockRepository)
	repo.On("LastID").Return(0, nil)
	repo.On("Store", input.ID, input.Name, input.Type, input.Count, input.Price).Return(input, nil)

	myService := NewService(repo)

	productResult1, err := myService.Store(input.Name, input.Type, input.Count, input.Price)

	assert.Nil(t, err)
	assert.Equal(t, input, productResult1)

}

func TestServiceCreateErrorOnCreate(t *testing.T) {
	input := domain.Product{
		ID:    1,
		Name:  "CellPhone",
		Type:  "Tech",
		Count: 3,
		Price: 250,
	}
	repo := new(MockRepository)
	repo.On("LastID").Return(0, nil)
	repo.On("Store", input.ID, input.Name, input.Type, input.Count, input.Price).Return(domain.Product{}, fmt.Errorf("error on store"))

	myService := NewService(repo)

	_, err := myService.Store(input.Name, input.Type, input.Count, input.Price)

	assert.NotNil(t, err)

}
