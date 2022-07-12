package store

import (
	"encoding/json"
	"os"
)

type Store interface {
	Read(data interface{}) error
	Write(data interface{}) error
}

type fileStore struct {
	filePath string
}

func NewStore(fileName string) Store {
	return &fileStore{fileName}
}

func (fs *fileStore) Read(data interface{}) error {
	file, err := os.ReadFile(fs.filePath)
	if err != nil {
		return err
	}
	return json.Unmarshal(file, &data)
}

func (fs *fileStore) Write(data interface{}) error {
	file, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return err
	}
	return os.WriteFile(fs.filePath, file, 0644)
}
