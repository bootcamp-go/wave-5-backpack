package main

import (
	"goweb/productos_capas/cmd/server/handler"
	"goweb/productos_capas/internal/products"
	"goweb/productos_capas/pkg/store"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("cmd/server/.env")
	if err != nil {
		log.Fatal("error al intentar cargar el archivo .env")
	}

	db := store.NewStore("products.json")
	if err := db.Ping(); err != nil {
		log.Fatal("error al intentar cargar archivo")
	}

	repo := products.NewRepository(db)
	service := products.NewService(repo)
	p := handler.NewProduct(service)

	router := gin.Default()
	pr := router.Group("/productos")
	pr.GET("/", p.GetAll())
	pr.GET("/:id", p.GetByID())
	pr.POST("/", p.Store())
	pr.PUT("/:id", p.Update())
	pr.PATCH("/:id", p.UpdateNamePrice())
	pr.DELETE("/:id", p.Delete())
	router.Run()
}
