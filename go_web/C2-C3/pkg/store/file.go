package store

import (
	"encoding/json"
	"os"
)

type Store interface {
	Read(data interface{}) error
	Write(data interface{}) error
	Open(data interface{}) error
}

func NewStore(pathFile string, fileName string) Store {
	return &fileStore{fileName}
}

type fileStore struct {
	filePath string
}

func (fs *fileStore) Write(data interface{}) error {
	fileData, err := json.MarshalIndent(data, "", "")
	if err != nil {
		return err
	}
	return os.WriteFile(fs.filePath, fileData, 0644)
}

func (fs *fileStore) Read(data interface{}) error {
	file, err := os.ReadFile(fs.filePath)
	if err != nil {
		return err
	}
	return json.Unmarshal(file, &data)
}

func (fs *fileStore) Open(data interface{}) error {
	_, err := os.ReadFile(fs.filePath)
	if err != nil {
		return err
	}
	return nil
}
