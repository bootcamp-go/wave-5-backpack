package main

import (
	"goweb/cmd/server/handler"
	"goweb/docs"
	"goweb/internal/transactions"
	"goweb/pkg/store"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//@title Go Web API
//@version 1.0
//@description This API is about Transactions
//@temsOfService https://example.com

//@contact.name API Support
//@contac.url https://developers.mercadolibre.com.cl/support

//@license.name FRP
//@license.url https://example.com
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error al intentar cargar archivo .env")
	}

	db := store.NewStore("transactions.json")

	if err := db.Ping(); err != nil {
		log.Fatal("error al intentar cargar el archivo")
	}

	repository := transactions.NewRepository(db)
	service := transactions.NewService(repository)
	transactions := handler.NewTransaction(service)

	router := gin.Default()

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	router.GET("docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r := router.Group("/transactions")
	r.POST("/", transactions.Store())
	r.GET("/", transactions.GetAll())
	r.PUT("/:id", transactions.Update())
	r.DELETE("/:id", transactions.Delete())
	r.PATCH("/:id", transactions.UpdateCodeAmount())
	router.Run()
}
