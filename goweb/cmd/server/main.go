package main

import (
	"goweb/cmd/server/handler"
	"goweb/internal/products"
	"goweb/pkg/store"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error al intentar cargar el archivo .env")
	}

	db := store.InitStore("products.json")
	if err := db.Ping(); err != nil {
		log.Fatal("Error al intentar cargar el archivo")
	}
	repo := products.InitRepository(db)
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
