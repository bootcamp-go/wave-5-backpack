package main

import (
	"github.com/gin-gonic/gin"
)

//EJERCICIO 3 DE LA GUIA
type Productos struct {
	Id            int     `json:"id"`
	Nombre        string  `json:"nombre"`
	Color         string  `json:"color"`
	Precio        float64 `json:"precio"`
	Stock         int     `json:"stock"`
	Codigo        string  `json:"codigo"`
	Publicado     bool    `json:"publicado"`
	FechaCreación string  `json:"fecha_creacion"`
}

func main() {
	router := gin.Default()

	//ENDPOINT para listar productos
	router.GET("/productos", GetAll)

	router.Run()
}

func GetAll(ctx *gin.Context) {
	//products := []Productos{}
	products := []Productos{
		{Id: 1, Nombre: "Papaya", Color: "Verde", Precio: 1500, Stock: 21, Codigo: "1s1", Publicado: true, FechaCreación: "06/06/2022"},
		{Id: 2, Nombre: "Aguacate", Color: "Verde", Precio: 2500, Stock: 11, Codigo: "1a1", Publicado: true, FechaCreación: "06/06/2022"},
		{Id: 3, Nombre: "Melon", Color: "Verde", Precio: 1200, Stock: 9, Codigo: "1m1", Publicado: true, FechaCreación: "06/06/2022"},
	}
	if len(products) <= 0 {
		ctx.JSON(200, gin.H{
			"All Products": "No hay productos disponibles",
		})
		return
	}
	ctx.JSON(200, gin.H{
		"All Products": products,
	})
}
