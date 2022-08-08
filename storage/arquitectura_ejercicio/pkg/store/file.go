package store

import (
	"encoding/json"
	"os"
)

type Store interface {
	Write(data interface{}) error
	Ping() error
	Read(data interface{}) error
}

type fileStore struct {
	FilePath string
}

func (fs *fileStore) Write(data interface{}) error {
	dataFile, err := json.Marshal(data)

	if err != nil {
		return err
	}

	return os.WriteFile(fs.FilePath, dataFile, 0644)
}

func (fs *fileStore) Read(data interface{}) error {
	dataFile, err := os.ReadFile(fs.FilePath)
	if err != nil {
		return err
	}

	return json.Unmarshal(dataFile, &data)
}
func (fs *fileStore) Ping() error {
	_, err := os.ReadFile(fs.FilePath)
	if err != nil {
		return err
	}
	return nil
}

func NewStore(filename string) Store {
	return &fileStore{filename}
}
