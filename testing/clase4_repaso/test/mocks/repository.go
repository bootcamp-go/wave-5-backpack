package mocks

import (
	"clase4_repaso/internal/domain"
	"fmt"
)

type MockRepository struct {
	DataMock []domain.Product
	Error    string
}

func (m *MockRepository) GetAll() ([]domain.Product, error) {
	if m.Error != "" {
		return nil, fmt.Errorf(m.Error)
	}
	return m.DataMock, nil
}

func (m *MockRepository) Store(id int, name, producType string, count int, price float64) (domain.Product, error) {
	if m.Error != "" {
		return domain.Product{}, fmt.Errorf(m.Error)
	}
	p := domain.Product{
		ID: id,
		Name: name,
		Type: producType,
		Count: count,
		Price: price,
	}
	return p, nil
}

func (m *MockRepository) LastID() (int, error) {
	if m.Error != "" {
		return 0, fmt.Errorf(m.Error)
	}
	return m.DataMock[len(m.DataMock)-1].ID, nil
}
