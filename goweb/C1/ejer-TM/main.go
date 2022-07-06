package main

// Crea dentro de la carpeta go-web un archivo llamado main.go
// Crea un servidor web con Gin que te responda un JSON que tenga una clave “message” y diga Hola seguido por tu nombre.
// Pegale al endpoint para corroborar que la respuesta sea la correcta.

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	const name string = "Seba"

	type transacciones struct {
		Id          int     `json:"id"`
		TranCode    string  `json:"tranCode"`
		Currency    string  `json:"currency"`
		Ammount     float64 `json:"ammount"`
		Transmitter string  `json:"transmitter"`
		Reciever    string  `json:"reciever"`
		TranDate    string  `json:"tranDate"`
	}

	// Se obtiene la data del archivo creado

	data, _ := os.ReadFile("./transacciones.json")

	// Se define variable [] de transacciones para transformar []byte data a formato transacciones struc
	var t []transacciones

	if err := json.Unmarshal(data, &t); err != nil {
		fmt.Println("error aqui")
		log.Fatal(err)
	}

	// Se crea el router y los endpoints

	router := gin.Default()

	router.GET("/saludo", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hola " + name,
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
