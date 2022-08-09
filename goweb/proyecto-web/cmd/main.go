package main

import (
	"log"
	"os"
	"proyecto-web/cmd/handlers"
	"proyecto-web/internal/transaction"
	"proyecto-web/internal/transaction/bdRepository"
	"proyecto-web/pkg/store"

	"github.com/joho/godotenv"

	"proyecto-web/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

// @title MELI Botocamp API
// @version 1.0
// @description Esta API es para manejar transacciones
// @termsOfService https://www.google.com
// @contact.name API Support
// @contact.url https://github.com/cgdesiderio96
// @license.name Gin
// @license.url https://github.com/gin-gonic/gin

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error al intentar cargar archivo .env")
	}
	bdPosta := store.NewDb()
	//bdFileSystem := store.NewStore("transacciones.json")

	//r := transaction.NewRepository(bdFileSystem)
	r := bdRepository.NewBdRepository(bdPosta)

	service := transaction.NewService(r)
	handler := handlers.NewTransactionHandler(service)
	servidor := gin.Default()

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	servidor.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	gr := servidor.Group("transacciones")
	{
		gr.GET("/", handler.GetAll())
		gr.GET("/:id", handler.GetById())
		gr.GET("/filters", handler.GetByCodigoTransaccion())
		gr.POST("/", handler.Create())
		gr.PUT("/:id", handler.Update())
		gr.PATCH("/:id", handler.UpdateParcial())
		gr.DELETE("/:id", handler.Delete())
	}

	err = servidor.Run()
	if err != nil {
		panic("cannot create server")
	}

}
