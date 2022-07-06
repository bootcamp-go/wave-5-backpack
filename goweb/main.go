package main

import (
	"encoding/json"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/products", func(ctx *gin.Context) {
		// ctx.JSON(200, gin.H{
		// 	"message": "Hola Francisco",
		// })
		products := GetAll()

		jsonData, err := json.Marshal(products)

		if err != nil {
			log.Fatal(err)
		} else {
			ctx.JSON(200, string(jsonData))
		}

	})
	router.Run()

}
