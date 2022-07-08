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
		log.Fatal("error al cargar archivo")
	}
	repository := transactions.NewRepository(db)
	service := transactions.NewService(repository)
	transactions := handler.NewTransaction(service)

	router := gin.Default()
	tr := router.Group("/transactions")
	tr.POST("/", transactions.Store())
	tr.GET("/", transactions.GetAll())
	tr.PUT("/:id", transactions.Update())
	tr.DELETE("/:id", transactions.Delete())
	tr.PATCH("/:id", transactions.UpdateFields())
	router.Run()
}
