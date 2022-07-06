package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Usuarios struct {
	Id       int     `json:"id"`
	Nombre   string  `json:"nombre"`
	Apellido string  `json:"apellido"`
	Email    string  `json:"email"`
	Edad     int     `json:"edad"`
	Altura   float64 `json:"altura"`
	Activo   bool    `json:"activo"`
	Creacion string  `json:"creacion"`
}

func main() {
	router := gin.Default()

	// Saludo estandar
	router.GET("/greet", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hola mundo!",
		})
	})

	// Saludo por nombre
	router.GET("/greet/:nombre", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Hola %s!", c.Param("nombre")),
		})
	})

	// Listado de usuarios
	path := "1/tm/usuarios.json"

	router.GET("/usuarios", func(c *gin.Context) {
		jsonData, err := os.ReadFile(path)
		if err != nil {
			panic(err)
		}

		var u []Usuarios
		err = json.Unmarshal(jsonData, &u)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, gin.H{
			"usuarios": u,
		})
	})

	router.Run()
}
