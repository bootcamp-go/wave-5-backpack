package store

import (
	"encoding/json"
	"os"
)

type Store interface {
	Read(data interface{}) error
	Write(data interface{}) error
	Ping() error
}

type fileStore struct {
	FilePath string
}

func NewStore(filePath string) Store {
	return &fileStore{filePath}
}

func (fs *fileStore) Write(data interface{}) error {
	file, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return os.WriteFile(fs.FilePath, file, 0644)
}

func (fs *fileStore) Read(data interface{}) error {
	file, err := os.ReadFile(fs.FilePath)
	if err != nil {
		return err
	}

	return json.Unmarshal(file, &data)
}

func (fs *fileStore) Ping() error {
	_, err := os.ReadFile(fs.FilePath)
	if err != nil {
		return err
	}

	return nil
}
