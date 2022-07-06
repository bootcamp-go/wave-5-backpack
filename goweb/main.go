package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)


func main() {
	router := gin.Default()

	router.GET("/hola/:name", func(ctx *gin.Context) {
		data := fmt.Sprintf("hola %s", ctx.Param("name"))

		ctx.JSON(200, gin.H{
			"message": data,
		})
	})

	if err := router.Run(":8080"); err != nil {
		log.Println("error en el server")
	}
}
