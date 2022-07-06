package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	// Definición estructura de json

	type transacciones struct {
		Id                 int     `json:"id"`
		Codigo_transaccion string  `json:"codigo_transaccion"`
		Moneda             string  `json:"moneda"`
		Monto              float64 `json:"monto"`
		Emisor             string  `json:"emisor"`
		Receptor           string  `json:"receptor"`
		Fecha_transaccion  string  `json:"fecha_transaccion"`
	}

	// Se obtiene la data del archivo creado

	data, _ := os.ReadFile("./transacciones.json")

	// Se define variable [] de transacciones para transformar []byte data a formato transacciones struc
	var t []transacciones

	if err := json.Unmarshal(data, &t); err != nil {
		fmt.Println("error aqui")
		log.Fatal(err)
	}

	// Se crea el routuer y los end points

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

	// Se crea
	router.Run()
}
