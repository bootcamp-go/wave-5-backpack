package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"strconv"
)

type users struct {
	Id           int     `json:"id"`
	FirstName    string  `json:"firstName"`
	LastName     string  `json:"lastName"`
	Email        string  `json:"email"`
	Age          int     `json:"age"`
	Height       float64 `json:"height"`
	Active       bool    `json:"active"`
	CreationDate string  `json:"creationDate"`
}

func GetAll(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hola")
}

func UsersHandler(c *gin.Context) {
	var data, _ = os.ReadFile("./users.JSON")
	var u []users

	if err := json.Unmarshal(data, &u); err != nil {
		fmt.Println("Error")
		log.Fatal(err)
	}
	c.JSON(200, gin.H{
		"Saludo": "hola estas pasando por el userHandler",
		"data":   u,
	})
}
func GetUsersByIdHandler(c *gin.Context) {
	var data, _ = os.ReadFile("./users.JSON")
	var u []users

	if err := json.Unmarshal(data, &u); err != nil {
		fmt.Println("Error")
		log.Fatal(err)
	}
	for key, _ := range u {
		if strconv.Itoa(u[key].Id) == c.Param("id") {
			c.JSON(200, gin.H{
				"saludo": "hola estas pasando por el GetuserByIdHandler",
				"data":   u[key],
			})
		}
	}
}
