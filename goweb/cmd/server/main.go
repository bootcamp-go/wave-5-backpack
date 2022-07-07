package main

import (
	"goweb/cmd/server/handler"
	"goweb/internal/products"

	"github.com/gin-gonic/gin"
)

func main() {
	//repository := products.NewRepository()
	repositoryJson := products.NewRepositoryJsonDB()
	service := products.NewService(repositoryJson)
	p := handler.NewProduct(service)

	router := gin.Default()

	// Clase 1 Ejercicio 1 Parte 1
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "Hola " + "Juan Pablo Ortiz"})
	})

	productos := router.Group("/productos")
	{
		// Clase 1 Ejercicio 2 Parte 1
		productos.GET("/", p.GetAll())
		// Clase 1 Ejercicio 2 Parte 2
		productos.GET("/:id", p.GetById())
		// Clase 2 Ejercicio 1 Parte 1
		productos.POST("/", p.Store())
	}

	router.Run(":8080")
}
