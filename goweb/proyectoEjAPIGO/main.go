package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

var id int = 3

func data() []Transaccion {
	data, err := os.ReadFile("./transacciones.json")

	if err != nil {
		fmt.Println("Error en la lectura: %v", err)
	}

	var t []Transaccion
	if err := json.Unmarshal(data, &t); err != nil {
		log.Fatal(err)
	}
	return t

}

var t []Transaccion = data()

func HandlerRaiz(ctx *gin.Context) {

	ctx.JSON(200, gin.H{
		"message": "Hola!!, bienvenidos a mi api :)",
	})

}

func AgregarTransacciones() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req Transaccion

		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(404, gin.H{

				"error": err.Error(),
			})
			return
		}
		req.Id = id + 1
		id = req.Id
		t = append(t, req)
		ctx.JSON(200, req)
	}
}

func HandlerGetID(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	var aux []Transaccion
	for _, transaccion := range t {
		if transaccion.Id == id {
			aux = append(aux, transaccion)
			break
		}
	}
	if aux != nil {
		ctx.JSON(200, gin.H{

			"transaccion": aux,
		})
	} else {
		ctx.JSON(200, gin.H{

			"error": "No se a encontrado una transaccion con este id.",
		})
	}
}
func HandlerFiltradoMoneda(ctx *gin.Context) {

	moneda := ctx.Query("moneda")

	if moneda != "" {
		var aux []Transaccion
		for _, transaccion := range t {
			if transaccion.Moneda == moneda {
				aux = append(aux, transaccion)
			}
		}
		if aux == nil {
			ctx.JSON(200, gin.H{

				"error": "No se a encontrado una transaccion con este filtro.",
			})
		} else {
			ctx.JSON(200, gin.H{

				"transacciones": aux,
			})
		}
	} else {
		ctx.JSON(200, gin.H{

			"transacciones": t,
		})
	}
}

func main() {

	// var transacciones []Transaccion

	server := gin.Default()
	server.GET("/", HandlerRaiz)

	transacciones := server.Group("/transacciones")
	{

		transacciones.GET("/", HandlerFiltradoMoneda)
		transacciones.GET("/:id", HandlerGetID)
		transacciones.POST("/", AgregarTransacciones())
	}

	server.Run(":8005")
}
