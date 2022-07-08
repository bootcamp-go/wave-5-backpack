package main

import (
	"github.com/bootcamp-go/wave-5-backpack/tree/lopez_cristian/goweb/cmd/handler"
	"github.com/bootcamp-go/wave-5-backpack/tree/lopez_cristian/goweb/internal/transactions"
	"github.com/gin-gonic/gin"
)

func main() {
	repo := transactions.NewRepository()
	service := transactions.NewService(repo)
	handler := handler.NewTransaction(service)

	router := gin.Default()

	rt := router.Group("/transactions")
	{
		rt.GET("", handler.GetAll)
		rt.GET("/search", handler.GetFilter)
		rt.GET("/:id", handler.GetByID)

		rt.POST("", handler.CreateTransaction)
	}

	router.Run()
}
