package main

import (
	"encoding/json"
	"os"

	"github.com/gin-gonic/gin"
)

type Product struct {
	Id         uint    `json:"id"`
	Name       string  `json:"name"`
	Color      string  `json:"color"`
	Price      float64 `json:"price"`
	Stock      uint    `json:"stock"`
	Code       string  `json:"code"`
	Published  bool    `json:"published"`
	Created_at string  `json:"created_at"`
}

func getProductList() ([]Product, error) {
	data, err := os.ReadFile("products.json")
	if err != nil {
		return []Product{}, err
	}
	var prList []Product
	err = json.Unmarshal(data, &prList)
	if err != nil {
		return []Product{}, err
	}
	return prList, nil
}

func main() {
	router := gin.Default()
	router.GET("/", func(ctx *gin.Context) {
		name := ctx.Request.URL.Query().Get("name")
		if name == "" {
			name = "Anonimo"
		}
		ctx.JSON(200, gin.H{
			"message": "Saludos " + name,
		})
	})

	getAll := func(ctx *gin.Context) {
		prList, err := getProductList()
		if err != nil {
			ctx.JSON(500, gin.H{
				"message": err.Error(),
			})
			return
		}
		ctx.JSON(200, prList)
	}

	router.GET("/products", getAll)
	router.Run()
}
