package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/gin-gonic/gin"
)

type User struct {
	Id           int     `json:"id"`
	Name         string  `json:"name"`
	LastName     string  `json:"lastname"`
	Email        string  `json:"email"`
	Age          int     `json:"age"`
	Height       float64 `json:"height"`
	Active       bool    `json:"active"`
	CreationDate string  `json:"creation-date"`
}

func main() {
	name := "Cristian Ladino"
	var users []User
	jsonData, err := ioutil.ReadFile("users.json")
	if err != nil {
		log.Fatal(err)
	}
	// Crea un router con gin
	router := gin.Default()

	// EJERCICIO 2
	router.GET("/hello-world", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hola " + name})
	})

	// EJERCICIO 3
	if err := json.Unmarshal([]byte(jsonData), &users); err != nil {
		log.Fatal(err)
	}

	router.GET("/users", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": &users})
	})

	router.Run(":8080")

	fmt.Println(users)

}
