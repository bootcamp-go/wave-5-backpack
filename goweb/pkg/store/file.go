package store

import (
	"encoding/json"
	"errors"
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

func NewStore(fileName string) Store {
	return &fileStore{
		FilePath: fileName,
	}
}

func (fs *fileStore) Ping() error {
	if _, err := os.ReadFile(fs.FilePath); err != nil {
		return err
	}

	return nil
}

func (fs *fileStore) Write(data interface{}) error {
	fileData, err := json.Marshal(data)
	if err != nil {
		return errors.New("error al escribir el archivo json")
	}

	return os.WriteFile(fs.FilePath, fileData, 0644)
}

func (fs *fileStore) Read(data interface{}) error {
	file, err := os.ReadFile(fs.FilePath)
	if err != nil {
		return errors.New("error al leer el archivo json")
	}

	return json.Unmarshal(file, &data)
}
