package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	type transaccion struct {
		id                int
		codigoTransaccion string
		moneda            string
		monto             int
		emisor            string
		receptor          string
		fecha             string
	}

	// Crea un router con gin
	router := gin.Default()

	// Captura la solicitud GET “/hello-world”
	router.GET("/transacciones", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hola Sebastián!"})
	})

	//var lista [] transaccion

	// Corremos nuestro servidor sobre el puerto 8080
	router.Run(":8080")
	fmt.Print("")
}
