package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/bootcamp-go/wave-5-backpack/tree/lopez_cristian/storage/cmd/db"
	"github.com/bootcamp-go/wave-5-backpack/tree/lopez_cristian/storage/cmd/handler"
	"github.com/bootcamp-go/wave-5-backpack/tree/lopez_cristian/storage/internal/transactions"
	"github.com/bootcamp-go/wave-5-backpack/tree/lopez_cristian/storage/pkg/web"
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

	db, err := db.NewConnection()
	if err != nil {
		log.Fatalf("error al conectar a la db: %v\n", err)
	}
	defer db.Close()

	// Init capas de transactions
	repo := transactions.NewRepository(db)
	service := transactions.NewService(repo)
	tr := handler.NewTransaction(service)

	gin.SetMode("test")
	router := gin.Default()

	// Middleware
	router.Use(web.TokenAuthMiddleware()) // el orden declarado de los Middleware afecta a su llamado

	// Router
	rt := router.Group("/transactions")
	{
		rt.GET("", tr.GetAll)
		rt.GET("/:id", tr.GetByID)

		rt.PUT("/:id", tr.Update)
		rt.PATCH("/:id", tr.Patch)

		rt.POST("", tr.CreateTransaction)

		rt.DELETE("/:id", tr.Delete)
	}

	if err := router.Run(); err != nil {
		panic(err)
	}
}
