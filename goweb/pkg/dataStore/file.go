package dataStore

import (
	"encoding/json"
	"os"
)

type DataStore interface {
	Read(data interface{}) error
	Write(data interface{}) error
	Ping() error
}

type Type string

const (
	FileType Type = "filestorage"
)

func NewStore(store Type, fileName string) DataStore {
	switch store {
	case FileType:
		return &FileStore{fileName}
	}
	return nil
}

type FileStore struct {
	FileName string
}

func (fs *FileStore) Ping() error {
	if _, err := os.ReadFile(fs.FileName); err != nil {
		return err
	}
	return nil
}

func (fs *FileStore) Write(data interface{}) error {
	fileData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return os.WriteFile(fs.FileName, fileData, 0644)
}

func (fs *FileStore) Read(data interface{}) error {
	file, err := os.ReadFile(fs.FileName)
	if err != nil {
		return err
	}
	return json.Unmarshal(file, data)
}
