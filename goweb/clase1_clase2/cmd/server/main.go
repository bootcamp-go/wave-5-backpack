package main

import (
	"goweb/clase1_clase2/cmd/server/handler"
	"goweb/clase1_clase2/docs"
	"goweb/clase1_clase2/internal/products"
	"goweb/clase1_clase2/pkg/store"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title MELI Bootcamp API | Jessica Escobar
// @version 1.0
// @description This API implements the CRUD method for MELI Products
// @termsOfService https://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones

// @contact.name API Support
// @contact.url https://developers.mercadolibre.com.ar/support

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error al intentar cargar el archivo .env")
	}

	db := store.NewStore("products.json")
	repo := products.NewRepository(db)
	service := products.NewService(repo)
	p := handler.NewProduct(service)

	router := gin.Default()
	pr := router.Group("/products")

	// Swagger Documentation - Endpoint
	docs.SwaggerInfo.Host = os.Getenv("HOST")
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// <--------------------------------------------------------------->
	pr.Use(p.TokenAuthMiddleware())
	pr.POST("/", p.Store())
	pr.GET("/", p.GetAll())
	pr.GET("/:id", p.GetById())
	pr.PUT("/:id", p.Update())
	pr.PATCH("/:id", p.UpdateFields())
	pr.DELETE("/:id", p.Delete())
	err = router.Run()
	if err != nil {
		panic(err)
	}
}
