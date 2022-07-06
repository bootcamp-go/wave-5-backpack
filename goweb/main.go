package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

type Transaccion struct {
	Id                 int
	Codigo_transaccion string
	Moneda             string
	Monto              float64
	Emisor             string
	Receptor           string
	Fecha_transaccion  string
}

func main() {

	// var transacciones []Transaccion

	data, err := os.ReadFile("./transacciones.json")

	if err != nil {
		fmt.Println("Error en la lectura: %v", err)
	}
	// jsonData := string(data)
	var t []Transaccion
	if err := json.Unmarshal(data, &t); err != nil {
		log.Fatal(err)
	}

	router := gin.Default()
	router.GET("/hello-world", func(c *gin.Context) {

		c.JSON(200, gin.H{
			"message": "Hola!!",
		})
	})

	router.GET("/transacciones", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"users": t,
		})
	})
	router.Run()
}
