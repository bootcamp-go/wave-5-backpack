package file

import (
	"encoding/json"
	"fmt"
	"os"
)

type File interface {
	Write(data interface{}) error
	Read(data interface{}) error
	Ping() error
}

type fileStruct struct {
	FilePath string
}

func NewFile(path string) File {
	return &fileStruct{
		FilePath: path,
	}
}

func (fs *fileStruct) Read(data interface{}) error {
	file, err := os.ReadFile(fs.FilePath)
	fmt.Println(err)
	if err != nil {
		return err
	}
	return json.Unmarshal(file, &data)
}

func (fs *fileStruct) Write(data interface{}) error {
	fileData, err := json.MarshalIndent(data, " ", "\t")
	if err != nil {
		return err
	}
	return os.WriteFile(fs.FilePath, fileData, 0644)
}

func (fs *fileStruct) Ping() error {
	_, err := os.ReadFile(fs.FilePath)
	return err
}
