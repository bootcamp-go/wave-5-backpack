package main

import (
	"proyecto_meli/cmd/server/handler"
	"proyecto_meli/internal/products"

	"github.com/gin-gonic/gin"
)

func main() {
	repo := products.NewRepository()
	service := products.NewService(repo)
	p := handler.NewProduct(service)

	r := gin.Default()
	pr := r.Group("/products")
	//pr.POST("/", p.Store())
	pr.GET("/", p.GetAll())
	r.Run()
}
