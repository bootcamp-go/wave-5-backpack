package store

import (
	"encoding/json"
	"fmt"
	"testing/4/tm/test_funcional/internal/domain"
)

type MockStore struct {
	ReadWasCalled  bool
	WriteWasCalled bool
	ErrRead        string
	ErrWrite       string
}

var mockProducts []domain.Product

func (m *MockStore) Read(data interface{}) error {
	m.ReadWasCalled = true

	if m.ErrRead != "" {
		return fmt.Errorf(m.ErrRead)
	}

	jsonProducts, _ := json.Marshal(mockProducts)
	return json.Unmarshal(jsonProducts, &data)
}

func (m *MockStore) Write(data interface{}) error {
	m.WriteWasCalled = true

	if m.ErrWrite != "" {
		return fmt.Errorf(m.ErrWrite)
	}

	jsonData, _ := json.Marshal(data)
	return json.Unmarshal(jsonData, &mockProducts)
}
