package mocks

import (
	"clase4_repaso/internal/domain"
	"fmt"
)

type MockService struct {
	DataMock []domain.Product
	Error string
}

func (m *MockService) GetAll() ([]domain.Product, error) {
	if m.Error != "" {
		return nil, fmt.Errorf(m.Error)
	}
	return m.DataMock, nil

}

func (m *MockService) Store(name, productType string, count int, price float64) (domain.Product, error) {
	if m.Error != "" {
		return domain.Product{}, fmt.Errorf(m.Error)
	}
	p := domain.Product{
		Name: name,
		Type: productType,
		Count: count,
		Price: price,
	}
	return p, nil
}
