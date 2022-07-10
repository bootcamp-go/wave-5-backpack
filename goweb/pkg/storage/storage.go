package storage

import (
	"encoding/json"
	"fmt"
	"os"
)

// Read recibe una interfaz vacia, este será el puntero donde guardaremos la data leida
type Storage interface {
  Read(p interface{}) error
  Write(data interface{}) error
}

func NewStorage(filePath string) Storage {
  return &storage{filePath: filePath}
}

type storage struct {
  filePath string
}

func (s storage) Read(data interface{}) error {
  file, err := os.ReadFile(s.filePath)
  if err != nil {
    return fmt.Errorf("error: al leer el archivo %v\n", err)
  }

  return json.Unmarshal(file, &data)
}

func (s storage) Write(data interface{}) error {
  fileData, err := json.MarshalIndent(data,"","")
  if err != nil {
    return fmt.Errorf("error: marshalling %v", err)
  }

  return os.WriteFile(s.filePath, fileData, 0644)
}
