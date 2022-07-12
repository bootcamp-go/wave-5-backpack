package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"clase4_parte1/cmd/server/handler"
	"clase4_parte1/internal/products"
	"clase4_parte1/pkg/store"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error al intentar cargar archivo .env")
	}

	db := store.NewStore("products.json")
	if err := db.Ping(); err != nil {
		log.Fatal("error al intentar cargar archivo")
	}

	repo := products.NewRepository(db)
	s := products.NewService(repo)
	p := handler.NewProduct(s)

	r := gin.Default()
	pr := r.Group("products")
	pr.GET("/", p.GetAll())
	pr.POST("/", p.Store())
	// pr.PUT("/:id", p.Update())
	// pr.PATCH("/:id", p.UpdateName())
	// pr.DELETE("/:id", p.Delete())

	r.Run()
}
