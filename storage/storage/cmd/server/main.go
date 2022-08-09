package main

import (
	"github.com/gin-gonic/gin"
	"storage/cmd/server/handler"
	"storage/internal/product"
	"storage/pkg/store"
)

func main() {
	store.Init()
	repo := product.NewRepository(store.StorageDB)
	service := product.NewService(repo)
	controller := handler.NewProduct(service)

	r := gin.Default()

	r.GET("/prod", controller.GetAll())
	r.GET("/prod/:name", controller.Search())
	r.POST("/prod", controller.Store())
	r.PATCH("/prod/:id", controller.Update())
	r.DELETE("/prod/:id", controller.Delete())

	r.Run()
}
