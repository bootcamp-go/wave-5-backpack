package main

import (
	"goweb/cmd/handler"
	"goweb/internal/transactions"

	"github.com/gin-gonic/gin"
)

func main() {
	repo := transactions.NewRepository()
	service := transactions.NewService(repo)
	t := handler.NewTransaction(service)

	router := gin.Default()
	router.GET("/transactions", t.GetAll)
	router.GET("/transaction/:id", t.GetOne)
	router.POST("/transaction", t.Create)
	router.Run()
}
