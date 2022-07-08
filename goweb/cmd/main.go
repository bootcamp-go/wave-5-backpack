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

	gt := router.Group("/transaction")
	gt.POST("/", t.Create)
	gt.GET("/:id", t.GetOne)
	gt.PUT("/:id", t.Update)
	gt.DELETE("/:id", t.Delete)
	gt.PATCH("/:id", t.Update2)
	router.Run()
}
