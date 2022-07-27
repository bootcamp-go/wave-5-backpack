package main

import (
	"WebServer/docs"
	"WebServer/internal/transactions"
	"WebServer/pkg/store"
	"WebServer/server/handler"
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
	_ = godotenv.Load()

	// it will be visible for this scope.
	db := store.NewStore("Transactions.json")
	if err := db.Ping(); err != nil {
		log.Fatal("Cannot read file")
	}

	repo := transactions.NewRepository(db)
	service := transactions.NewService(repo)
	t := handler.NewTransaction(service)

	r := gin.Default()

	// Swagger Documentation - Endpoint
	docs.SwaggerInfo.Host = os.Getenv("HOST")
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// <--------------------------------------------------------------->
	tr := r.Group("/transactions")
	tr.POST("/", t.Create())
	tr.GET("/", t.GetAll())
	tr.PUT("/:id", t.Update())
	tr.PATCH("/:id", t.UpdatePartial())
	tr.DELETE("/:id", t.Delete())
	r.Run()
}
