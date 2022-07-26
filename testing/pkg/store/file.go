package store

import (
	"encoding/json"
	"os"
)

type Store interface {
	Read(data interface{}) error
	Write(data interface{}) error
	AddMock(mock *Mock)
	ClearMock()
}

type Mock struct {
	Data        []byte
	Error       error
	ReadInvoked bool
}

type FileStore struct {
	FileName string
	Mock     *Mock
}

func NewFileStore(fileName string) Store {
	return &FileStore{FileName: fileName}
}

func (f *FileStore) AddMock(mock *Mock) {
	f.Mock = mock
}

func (f *FileStore) ClearMock() {
	f.Mock = nil
}

func (f *FileStore) Read(data interface{}) error {
	if f.Mock != nil {
		if f.Mock.Error != nil {
			return f.Mock.Error
		}

		f.Mock.ReadInvoked = true
		return json.Unmarshal(f.Mock.Data, &data)
	}
	fileData, err := os.ReadFile(f.FileName)
	if err != nil {
		return err
	}

	return json.Unmarshal(fileData, &data)
}

func (f *FileStore) Write(data interface{}) error {
	if f.Mock != nil {
		if f.Mock.Error != nil {
			return nil
		}

		return nil
	}
	dataFile, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return err
	}

	return os.WriteFile(f.FileName, dataFile, 0644)
}
