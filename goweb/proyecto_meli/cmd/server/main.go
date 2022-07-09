package main

import (
	"log"
	"proyecto_meli/cmd/server/handler"
	"proyecto_meli/internal/products"
	store "proyecto_meli/pkg/store"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error al intentar cargar archivo .env")
	}
	db := store.NewStore("products.json")
	if err := db.Ping(); err != nil {
		log.Fatal("error al intentar cargar el archivo store")
	}
	repo := products.NewRepository(db)
	service := products.NewService(repo)
	p := handler.NewProduct(service)

	r := gin.Default()
	pr := r.Group("/products")
	pr.GET("/", p.GetAll())
	pr.GET("/:id", p.GetById())
	pr.GET("/filter", p.FilterList())
	pr.POST("/", p.Store())
	pr.PUT("/:id", p.Update())
	pr.DELETE("/:id", p.Delete())
	pr.PATCH("/:id", p.Update_Name_Price())
	r.Run()
}
