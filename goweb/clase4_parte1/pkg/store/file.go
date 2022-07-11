package store

import (
	"encoding/json"
	"os"
)

/* Se implementa la interfaz Store con los métodos Read y Write,
ambos métodos reciben una interfaz y devolverán un error */
type Store interface {
	Read(data interface{}) error
	Write(data interface{}) error
	Ping() error
}

/* Se implementa la función NewStore encargada de generar la estructura
fileStore y recibe el parámetro fileName que es el nombre del archivo */
func NewStore(fileName string) Store {
	return &fileStore{fileName}
}

// Se declara la estructura fileStore con el campo filePath que guardará el nombre del archivo
type fileStore struct {
	FilePath string
}

/* Se implementa el método Write para escribir datos de la estructura en el archivo.
Recibe una interfaz y lo convertirá a una representación JSON en bytes para guardarlo
en el archivo que especificamos al momento de instanciar la función NewStore. */
func (fs *fileStore) Write(data interface{}) error {
	fileData, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return err
	}
	// "0644" permiso de sólo lectura, y escritura para propietario
	return os.WriteFile(fs.FilePath, fileData, 0644)
}

/* Se implementa el método Read para leer el archivo y guardar su contenido
empleando la interfaz que recibirá como parámetro */
func (fs *fileStore) Read(data interface{}) error {
	file, err := os.ReadFile(fs.FilePath)
	if err != nil {
		return err
	}
	return json.Unmarshal(file, &data)
}

/* Se implementa el método Ping para validar que el archivo que se cargó
existe en el directorio */
func (fs *fileStore) Ping() error {
	if _, err := os.ReadFile(fs.FilePath); err != nil {
		return err
	}
	return nil
}