package repository

import (
	"encoding/json"
	"goweb/4/tt/internal/domain"
)

type MockStore struct {
	ReadWasCalled  bool
	WriteWasCalled bool
	mockProducts   []domain.Product
}

func (m *MockStore) Read(data interface{}) error {
	m.ReadWasCalled = true
	m.mockProducts = append(m.mockProducts, domain.NewProduct(1, "Before Update", 1.5, 2))
	jsonProducts, _ := json.Marshal(m.mockProducts)
	return json.Unmarshal(jsonProducts, &data)
}

func (m *MockStore) Write(data interface{}) error {
	m.WriteWasCalled = true
	jsonData, _ := json.Marshal(data)
	return json.Unmarshal(jsonData, &m.mockProducts)
}
