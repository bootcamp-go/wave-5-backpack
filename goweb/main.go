package main

import (
	"fmt"
	"goweb/models"
	"goweb/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ejercicio 1 parte 1
func GetAll(ctx *gin.Context) {
	products, err := services.ReadAllProducts()
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// Ejercicio 1 parte 2
	var filtrados []models.Product
	for _, product := range products {

		var idVerif, nombreVerif, colorVerif, precioVerif, stockVerif, codigoVerif, publicadoVerif, fechaCreacionVerif bool

		idVerif = ctx.Query("id") != "" && ctx.Query("id") != fmt.Sprintf("%d", product.Id)
		nombreVerif = ctx.Query("nombre") != "" && ctx.Query("nombre") != product.Nombre
		colorVerif = ctx.Query("color") != "" && ctx.Query("color") != product.Color
		precioVerif = ctx.Query("precio") != "" && ctx.Query("precio") != fmt.Sprintf("%f", product.Precio)
		stockVerif = ctx.Query("stock") != "" && ctx.Query("stock") != fmt.Sprintf("%d", product.Stock)
		codigoVerif = ctx.Query("codigo") != "" && ctx.Query("codigo") != product.Codigo
		publicadoVerif = ctx.Query("publicado") != "" && ctx.Query("publicado") != fmt.Sprintf("%t", product.Publicado)
		fechaCreacionVerif = ctx.Query("fechaCreacion") != "" && ctx.Query("fechaCreacion") != product.FechaCreacion

		if idVerif || nombreVerif || colorVerif || precioVerif || stockVerif || codigoVerif || publicadoVerif || fechaCreacionVerif {
			continue
		}

		filtrados = append(filtrados, product)

	}

	ctx.JSON(200, filtrados)
}

// Ejercicio 2 parte 2
func GetById(ctx *gin.Context) {
	producto, err := services.GetById(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, producto)
}

func main() {
	router := gin.Default()

	// Ejercicio 1 parte 1
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "Hola " + "Juan Pablo Ortiz"})
	})

	productos := router.Group("/productos")
	{
		// Ejercicio 2 parte 1
		productos.GET("/", GetAll)
		// Ejercicio 2 parte 2
		productos.GET("/:id", GetById)
	}

	router.Run(":8080")
}
