package store

import (
	"encoding/json"
	"os"
)

type Store interface {
	Read(data interface{}) error
	Write(data interface{}) error
}

func NewStore(fileName string) Store {
	return &fileStore{fileName}
}

type fileStore struct {
	filePath string
}

func (f *fileStore) Write(data interface{}) error {
	fileData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return os.WriteFile(f.filePath, fileData, 0644)
}

func (f *fileStore) Read(data interface{}) error {
	file, err := os.ReadFile(f.filePath)
	if err != nil {
		return err
	}
	return json.Unmarshal(file, &data)
}
