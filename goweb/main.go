package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/products", func(ctx *gin.Context) {
		// ctx.JSON(200, gin.H{
		// 	"message": "Hola Francisco",
		// })

		GetAll(ctx.Writer, ctx.Request)

	})
	router.Run()

}
