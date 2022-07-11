package registro

import (
	"encoding/json"
	"os"
)

type Registro interface {
	Write(data interface{}) error
	Read(data interface{}) error
}

type FileRegistro struct {
	FileName string
}

type Type string

const (
	FileType Type = "file"
)

func NewFileStore(store Type, fileName string) Registro {
	switch store {
	case FileType:
		return &FileRegistro{FileName: fileName}
	}
	return nil
}

func (f *FileRegistro) Write(data interface{}) error {
	fileData, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return err
	}

	return os.WriteFile(f.FileName, fileData, 0644)
}

func (f *FileRegistro) Read(data interface{}) error {
	fileData, err := os.ReadFile(f.FileName)
	if err != nil {
		return err
	}

	return json.Unmarshal(fileData, &data)
}
