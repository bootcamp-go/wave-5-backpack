package store

import (
	"errors"
	"reflect"
)

type Mock struct {
	ReadFlag  bool
	Db        interface{}
	FailRead  bool
	FailWrite bool
}

func (s *Mock) Ping() error {
	return nil
}
func (s *Mock) Read(data interface{}) error {
	if s.FailRead {
		return errors.New("error: at read")
	}
	s.ReadFlag = true
	dbData := reflect.ValueOf(s.Db).Elem()
	reflect.ValueOf(data).Elem().Set(dbData)
	return nil
}
func (s *Mock) Write(data interface{}) error {
	if s.FailWrite {
		return errors.New("error: at write")
	}
	reflect.ValueOf(s.Db).Elem().Set(reflect.ValueOf(data))
	return nil
}
