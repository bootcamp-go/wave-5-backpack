package main

import (
	"github.com/bootcamp-go/wave-5-backpack/tree/Ramos_Andres/goweb/practica/cmd/server/handler"
	"github.com/bootcamp-go/wave-5-backpack/tree/Ramos_Andres/goweb/practica/internal/products"
	"github.com/bootcamp-go/wave-5-backpack/tree/Ramos_Andres/goweb/practica/pkg/file"
	"github.com/ncostamagna/meli-bootcamp/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"os"

	"github.com/gin-gonic/gin"
)

// @title Bootcamp MELI API
// @version 1.0
// @description esta api maneja los productos de nuestro proyecto
// @termsOfService N/A

// @contact.name API Support
// @contact.url www.com

// @license.name Apache 2.0
// @license.url www.com
func main() {
	fileDB := file.NewFile("/Users/pmelegatti/Documents/wave-5-backpack/goweb/resources/products.json")
	if err := fileDB.Ping(); err != nil {
		panic(err)
	}
	repository := products.NewRepository(fileDB)
	service := products.NewService(repository)
	p := handler.NewProduct(service)

	router := gin.Default()

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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
