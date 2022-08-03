package store

import (
	"encoding/json"
	"os"
)

type Store interface {
	Write(data interface{}) error
	Read(data interface{}) error
}

type FileStore struct {
	FileName string
	Mock     *Mock
}

type Mock struct {
	Data        []byte
	Error       error
	ReadInvoked bool
}

type Type string

const (
	FileType Type = "file"
)

func NewFileStore(store Type, fileName string) Store {
	switch store {
	case FileType:
		return &FileStore{FileName: fileName}
	}
	return nil
}

func (f *FileStore) Write(data interface{}) error {
	if f.Mock != nil {
		if f.Mock.Error != nil {
			return f.Mock.Error
		}

		return nil
	}
	fileData, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return err
	}

	return os.WriteFile(f.FileName, fileData, 0644)
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
