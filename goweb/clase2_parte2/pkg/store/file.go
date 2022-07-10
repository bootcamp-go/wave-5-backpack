package store

import (
	"encoding/json"
	"os"
)

type Store interface {
	Read(data interface{}) error
	Write(data interface{}) error
	Val() error
}

func NewStore(fileName string) Store {
	return &fileStore{fileName}
}

type fileStore struct {
	FilePath string
}

func (fs *fileStore) Val() error {
	if _, err := os.ReadFile(fs.FilePath); err != nil {
		return err
	}
	return nil
}
func (fs *fileStore) Read(data interface{}) error {
	file, err := os.ReadFile(fs.FilePath)
	if err != nil {
		return err
	}
	return json.Unmarshal(file, &data)
}

func (fs *fileStore) Write(data interface{}) error {
	fileData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return os.WriteFile(fs.FilePath, fileData, 0644)
}
