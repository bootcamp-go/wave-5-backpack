package main

import (
	"os"

	"github.com/bootcamp-go/cmd/server/handler"
	"github.com/bootcamp-go/docs"
	"github.com/bootcamp-go/internal/transactions"
	"github.com/bootcamp-go/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Bootcamp Go
// @version 1.0
// @descripcion This API handle MELI Transactions.
// @termsOfService https://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones

// @contact.name API Support
// @contact.url https://developers.mercadolibre.com.ar/support

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	loadEnv()

	router := gin.Default()
	db := store.NewFileStore(store.FileType, "./transactions.json")
	r := transactions.NewRepository(db)
	s := transactions.NewService(r)
	h := handler.NewTransaction(s)

	// Documentación
	docs.SwaggerInfo.Host = os.Getenv("HOST")
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	rTransaction := router.Group("transactions")
	rTransaction.GET("/", h.GetAll())
	rTransaction.POST("/", h.Store())
	rTransaction.PUT("/:id", h.Update())
	rTransaction.DELETE("/:id", h.Delete())
	rTransaction.PATCH("/:id", h.UpdateReceptorYMonto())

	err := router.Run(":" + os.Getenv("HOST"))
	if err != nil {
		panic(err)
	}
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}
