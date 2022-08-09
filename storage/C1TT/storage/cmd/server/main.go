package main

import (
	"storage/cmd/server/handler"
	cnn "storage/db"
	"storage/internal/products"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	loadEnv()

	db := cnn.MySQLConnection()
	repo := products.NewRepository(db)
	serv := products.NewService(repo)
	p := handler.NewProduct(serv)

	r := gin.Default()
	pr := r.Group("/api/v1/products")

	pr.POST("/", p.Store())
	pr.GET("/", p.GetByName())

	r.Run()
}

func loadEnv() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
}
