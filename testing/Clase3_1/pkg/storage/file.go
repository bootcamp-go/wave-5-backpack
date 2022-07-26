package storage

import (
	"encoding/json"
	"os"
)

type Store interface {
	Read(data interface{}) error
	Write(data interface{}) error
}

type fileStore struct {
	FilePath string
}

func NewStore(fileName string) Store {
	return &fileStore{fileName}
}

func (f *fileStore) Read(data interface{}) error {
	file, err := os.ReadFile(f.FilePath)
	if err != nil {
		return err
	}
	return json.Unmarshal(file, &data)
}
func (f *fileStore) Write(data interface{}) error {
	file, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return os.WriteFile(f.FilePath, file, 0644)

}
