package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type transactions struct {
	Id       int
	Codigo   string
	Moneda   string
	Monto    float64
	Emisor   string
	Receptor string
	Fecha    string
}

var t []transactions

func main() {

	file, err := os.ReadFile("./transactions.json")
	if err != nil {
		panic(err)
	}

	if e := json.Unmarshal(file, &t); e != nil {
		log.Fatal(e)
	}

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hola Simón",
		})
	})

	router.GET("/transactions", GetAll)

	router.Run()
}

func GetAll(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, t)
}
