package main

import (
	"encoding/json"
	"os"

	"github.com/gin-gonic/gin"
)

type Transaction struct {
	Id       string `json:"id"`
	Code     string `json:"code"`
	Moneda   string `json:"moneda"`
	Monto    int    `json:"monto"`
	Emisor   string `json:"emisor"`
	Receptor string `json:"receptor"`
	Fecha    string `json:"fecha"`
}

func getAll(c *gin.Context) {
	var transactions []Transaction
	file, err := os.ReadFile("./transactions.json")
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
	}
	err2 := json.Unmarshal(file, &transactions)
	if err2 != nil {
		c.JSON(500, gin.H{
			"error": err2.Error(),
		})
	}
	c.JSON(200, gin.H{
		"transacciones": transactions,
	})
}

func main() {
	router := gin.Default()

	router.GET("/hola", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hola Cristobal",
		})
	})

	router.GET("/transacciones", getAll)

	router.Run()
}
