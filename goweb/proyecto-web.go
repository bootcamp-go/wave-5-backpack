package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Transaccion struct {
	Id                int     `json:"Id"`
	CodigoTransaccion string  `json:"codigo_transaccion"`
	Moneda            string  `json:"moneda"`
	Monto             float64 `json:"monto"`
	Emisor            string  `json:"emisor"`
	Receptor          string  `json:"receptor"`
	FechaTransaccion  string  `json:"fecha_transaccion"`
}

var transaciones []Transaccion = []Transaccion{
	{
		Id:                1,
		CodigoTransaccion: "A-1",
		Moneda:            "$",
		Monto:             50,
		Emisor:            "ARCOR",
		Receptor:          "BCRA",
		FechaTransaccion:  "12-05-2022",
	},
	{
		Id:                2,
		CodigoTransaccion: "B-1",
		Moneda:            "USD",
		Monto:             5000,
		Emisor:            "TOYOTA",
		Receptor:          "BCRA",
		FechaTransaccion:  "10-02-2022",
	},
	{
		Id:                3,
		CodigoTransaccion: "B-2",
		Moneda:            "EUR",
		Monto:             500,
		Emisor:            "FORD",
		Receptor:          "BCRA",
		FechaTransaccion:  "01-03-2022",
	},
}

func main() {
	router := gin.Default()
	nombre := "Charly"

	router.GET("transacciones", getAll)

	router.GET("hola", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hola " + nombre,
		})
	})
	router.Run()
}

func getAll(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, transaciones)
}
