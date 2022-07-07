package main

import (
	"bufio"
	"fmt"
	"goweb/clase1_parte1/internal"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Ingresa tu nombre: ")
	name, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	msg := "Hola " + name
	router.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": msg,
		})
	})

	router.GET("/productos", internal.GetAll)
	router.GET("/productos/:id", internal.GetById)

	// POST

	router.POST("/productos", internal.Post)

	router.Run()
}
