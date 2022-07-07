package main

import (
	"goweb/2/tt/cmd/handler"
	"goweb/2/tt/internal/repository"
	"goweb/2/tt/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	pr := repository.NewRepository()
	ps := service.NewService(pr)
	ph := handler.NewHandler(ps)

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	products := router.Group("/products")
	{
		products.GET("", ph.GetAll)
		products.POST("", ph.Post())
	}

	router.Run()
}
