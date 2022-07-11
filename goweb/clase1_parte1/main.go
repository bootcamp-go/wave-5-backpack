package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

//EJERCICIO 1 Y 2 DE LA GUIA GOWEB CLASE 1 PARTE 1
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

	products := []Productos{
		{Id: 1, Nombre: "Papaya", Color: "Verde", Precio: 1500, Stock: 21, Codigo: "1s1", Publicado: true, FechaCreación: "06/06/2022"},
		{Id: 2, Nombre: "Aguacate", Color: "Verde", Precio: 2500, Stock: 11, Codigo: "1a1", Publicado: true, FechaCreación: "06/06/2022"},
		{Id: 3, Nombre: "Melon", Color: "Verde", Precio: 1200, Stock: 9, Codigo: "1m1", Publicado: true, FechaCreación: "06/06/2022"},
	}

	uri := "/hello"

	//ENDPOINT que nos saluda
	router.GET(uri, func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello Christian! 👋 Bootcampers",
			"uri":     uri,
		})
	})

	jsonProducts, err := json.MarshalIndent(products, "", " ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(jsonProducts))
	os.WriteFile("productos.json", jsonProducts, 0644)
	router.Run()
}
