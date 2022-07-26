package store

import (
	"encoding/json"
	"fmt"
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

func (fs *fileStore) Ping() error {
	if _, err := os.ReadFile(fs.FilePath); err != nil {
		return err
	}
	return nil
}

func (fs *fileStore) Write(data interface{}) error {
	fileData, err := json.Marshal(data)
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
	fmt.Println("voy a deserializar json:" + string(file))
	return json.Unmarshal(file, &data)
}
