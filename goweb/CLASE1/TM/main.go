package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

type users struct {
	Id       int       `json:"id"`
	Nombre   string    `json:"nombre" binding:"required"`
	Apellido string    `json:"apellido" binding:"required"`
	Email    string    `json:"email" binding:"required"`
	Edad     int       `json:"edad" binding:"required"`
	Altura   float64   `json:"altura" binding:"required"`
	Activo   time.Time `json:"activo" binding:"required"`
}

var usuarios []users
var U = users{}

func main() {
	router := gin.Default()

	//Ejercicio 2
	router.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hola Luz",
		})
	})
	//

	//Ejercicio 3
	router.GET("/users", GetAll)
	router.Run()

}

func GetAll(c *gin.Context) {

	data, err := os.ReadFile("./users.json")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Se abrio el archivo.")
	if err = json.Unmarshal([]byte(data), &U); err != nil {
		log.Fatal(err)
	}
	c.JSON(200, U)
}

///
