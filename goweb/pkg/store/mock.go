package store

import "reflect"

type Mock struct {
	ReadFlag bool
	Db       interface{}
}

func (s *Mock) Ping() error {
	return nil
}
func (s *Mock) Read(data interface{}) error {
	s.ReadFlag = true
	dbData := reflect.ValueOf(s.Db).Elem()
	reflect.ValueOf(data).Elem().Set(dbData)
	return nil
}
func (s *Mock) Write(data interface{}) error {
	reflect.ValueOf(s.Db).Elem().Set(reflect.ValueOf(data))
	return nil
}
