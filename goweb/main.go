package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	type transacciones struct {
		Id                 int     `json:"id"`
		Codigo_transaccion string  `json:"codigo_transaccion"`
		Moneda             string  `json:"moneda"`
		Monto              float64 `json:"monto"`
		Emisor             string  `json:"emisor"`
		Receptor           string  `json:"receptor"`
		Fecha_transaccion  string  `json:"fecha_transaccion"`
	}

	data, _ := os.ReadFile("./transacciones.json")

	var t []transacciones

	if err := json.Unmarshal(data, &t); err != nil {
		fmt.Println("error aqui")
		log.Fatal(err)
	}

	router := gin.Default()

	router.GET("/saludo", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hola Sala 2 saludo",
		})
	})

	router.GET("/transacciones", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"data": t,
		})
	})

	router.Run()
}
