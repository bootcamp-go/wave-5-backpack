/*---------------------------------------------------------------------------------*

     Assignment:	Practica #1
         Author:	Israel Fabela
	   Language:	go1.18.3 darwin/arm64
		  Topic:	Go Web

	Description:
		‣	Exercise 1 - Generate PUT method
		‣	Exercise 2 - Generate DELETE method
		‣	Exercise 3 - Generate PATCH method

	© Mercado Libre - IT Bootcamp 2022

----------------------------------------------------------------------------------*/

package main

import (
	"goweb/clase3-go-web-tm/cmd/handler"
	"goweb/clase3-go-web-tm/internal/transactions"

	"github.com/gin-gonic/gin"
)

func main() {
	repo := transactions.NewRepository()
	service := transactions.NewService(repo)
	t := handler.NewTransaction(service)

	r := gin.Default()

	r.GET("/", handler.PaginaPrincipal)
	r.GET("/hola", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hola Israel ! 👋",
		})
	})

	pr := r.Group("/transactions")
	{
		pr.POST("/", t.Ecommerce())
		pr.GET("/", t.GetAll())
		pr.GET("/:id", t.GetOne())
		pr.PUT("/:id", t.Update())
		pr.PATCH("/:id", t.UpdateOne())
		pr.DELETE("/:id", t.Delete())
	}

	r.Run()
}
