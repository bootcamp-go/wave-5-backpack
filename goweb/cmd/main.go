package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/bootcamp-go/wave-5-backpack/tree/lopez_cristian/goweb/cmd/handler"
	"github.com/bootcamp-go/wave-5-backpack/tree/lopez_cristian/goweb/internal/transactions"
	"github.com/bootcamp-go/wave-5-backpack/tree/lopez_cristian/goweb/pkg/storage"
)

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

	rt := router.Group("/transactions")
	{
		rt.GET("", t.GetAll)
		rt.GET("/search", t.GetFilter)
		rt.GET("/:id", t.GetByID)

		rt.PUT("/:id", t.Update)
		rt.PATCH("/:id", t.UpdateMontoCod)

		rt.POST("", t.CreateTransaction)

		rt.DELETE("/:id", t.Delete)
	}

	router.Run()
}
