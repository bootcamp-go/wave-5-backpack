package store

import (
	"encoding/json"
	"os"
)

type Store interface {
	Write(data interface{}) error
	Read(data interface{}) error
	Ping() error
}

type fileStore struct {
	filePath string
}

func NewStore(filepath string) Store {
	return &fileStore{
		filePath: filepath,
	}
}

func (f *fileStore) Ping() error {
	_, err := os.ReadFile(f.filePath)
	return err
}

func (f *fileStore) Write(data interface{}) error {
	writeData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return os.WriteFile(f.filePath, writeData, 0644)
}

func (f *fileStore) Read(data interface{}) error {
	readData, err := os.ReadFile(f.filePath)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(readData, data); err != nil {
		return err
	}
	return nil
}
