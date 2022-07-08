package main

import (
	"github.com/bootcamp-go/wave-5-backpack/tree/Ramos_Andres/goweb/practica/cmd/sever/handler"
	"github.com/bootcamp-go/wave-5-backpack/tree/Ramos_Andres/goweb/practica/internal/products"

	"github.com/gin-gonic/gin"
)

func main() {
	repository := products.NewRepository()
	service := products.NewService(repository)
	p := handler.NewProduct(service)

	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		name := ctx.Request.URL.Query().Get("name")
		if name == "" {
			name = "Anonimo"
		}
		ctx.JSON(200, gin.H{
			"message": "Saludos " + name,
		})

	})

	productos := router.Group("/products")
	{
		productos.GET("/", p.GetAll())
		productos.GET("/:id", p.GetById())
		productos.POST("/", p.Store())
		productos.PUT("/:id", p.UpdateTotal())
		productos.PATCH("/:id", p.UpdatePartial())
		productos.DELETE("/:id", p.Delete())
	}

	router.Run(":8080")
}
