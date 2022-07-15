package main

import (
	//"fmt"
	"net/http"
	//"time"

	"github.com/gin-gonic/gin"
)

type Productos struct {
	Id int
	Nombre string
	Color string
	Precio int
	Stock int
	Codigo string
	Publicado bool
	Fecha string
}

func main()  {
	// creacion del router
	router := gin.Default()

	router.GET("/hello-world", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "Hola Nicolas!"})
	})

	router.GET("/productos", GetAll)

	router.Run()
}

func GetAll(ctx *gin.Context) {
	products := []Productos{
		{Id: 1, Nombre: "Chocramo", Color: "Cafe", Precio: 3000, Stock: 100, Codigo: "CASD23", Publicado: true, Fecha: "1992"},
		{Id: 2, Nombre: "Gala", Color: "Cafe", Precio: 2000, Stock: 100, Codigo: "FRE341", Publicado: true, Fecha: "1996"},
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": products,
	})
}
