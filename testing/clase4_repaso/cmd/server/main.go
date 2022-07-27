package main

import (
	"clase4_repaso/cmd/server/handler"
	"clase4_repaso/internal/products"
	"clase4_repaso/pkg/store"

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

	r.Run()
}
