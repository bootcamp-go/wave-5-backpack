package main

import (
	"GoWeb/cmd/server/handler"
	"GoWeb/internals/transactions"

	"github.com/gin-gonic/gin"
)

func main() {
	repo := transactions.NewRepository()
	service := transactions.NewService(repo)
	tran := handler.NewTransaction(service)

	router := gin.Default()
	tr := router.Group("/transacciones")
	tr.POST("/", tran.Store())
	tr.GET("/", tran.GetAll())
	tr.PUT("/:id", tran.Update())
	tr.DELETE("/:id", tran.Delete())
	tr.PATCH("/:id", tran.UpdateCode())
	router.Run()
}
