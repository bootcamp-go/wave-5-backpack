package main

import (
	"encoding/json"
	"fmt"
	"goweb/noStructure/product"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

var productos []product.Products
var lastID int

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
	idProduct, err := strconv.Atoi(ctx.Param("Id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error: ": err.Error(),
		})
		return
	}
	var arrayResult []product.Products
	exist := false

	for _, data := range productos {
		if data.Id == idProduct {
			exist = true
			arrayResult = append(arrayResult, data)
		}
	}

	if exist {
		ctx.JSON(http.StatusAccepted, gin.H{
			"Producto: ": arrayResult,
		})
	} else {
		ctx.String(400, "No existe un producto con ese id")
	}

}

func getProductsPublish(ctx *gin.Context) {
	filter, err := strconv.ParseBool(ctx.Query("publish"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error: ": err.Error(),
		})
		return
	}
	var arrayResult []product.Products

	for _, data := range productos {
		if data.Publicado == filter {
			arrayResult = append(arrayResult, data)
		}

	}

	ctx.JSON(http.StatusOK, gin.H{
		"responseData": arrayResult,
	})
}

func createProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token != "123456" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "token inválido",
			})
			return
		}
		var req product.Products

		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Error: ": err.Error(),
			})
			return
		}

		lastID++
		req.Id = lastID
		productos = append(productos, req)

		ctx.JSON(http.StatusAccepted, gin.H{
			"Productos: ": productos,
		})
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
		products.GET("/:Id", getProductId)
		products.GET("/publish", getProductsPublish)
		products.POST("/create", createProduct())
	}

	if err := server.Run(); err != nil {
		fmt.Println("error: ", err.Error())
		return
	}
}
