package main

import (
	"goweb/cmd/server/handler"
	"goweb/docs"
	"goweb/internal/products"
	"goweb/pkg/store"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title MELI Bootcamp API
// @version 1.0
// @description This API Handle MELI Products.
// @termsOfService https://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones

// @contact.name API Support
// @contact.url https://developers.mercadolibre.com.ar/support

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	// Cargar las variables de entorno
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("error al intentar cargar archivo .env")
	}

	// Cargar la base de datos
	db := store.NewStore(os.Getenv("DB_PATH"))
	if err := db.Ping(); err != nil {
		log.Fatal("error al conectar con la base de datos")
	}

	//repository := products.NewRepositoryMemoria()
	//repository := products.NewRepositoryJsonDB()
	repository := products.NewRepository(db)
	service := products.NewService(repository)
	p := handler.NewProduct(service)

	router := gin.Default()

	// Cargar la API de documentación
	docs.SwaggerInfo.Host = os.Getenv("HOST")
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	productos := router.Group("/products")
	{
		productos.GET("/", p.GetAll())
		productos.GET("/:id", p.GetById())
		productos.POST("/", p.Store())
		productos.PUT("/:id", p.Update())
		productos.DELETE("/:id", p.Delete())
		productos.PATCH("/:id", p.UpdateNombreYPrecio())
	}

	// Clase 1 Ejercicio 1 Parte 1
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "Hola " + "Juan Pablo Ortiz"})
	})

	err = router.Run(os.Getenv("PORT"))
	if err != nil {
		log.Fatal("error al iniciar el servidor")
	}
}
