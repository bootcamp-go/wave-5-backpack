package main

import (
	"log"

	"github.com/bootcamp-go/wave-5-backpack/tree/ospina_christian/goweb/clase3_parte2/cmd/server/handler"
	"github.com/bootcamp-go/wave-5-backpack/tree/ospina_christian/goweb/clase3_parte2/internal/productos"
	"github.com/bootcamp-go/wave-5-backpack/tree/ospina_christian/goweb/clase3_parte2/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error al intentar cargar el archivo .env")
	}

	db := store.NewStore("productos.json")
	if err := db.Ping(); err != nil {
		log.Fatal("error al intentar cargar el archivo")
	}

	repo := productos.NewRepository(db)
	s := productos.NewService(repo)
	p := handler.NewProduct(s)

	router := gin.Default()
	rProductos := router.Group("productos")
	rProductos.GET("/", p.GetAll())
	rProductos.GET("/:id", p.GetForId())
	rProductos.POST("/", p.Store())
	rProductos.PUT("/:id", p.Update())
	rProductos.PATCH("/:id", p.UpdatePrecio())
	rProductos.DELETE("/:id", p.Delete())

	router.Run()
}
