package main

import (
	"github.com/bootcamp-go/wave-5-backpack/goweb/cmd/handler"
	"github.com/bootcamp-go/wave-5-backpack/goweb/internal/products"
	"github.com/gin-gonic/gin"
)

func main() {

	repository := products.NewRepository()
	service := products.NewService(repository)
	product := handler.NewProduct(service)

	router := gin.Default()

	group := router.Group("products")

	group.GET("/", product.GetAll())
	group.POST("/", product.Store())
	group.PUT("/:id", product.UpdateAll())
	group.PATCH("/:id", product.Update())
	group.DELETE("/:id", product.Delete())

	//group.GET("/", GetFilter)
	//group.GET("/:id", GetProduct)

	router.Run()

}
