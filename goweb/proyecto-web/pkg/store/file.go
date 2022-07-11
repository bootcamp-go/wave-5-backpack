package store

import (
	"encoding/json"
	"os"
)

type Store interface {
	Read(data interface{}) error
	Write(date interface{}) error
}

type fileStore struct {
	FilePath string
}

func NewStore(filePath string) Store {
	return &fileStore{FilePath: filePath}
}

func (fs *fileStore) Read(data interface{}) error {
	fileBytes, err := os.ReadFile(fs.FilePath)
	if err != nil {
		return err
	}
	return json.Unmarshal(fileBytes, data)
}

func (fs *fileStore) Write(data interface{}) error {
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return os.WriteFile(fs.FilePath, dataBytes, 0644)
}

func (fs *fileStore) Ping() error {
	err := fs.Read(fs.FilePath)
	if err != nil {
		return err
	}
	return nil
}
