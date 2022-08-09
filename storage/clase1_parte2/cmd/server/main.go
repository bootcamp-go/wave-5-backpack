package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"goweb/docs"
	"goweb/pkg/db"

	"goweb/cmd/server/handler"
	"goweb/internal/transactions"
)

// @title MELI Bootcamp API
// @version 1.0
// @description This API handle MELI Transactions
// termsOfService https://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones

//@contact.name API Support
//@contact.url https://developers.mercadolibre.com.ar/support

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error al cargar el archivo .env")
	}

	StorageDB := db.MySQLConnection()
	repositoryTransaction := transactions.NewRepository(StorageDB)
	serviceTransaction := transactions.NewService(repositoryTransaction)
	handlerTransaction := handler.NewHandler(serviceTransaction)

	r := gin.Default()

	//Implementacion swagger
	docs.SwaggerInfo.Host = os.Getenv("HOST")
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	tg := r.Group("/transactions")
	tg.DELETE("/:id", handlerTransaction.Delete())
	tg.GET("/", handlerTransaction.GetAll())
	tg.GET("/:sender", handlerTransaction.GetBySender())
	tg.POST("/", handlerTransaction.Store())
	tg.PUT("/:id", handlerTransaction.Upddate())
	tg.PATCH("/:id", handlerTransaction.UpdateAmount())
	r.Run()
}