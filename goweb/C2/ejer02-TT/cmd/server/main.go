package main

import (
	"ejer02-TT/cmd/server/handler"
	"ejer02-TT/internal/transactions"

	"github.com/gin-gonic/gin"
)

func main() {
	repository := transactions.NewRepository()
	service := transactions.NewService(repository)
	transactions := handler.NewTransaction(service)

	router := gin.Default()
	router.POST("/transactions", transactions.Store())
	router.GET("/transactions", transactions.GetAll())
	router.Run()
}
