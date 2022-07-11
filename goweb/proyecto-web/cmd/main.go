package main

import (
	"log"
	"proyecto-web/cmd/handlers"
	"proyecto-web/internal/transaction"
	"proyecto-web/pkg/store"

	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error al intentar cargar archivo .env")
	}
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
	servidor.DELETE("/transacciones/:id", handler.Delete())
	servidor.Run()
}
