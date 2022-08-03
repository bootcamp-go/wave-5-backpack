package store

import (
	"encoding/json"
	"os"
)

type Store interface {
	Read(data interface{}) error
	Write(data interface{}) error
}

type FileStore struct {
	FileName string
}

func NewFileStore(fileName string) Store {
	return &FileStore{FileName: fileName}
}

func (f *FileStore) Read(data interface{}) error {
	fileData, err := os.ReadFile(f.FileName)
	if err != nil {
		return err
	}

	return json.Unmarshal(fileData, &data)
}

func (f *FileStore) Write(data interface{}) error {
	dataFile, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return err
	}

	return os.WriteFile(f.FileName, dataFile, 0644)
}
