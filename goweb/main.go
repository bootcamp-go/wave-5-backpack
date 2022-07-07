package main

import (
	"encoding/json"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

type usuarios struct {
	Id       int       `json:"id" binding:"required"`
	Nombre   string    `json:"nombre"`
	Apellido string    `json:"apellido"`
	Email    string    `json:"email"`
	Edad     int       `json:"edad"`
	Altura   float64   `json:"altura"`
	Activo   bool      `json:"activo"`
	Fecha    time.Time `json:"fecha"`
}

func main() {
	//router con logger y recover
	router := gin.Default()

	router.GET("/hola", func(c *gin.Context) {
		//respuesta tipo JSON
		c.JSON(200, gin.H{
			"message": "Hola Luz! :D",
		})
	})

	router.GET("/usuarios", getAll)
	// Corriendo servidor sobre el puerto 8080
	router.Run(":8081")
}

func getAll(c *gin.Context) {
	//Leo el json y lo envío como retorno
	//para mostrarlo en el path /usuarios
	jsonUsers, err := os.ReadFile("./users.json")
	if err != nil {
		panic(err)
	}
	var users []usuarios
	err = json.Unmarshal(jsonUsers, &users)
	if err != nil {
		panic(err)
	}
	c.JSON(200, gin.H{"usuarios": users})
}
