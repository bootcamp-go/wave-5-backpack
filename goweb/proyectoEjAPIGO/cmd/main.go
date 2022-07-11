package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()
	// server.GET("/", HandlerRaiz)

	transacciones := server.Group("/transacciones")
	{

		transacciones.GET("/", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "holaa",
			})
		})
		// transacciones.GET("/:id", HandlerGetID)
		// transacciones.POST("/", AgregarTransacciones())
	}

	server.Run(":8005")
}
