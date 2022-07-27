package main

import (
	"clase4_parte1/cmd/server/handler"
	"clase4_parte1/internal/products"
	"clase4_parte1/pkg/store"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("error al cargar archivo .env \n")
	}
	db := store.New(store.FileType, "products.json")
	repo := products.NewRepository(db)
	service := products.NewService(repo)
	p := handler.NewProduct(service)

	r := gin.Default()
	pr := r.Group("/products")
	pr.POST("/", p.Store())
	pr.GET("/", p.GetAll())
	r.Run()
}
