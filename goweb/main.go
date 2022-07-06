package main

import (
	"goweb/services"

	"github.com/gin-gonic/gin"
)

func GetAll(ctx *gin.Context) {
	products, err := services.ReadAllProducts()
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, products)
}

func main() {
	router := gin.Default()

	// Ejercicio 1
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "Hola " + "Juan Pablo Ortiz"})
	})

	// Ejercicio 2
	router.GET("/productos", GetAll)

	router.Run(":8080")
}
