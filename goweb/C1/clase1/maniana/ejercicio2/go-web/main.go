package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*Ejercicio 2 - Hola {nombre}

1. Crea dentro de la carpeta go-web un archivo llamado main.go
2. Crea un servidor web con Gin que te responda un JSON que tenga una clave “message”
y diga Hola seguido por tu nombre.
3. Pegale al endpoint para corroborar que la respuesta sea la correcta.
*/

func main() {
	router := gin.Default()

	router.GET("/nombre", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hola 👋 Bootcampers",
		})
	})

	router.Run(":8000") // Por defecto gin arranca el server on 8080
}
