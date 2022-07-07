package main

import (
	"encoding/json"
	"fmt"
	"goweb/product"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

var productos = map[string]product.Products{
	"1": {
		Id:            1,
		Nombre:        "sandia",
		Color:         "verde",
		Precio:        20000,
		Stock:         5,
		Codigo:        "23fe2",
		Publicado:     true,
		FechaCreacion: "23/10/2022",
	},
	"2": {
		Id:            2,
		Nombre:        "Manzana",
		Color:         "Rojo",
		Precio:        50000,
		Stock:         12,
		Codigo:        "22fe2",
		Publicado:     true,
		FechaCreacion: "10/10/2022",
	},
}

func getAll(ctx *gin.Context) {
	path := "./products.json"
	file, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("Ocurrio un error: %s", err)
	}
	var p product.Products
	if err := json.Unmarshal([]byte(file), &p); err != nil {
		log.Fatal(err)
	}

	ctx.JSON(200, gin.H{
		"data": p,
	})
}

func getUserId(ctx *gin.Context) {

	producto, ok := productos[ctx.Param("Id")]

	if ok {
		ctx.String(200, "El producto es %s", producto.Nombre)
	} else {
		ctx.String(400, "No existe producto con esas caracteristicas")
	}
}

func main() {

	server := gin.Default()

	server.GET("/hello-world", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hola sebas!!",
		})
	})

	products := server.Group("/products")
	{
		products.GET("/", getAll)
		products.GET("/:Id", getUserId)
	}

	server.Run()

}
