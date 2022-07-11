package main

import (
	"github.com/bootcamp-go/wave-5-backpack/tree/olivera_sebastian/goweb/clase2_parte2+siguientes/cmd/server/handler"
	"github.com/bootcamp-go/wave-5-backpack/tree/olivera_sebastian/goweb/clase2_parte2+siguientes/internal/transactions"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
	repository := transactions.NewRepository()
	service := transactions.NewService(repository)
	handler := handler.NewTransaction(service)

	router := gin.Default()
	rTransaction := router.Group("transactions")
	rTransaction.GET("/", handler.GetAll())
	rTransaction.POST("/", handler.Store())
	rTransaction.PUT("/:id", handler.Update())
	rTransaction.DELETE("/:id", handler.Delete())
	rTransaction.PATCH("/:id", handler.UpdateReceptorYMonto())

	router.Run(":8000")
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}
