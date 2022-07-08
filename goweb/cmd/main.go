package main

import (
	"log"

	"github.com/bootcamp-go/wave-5-backpack/tree/lopez_cristian/goweb/cmd/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	rt := router.Group("/transactions")
	{
		rt.GET("", handler.GetAll)
		rt.GET("/search", handler.GetFilter)
		rt.GET("/:id", handler.GetByID)

		rt.POST("", handler.CreateTransaction)
	}

	if err := router.Run(":8080"); err != nil {
		log.Println("error en el server")
	}
}
