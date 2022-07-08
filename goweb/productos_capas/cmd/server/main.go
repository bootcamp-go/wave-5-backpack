package main

import (
	"goweb/productos_capas/cmd/server/handler"
	"goweb/productos_capas/internal/products"

	"github.com/gin-gonic/gin"
)

func main() {
	repo := products.NewRepository()
	service := products.NewService(repo)
	p := handler.NewProduct(service)
	repo.ReadJSON()

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
