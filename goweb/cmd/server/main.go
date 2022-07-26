package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"goweb/cmd/server/docs"
	"goweb/cmd/server/handler"
	"goweb/internal/products"
	"goweb/pkg/file"
	"os"
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
