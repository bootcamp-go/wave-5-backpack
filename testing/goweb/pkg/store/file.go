package store

import (
	"encoding/json"
	"os"
)

// 1) Se crea la interface Store con los métodos Read y Write. Ambos métodos reciben una interfaz vacía y devuelven un error
type Store interface {
	Read(data interface{}) error
	Write(data interface{}) error

	// este método lo agrego para validar que, cuando inicio la aplicación, el archivo se lea correctamente. Entonces termino haciendo un Open.
	Ping() error
}

// 2) Factory de Store: Se debe implementar la función Factory que se encarga de generar la estructura que deseamos y recibe el tipo de store que queremos implementar y el nombre del archivo. Se declara la estructura FileStore con el campo que guarde el nombre del archivo.
// Básicamente esta función nos "devuelve la implementación de Store"
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

// Metodo Write: Toma la info que le mandamos y la escribe en el archivo json. Recibe una interfaz y lo va a convertir a una representación de json en bytes para poder escribir el archivo
func (fs *fileStore) Write(data interface{}) error {
	fileData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return os.WriteFile(fs.FilePath, fileData, 0644)
}


// Read: toma la dirección del contenido y con el unMarshall lo transforma en algo que podamos usar desde la aplicación
func (fs *fileStore) Read(data interface{}) error {
	file, err := os.ReadFile(fs.FilePath)
	if err != nil {
		return err
	}
	return json.Unmarshal(file, &data)
}




