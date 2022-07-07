package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*Ejercicio 3 - Listar Entidad

Ya habiendo creado y probado nuestra API que nos saluda, generamos una ruta que
devuelve un listado de la temática elegida.
Dentro del “main.go”, crea una estructura según la temática con los campos correspondientes.
	1. Genera un endpoint cuya ruta sea /temática (en plural). Ejemplo: “/productos”
	2. Genera un handler para el endpoint llamado “GetAll”.
	3. Crea una slice de la estructura, luego devuelvela a través de nuestro endpoint.
*/

type Transaction struct {
	Id                int64
	CodigoTransaccion string
	Moneda            string
	Monto             float64
	Emisor            string
	Receptor          string
}

func main() {
	router := gin.Default()

	router.GET("/nombre", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hola 👋 Bootcampers",
		})
	})

	router.GET("/transacciones", GetAll)

	router.Run(":8000") // Por defecto gin arranca el server on 8080
}

func GetAll(ctx *gin.Context) {
	transactions := []Transaction{
		{Id: 1, CodigoTransaccion: "abc123", Moneda: "peso", Monto: 100, Emisor: "Martín", Receptor: "Luisa"},
		{Id: 2, CodigoTransaccion: "abc134", Moneda: "dolar", Monto: 200, Emisor: "Marcos", Receptor: "Luisa"},
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": transactions,
	})
}
