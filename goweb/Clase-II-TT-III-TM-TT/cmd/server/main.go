package main

import (
	"arquitectura/cmd/server/handler"
	"arquitectura/internal/transactions"
	"arquitectura/pkg/store"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

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
	router.POST("/transactions", transactions.Store())
	router.GET("/transactions", transactions.GetAll())
	router.PUT("/transactions/:id", transactions.Update())
	router.DELETE("/transactions/:id", transactions.Delete())
	router.PATCH("/transactions/:id", transactions.UpdateCodeAmount())
	router.Run()
}
