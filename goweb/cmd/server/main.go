package main

import (
	"github.com/gin-gonic/gin"
	"goweb/cmd/server/handler"
	"goweb/internal/transactions"
)

func main() {

	repository := transactions.NewRepository()
	service := transactions.NewService(repository)
	handler := handler.NewTransaction(service)

	router := gin.Default()
	rTransaction := router.Group("transactions")
	rTransaction.GET("/", handler.GetAll())
	rTransaction.POST("/", handler.Store())
	rTransaction.PUT("/:id", handler.Update())
	rTransaction.PATCH("/:id", handler.UpdateTransmitter())
	rTransaction.DELETE("/:id", handler.Delete())
	router.Run(":3000")
}
