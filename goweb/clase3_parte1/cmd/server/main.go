package main

import (
	"github.com/bootcamp-go/wave-5-backpack/tree/ospina_christian/goweb/clase3_parte1/cmd/server/handler"
	"github.com/bootcamp-go/wave-5-backpack/tree/ospina_christian/goweb/clase3_parte1/internal/productos"
	"github.com/gin-gonic/gin"
)

func main() {
	repo := productos.NewRepository()
	s := productos.NewService(repo)
	p := handler.NewProduct(s)

	router := gin.Default()
	rProductos := router.Group("/productos")
	rProductos.GET("/", p.GetAll())
	rProductos.POST("/", p.Store())
	rProductos.PUT("/:id", p.Update())
	rProductos.PATCH("/:id", p.UpdatePrecio())
	rProductos.DELETE("/:id", p.Delete())

	router.Run()
}
