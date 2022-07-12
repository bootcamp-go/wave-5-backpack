package main

import (
	"github.com/abelardolugo/go-web/cmd/server/handler"
	"github.com/abelardolugo/go-web/internal/products"

	"github.com/gin-gonic/gin"
)

func main() {
	repo := products.NewRepository()
	s := products.NewService(repo)
	h := handler.NewProduct(s)

	router := gin.Default()
	pr := router.Group("/productos")
	pr.GET("/", h.GetAll())
	pr.POST("/create", h.Store())
	pr.PUT("/:id/update", h.Update())
	pr.PATCH("/:id/update-name", h.UpdateName())
	pr.PATCH("/:id/update-name-price", h.UpdateNamePrice())
	pr.DELETE("/:id/delete", h.Delete())
	router.Run(":8080")
}
