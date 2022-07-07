package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

var products []product

func main() {
	router := gin.Default()

	if err := Read(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(products)

	//router.GET("/products", func(ctx *gin.Context) {
	// ctx.JSON(200, gin.H{
	// 	"message": "Hola Francisco",
	// })
	//})

	router.GET("/products", GetAll)
	router.GET("/products/", GetFilter)
	router.GET("/products/:id", GetProduct)
	router.POST("/newProduct", NewProduct())

	router.Run()

}
