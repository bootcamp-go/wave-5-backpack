package main

import (
	"goweb/clase1_clase2/cmd/server/handler"
	"goweb/clase1_clase2/internal/products"

	"github.com/gin-gonic/gin"
)

func main() {

	repo := products.NewRepository()
	service := products.NewService(repo)
	p := handler.NewProduct(service)

	repo.ReadJson()

	router := gin.Default()
	pr := router.Group("/products")

	pr.POST("/", p.Store())
	pr.GET("/", p.GetAll())
	pr.GET("/:id", p.GetById())
	router.Run()
}
