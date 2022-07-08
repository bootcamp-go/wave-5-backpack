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

func NewStore(fileName string) Store {
	return &fileStore{fileName}
}

type fileStore struct {
	FilePath string
}

func (fs *fileStore) Read(data interface{}) error {
	file, err := os.ReadFile(fs.FilePath)
	if err != nil {
		return err
	}

	return json.Unmarshal(file, data)
}

func (fs *fileStore) Write(data interface{}) error {
	file, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	return os.WriteFile(fs.FilePath, file, 0644)
}

func (fs *fileStore) Ping() error {
	if _, err := os.ReadFile(fs.FilePath); err != nil {
		return err
	}
	return nil
}
