package main

import (
	"encoding/json"
	"fmt"
	"goweb/product"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var productos = []product.Products{
	{
		Id:            "1",
		Nombre:        "sandia",
		Color:         "verde",
		Precio:        20000,
		Stock:         5,
		Codigo:        "23fe2",
		Publicado:     "true",
		FechaCreacion: "23/10/2022",
	},
	{
		Id:            "2",
		Nombre:        "Manzana",
		Color:         "Rojo",
		Precio:        50000,
		Stock:         12,
		Codigo:        "22fe2",
		Publicado:     "true",
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

func getProductId(ctx *gin.Context) {

	product := ctx.Param("Id")
	var result string
	exist := false

	for _, data := range productos {
		if data.Id == product {
			exist = true
			result = "El producto consultado es: " + data.Nombre
		}
	}

	if exist {
		ctx.String(200, result)
	} else {
		ctx.String(400, "No existe producto con esas caracteristicas")
	}

}

func getProductsPublish(ctx *gin.Context) {
	filter := ctx.Query("publish")
	var arrayResult []product.Products

	for _, data := range productos {
		if data.Publicado == filter {
			arrayResult = append(arrayResult, data)
		}

	}

	ctx.JSON(http.StatusAccepted, gin.H{
		"responseData": arrayResult,
	})
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
		products.GET("/:Id", getProductId)
		products.GET("/publish", getProductsPublish)
	}

	server.Run()

}
