package main

import (
	"clase3_parte1/cmd/server/handler"
	"clase3_parte1/internal/products"
	"clase3_parte1/pkg/store"

	"github.com/gin-gonic/gin"
)

func main() {

	db := store.NewStore("products.json")

	repo := products.NewRepository(db)
	s := products.NewService(repo)
	p := handler.NewProduct(s)

	r := gin.Default()
	pr := r.Group("products")

	pr.GET("/", p.GetAll())
	pr.POST("/", p.Store())
	pr.PUT("/:id", p.Update())
	pr.PATCH("/:id", p.UpdateName())
	pr.DELETE("/:id", p.Delete())

	r.Run()
}
