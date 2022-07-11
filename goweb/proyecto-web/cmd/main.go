package main

import (
	"proyecto-web/cmd/handlers"
	"proyecto-web/internal/transaction"
	"proyecto-web/pkg/store"

	"github.com/gin-gonic/gin"
)

func main() {
	bd := store.NewStore("transacciones.json")
	r := transaction.NewRepository(bd)
	service := transaction.NewService(r)
	handler := handlers.NewTransactionHandler(service)
	servidor := gin.Default()

	servidor.GET("/transacciones", handler.GetAll())
	servidor.GET("/transacciones/:id", handler.GetById())
	servidor.POST("/transacciones", handler.Create())
	servidor.PUT("/transacciones/:id", handler.Update())
	servidor.PATCH("/transacciones/:id", handler.UpdateParcial())
	servidor.Run()
}
