package main

import (
	"goweb/cmd/server/handler"
	"goweb/internal/products"

	"github.com/gin-gonic/gin"
)

func main() {

	repo := products.InitRepository()
	service := products.InitService(repo)
	p := handler.InitProduct(service)

	r := gin.Default()
	pr := r.Group("/products")
	pr.POST("/create", p.CreateProduct())
	pr.GET("/", p.GetAll())
	pr.PUT("/:id", p.Update())
	pr.DELETE("/:id", p.Delete())
	pr.PATCH("/:id", p.UpdateOne())
	r.Run()
}
