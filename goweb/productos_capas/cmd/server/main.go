package main

import (
	"goweb/productos_capas/cmd/server/handler"
	"goweb/productos_capas/docs"
	"goweb/productos_capas/internal/products"
	"goweb/productos_capas/pkg/store"
	"log"
	"os"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// @title MELI Camilo Calderón API
// @version 1.0
// @description Product management CRUD
// @termsOfService https://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones

// @contact.name API Support
// @contact.url https://developers.mercadolibre.com.ar/support

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	err := godotenv.Load("cmd/server/.env")
	if err != nil {
		log.Fatal("error al intentar cargar el archivo .env")
	}

	db := store.NewStore("products.json")
	if err := db.Ping(); err != nil {
		log.Fatal("error al intentar cargar archivo")
	}

	repo := products.NewRepository(db)
	service := products.NewService(repo)
	p := handler.NewProduct(service)
	router := gin.Default()

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	pr := router.Group("/productos")
	pr.Use(p.TokenAuthMiddleware())
	pr.GET("/", p.GetAll())
	pr.GET("/:id", p.GetByID())
	pr.POST("/", p.Store())
	pr.PUT("/:id", p.Update())
	pr.PATCH("/:id", p.UpdateNamePrice())
	pr.DELETE("/:id", p.Delete())
	router.Run()
}
