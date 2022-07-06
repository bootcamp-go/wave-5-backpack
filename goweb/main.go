package main

import (
	"encoding/json"
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

// Estructura de usuarios
type Usuarios struct {
	Id            int     `json:"id"`
	Nombre        string  `json:"nombre"`
	Apellido      string  `json:"apellido"`
	Email         string  `json:"email"`
	Edad          int     `json:"edad"`
	Altura        float64 `json:"altura"`
	Activo        bool    `json:"activo"`
	FechaCreacion string  `json:"fecha_creacion"`
}

func GetAll(c *gin.Context) {
	// Abrimos el archivo
	jsonFile, err := os.Open("./usuarios.json")
	if err != nil {
		log.Println("Error opening json file: ", err)
	}
	// Cerramos el archivo
	defer jsonFile.Close()

	// Creamos el objeto decodificador
	decoder := json.NewDecoder(jsonFile)

	// Creamos el Slice de usuarios
	var user []Usuarios

	// Decodificamos el archivo y se asigna al slice
	err = decoder.Decode(&user)
	if err == io.EOF {
		log.Println("EOF: ", err)
	}
	if err != nil {
		log.Println("Error decoding json data: ", err)
	}

	// Devolvemos el JSON al cliente
	c.JSON(200, user)
}

func main() {
	// Crea un router con gin
	router := gin.Default()

	// Captura la solicitud GET "/hello"
	router.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hola Arturo"})
	})

	// Devuelve los elementos de los usuarios.json
	router.GET("/usuarios", GetAll)

	// Corremos nuestro servidor sobre el puerto 8080
	router.Run()
}
