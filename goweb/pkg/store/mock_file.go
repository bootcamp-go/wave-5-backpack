package store

import (
	"fmt"
	"goweb/internal/domain"
)

type MockStorage struct {
	DataMock []domain.Product
	ErrWrite string
	ErrRead  string
	ReadFile bool
 }
 
 func (m *MockStorage) Read(data interface{}) error {
	if m.ErrRead != "" {
		return fmt.Errorf(m.ErrRead)
	}
	a := data.(*[]domain.Product)
	*a = m.DataMock
	m.ReadFile = true
	return nil
 }
 
 func (m *MockStorage) Write(data interface{}) error {
	if m.ErrWrite != "" {
		return fmt.Errorf(m.ErrWrite)
	}
	a := data.([]domain.Product)
	m.DataMock = append(m.DataMock, a...)
	return nil
 }
 
 func (m *MockStorage) Ping() error  {
	//	err := os.OpenFile()
		return nil
	}