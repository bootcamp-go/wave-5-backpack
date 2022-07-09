/*---------------------------------------------------------*

     Assignment:	Practica #2
         Author:	Israel Fabela
	   Language:	go1.18.3 darwin/arm64
		  Topic:	Go Web

	Description:
		‣	Exercise 1 - Generate internal package
		‣	Exercise 2 - Generate server package

	© Mercado Libre - IT Bootcamp 2022

----------------------------------------------------------*/

package main

import (
	"goweb/clase2-go-web-tt/cmd/handler"
	"goweb/clase2-go-web-tt/internal/transactions"

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
	}

	r.Run()
}
