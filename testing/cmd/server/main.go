package main

import (
	"github.com/bootcamp-go/go-testing/cmd/server/handler"
	"github.com/bootcamp-go/go-testing/internal/products"
	"github.com/bootcamp-go/go-testing/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	loadEnv()

	r := gin.Default()
	db := store.NewFileStore("./products.json")
	repo := products.NewRepository(db)
	serv := products.NewService(repo)
	p := handler.NewProducts(serv)

	pr := r.Group("products")
	pr.GET("/", p.GetAll())

	r.Run()
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}
