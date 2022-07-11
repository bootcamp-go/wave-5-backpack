package store

import (
	"encoding/json"
	"os"
)

// --------------------------------------------
// --------------- Estructuras ----------------
// --------------------------------------------

// Clase 3 Ejercicio 1 Parte 2
type Store interface {
	Read(data interface{}) error
	Write(data interface{}) error
	Ping() error
}

type fileStore struct {
	FilePath string
}

func NewStore(filename string) Store {
	return &fileStore{FilePath: filename}
}

// --------------------------------------------
// ----------------- Métodos ------------------
// --------------------------------------------

func (fs *fileStore) Ping() error {
	if _, err := os.ReadFile(fs.FilePath); err != nil {
		return err
	}
	return nil
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
