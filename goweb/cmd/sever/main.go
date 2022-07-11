package main

import (
	"fmt"

	"github.com/bootcamp-go/wave-5-backpack/tree/Ramos_Andres/goweb/practica/cmd/sever/handler"
	"github.com/bootcamp-go/wave-5-backpack/tree/Ramos_Andres/goweb/practica/docs"
	"github.com/bootcamp-go/wave-5-backpack/tree/Ramos_Andres/goweb/practica/internal/products"
	"github.com/bootcamp-go/wave-5-backpack/tree/Ramos_Andres/goweb/practica/pkg/file"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

func main() {
	fileDB := file.NewFile("/Users/andreramos/Documents/Bootcamp-Go Meli/wave-5-backpack/goweb/resources/products.json")
	if err := fileDB.Ping(); err != nil {
		panic(err)
	} else {
		fmt.Println("good to go")
	}
	repository := products.NewRepository(fileDB)
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

	docs.SwaggerInfo.Host = "localhost:8080"
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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
