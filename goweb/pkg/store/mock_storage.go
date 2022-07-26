package store

import "reflect"

type MockStorage struct {
	DataMock      interface{}
	ReadWasCalled bool
}

func (ms *MockStorage) Write(data interface{}) error {
	ms.DataMock = data
	return nil
}

func (ms *MockStorage) Read(data interface{}) error {
	rv := reflect.Indirect(reflect.ValueOf(data))
	rv.Set(reflect.ValueOf(ms.DataMock))
	ms.ReadWasCalled = true
	return nil
}
