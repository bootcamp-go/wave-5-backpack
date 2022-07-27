package products

import "fmt"

type MockStorage struct {
	ReadWasCalled bool
	dataMock      []Product
	errWrite      string
	errRead       string
}

func (m *MockStorage) Write(data interface{}) error {
	if m.errWrite != "" {
		return fmt.Errorf(m.errWrite)
	}
	a := data.([]Product)
	m.dataMock = a
	return nil
}

func (m *MockStorage) Read(data interface{}) error {
	m.ReadWasCalled = true
	if m.errRead != "" {
		return fmt.Errorf(m.errRead)
	}
	a := data.(*[]Product)
	*a = m.dataMock
	return nil
}

func (m *MockStorage) Ping() error {
	return nil
}
