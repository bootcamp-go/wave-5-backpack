package main

import (
	"log"

	"github.com/bootcamp-go/wave-5-backpack/tree/lopez_cristian/goweb/internal/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/transactions", handler.GetAll)
	router.GET("/transactions/search", handler.GetFilter)
	router.GET("/transactions/:id", handler.GetByID)

	if err := router.Run(":8080"); err != nil {
		log.Println("error en el server")
	}
}
