package bank

import (
	"encoding/json"
	"os"
)

type Bank interface {
	Read(data interface{}) error
	Write(data interface{}) error
	Ping() error
}

type Type string

const (
	FileType Type = "filestorage"
)

type fileBank struct {
	FilePath string
}

func NewBank(fileName string) Bank {
	return &fileBank{fileName}
}

func New(store Type, fileName string) Bank {
	switch store {
	case FileType:
		return &fileBank{fileName}
	}
	return nil
}

func (fs *fileBank) Ping() error {
	if _, err := os.ReadFile(fs.FilePath); err != nil {
		return err
	}
	return nil
}

func (fs *fileBank) Write(data interface{}) error {
	fileData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return os.WriteFile(fs.FilePath, fileData, 0644)
}

func (fs *fileBank) Read(data interface{}) error {
	file, err := os.ReadFile(fs.FilePath)
	if err != nil {
		return err
	}
	return json.Unmarshal(file, &data)
}
