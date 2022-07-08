package store

import (
	"encoding/json"
	"os"
)

type fileStore struct {
	FilePath string
}



type Store interface {
	Write(data interface{}) error
	Read(data interface{}) error
	Ping() error
}

func NewStore(filePath string) Store {
	return &fileStore{FilePath: filePath}

}

// Read implements Store
func (fs *fileStore) Read(data interface{}) error {
	file, err := os.ReadFile(fs.FilePath)

	if err != nil{
		return err
	}
	return json.Unmarshal(file, &data)
}

func (fs *fileStore) Ping() error  {
//	err := os.OpenFile()
	return nil
}

func (fs *fileStore) Write(data interface{}) error {
	fileData, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return err
	}

	return os.WriteFile(fs.FilePath, fileData, 0644)
}
