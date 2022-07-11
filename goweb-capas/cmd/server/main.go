package main

import (
	"goweb-capas/cmd/server/handler"
	"goweb-capas/internal/products"

	"github.com/gin-gonic/gin"
)

func main() {
	repo := products.NewRepository()
	service := products.NewService(repo)
	p := handler.NewProduct(service)

	r := gin.Default()
	pr := r.Group("/products")
	pr.POST("/", p.Store())
	pr.GET("/", p.GetAll())
	pr.PUT("/:id", p.Update())
	pr.PATCH("/:id", p.Patch())
	pr.DELETE("/:id", p.Delete())
	r.Run()
}
