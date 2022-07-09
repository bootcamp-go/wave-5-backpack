/*---------------------------------------------------------*

     Assignment:	Practica #2
         Author:	Israel Fabela
	   Language:	go1.18.3 darwin/arm64
		  Topic:	Go Web

	Description:
		‣	Exercise 1 - ENV Configuration
		‣	Exercise 2 - Save information
		‣	Exercise 3 - Read information

	© Mercado Libre - IT Bootcamp 2022

----------------------------------------------------------*/

package main

import (
	"goweb/clase3-go-web-tt/cmd/handler"
	"goweb/clase3-go-web-tt/internal/transactions"
	"goweb/clase3-go-web-tt/pkg/bank"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("error al intentar cargar archivo .env")
	}

	db := bank.NewBank("transacciones.json")
	if err := db.Ping(); err != nil {
		log.Fatal("error al intentar cargar archivo")
	}

	repo := transactions.NewRepository(db)
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
		pr.GET("/", t.GetAll())
		pr.GET("/:id", t.GetOne())
		pr.PUT("/:id", t.Update())
		pr.POST("/", t.Ecommerce())
		pr.PATCH("/:id", t.UpdateOne())
		pr.DELETE("/:id", t.Delete())
	}

	r.Run()
}
