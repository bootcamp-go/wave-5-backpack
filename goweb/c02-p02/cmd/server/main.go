package main

import (
	"github.com/bootcamp-go/wave-5-backpack/tree/lugo_abelardo/cmd/server/handler"
	"github.com/bootcamp-go/wave-5-backpack/tree/lugo_abelardo/internal/products"
	"github.com/gin-gonic/gin"
)

func main() {
	repo := products.NewRepository()
	service := products.NewService(repo)
	p := handler.NewProduct((service))

	r := gin.Default()
	pr := r.Group("/products")
	pr.POST("/", p.Store())
	pr.GET("/", p.GetAll())
	r.Run()
}
