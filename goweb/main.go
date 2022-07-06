package main

import (
	"encoding/json"
	"os"

	"github.com/gin-gonic/gin"
)

type transaction struct {
	Id        int
	Code      string
	Currency  string
	Amount    float64
	Issuer    string
	Recipient string
	Date      string
}

func hello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hola",
	})
}

func GetAll(c *gin.Context) {
	file, _ := os.ReadFile("transactions.json")
	var transactions []transaction
	json.Unmarshal([]byte(file), &transactions)
	c.JSON(200, transactions)
}

func main() {
	router := gin.Default()
	router.GET("/hello", hello)
	router.GET("/transactions", GetAll)
	router.Run()
}
