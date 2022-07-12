package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/bootcamp-go/wave-5-backpack/tree/lopez_cristian/goweb/docs"
	"github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/bootcamp-go/wave-5-backpack/tree/lopez_cristian/goweb/cmd/handler"
	"github.com/bootcamp-go/wave-5-backpack/tree/lopez_cristian/goweb/internal/transactions"
	"github.com/bootcamp-go/wave-5-backpack/tree/lopez_cristian/goweb/pkg/storage"
)

// @title MELI Bootcamp-go practice API
// @version 1.0
// @description This API is from Bootcamp-go

// @license Apache 2.0
// @license.url https://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error al intentar leer el archivo .env")
	}

	file, err := os.Open("transactions.json")
	defer file.Close()

	storage := storage.NewStorage("transactions.json")

	repo := transactions.NewRepository(storage)
	service := transactions.NewService(repo)
	t := handler.NewTransaction(service)

	router := gin.Default()

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	rt := router.Group("/transactions")
	{
		rt.GET("", t.GetAll)
		rt.GET("/:id", t.GetByID)

		rt.PUT("/:id", t.Update)
		rt.PATCH("/:id", t.Patch)

		rt.POST("", t.CreateTransaction)

		rt.DELETE("/:id", t.Delete)
	}

	router.Run()
}
