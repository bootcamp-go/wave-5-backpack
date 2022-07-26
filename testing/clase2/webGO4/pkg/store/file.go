package store

import (
	"encoding/json"
	"log"
	"os"
)

type Store interface {
	Read(data interface{}) error
	Write(data interface{}) error
}
type fileStore struct {
	FilePath string
}

func NewStore(pathFile string) Store {
	_, err := os.OpenFile(pathFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	return &fileStore{FilePath: pathFile}
}
func (fs *fileStore) Write(data interface{}) error {
	fileData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(fs.FilePath, fileData, 0644)
}
func (fs *fileStore) Read(data interface{}) error {
	file, err := os.ReadFile(fs.FilePath)
	if err != nil {
		return err
	}
	return json.Unmarshal(file, &data)
}
