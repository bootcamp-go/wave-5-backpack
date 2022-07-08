package main

import (
	"arquitectura/cmd/server/handler"
	"arquitectura/internal/transactions"

	"github.com/gin-gonic/gin"
)

func main() {
	repository := transactions.NewRepository()
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
