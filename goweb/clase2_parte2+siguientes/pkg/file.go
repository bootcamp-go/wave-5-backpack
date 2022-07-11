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
	fileData, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile(f.FileName, fileData, 0644)
}

func (f *FileStore) Read(data interface{}) error {
	fileData, err := os.ReadFile(f.FileName)
	if err != nil {
		return err
	}
	return json.Unmarshal(fileData, &data)
}
