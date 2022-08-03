package mocks

import (
	"clase4_repaso/internal/domain"
	"fmt"
)

type MockStorage struct {
	DataMock []domain.Product
	ErrWrite string
	ErrRead  string
}

func (m *MockStorage) Read(data interface{}) error {
	if m.ErrRead != "" {
		return fmt.Errorf(m.ErrRead)
	}
	a := data.(*[]domain.Product)
	*a = m.DataMock
	return nil
}

func (m *MockStorage) Write(data interface{}) error {
	if m.ErrWrite != "" {
		return fmt.Errorf(m.ErrWrite)
	}
	a := data.([]domain.Product)
	m.DataMock = append(m.DataMock, a[len(a)-1])
	return nil
}
