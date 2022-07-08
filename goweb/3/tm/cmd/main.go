package main

import (
	"goweb/3/tm/cmd/handler"
	"goweb/3/tm/internal/repository"
	"goweb/3/tm/internal/service"

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
		products.POST("", ph.Create())

		products.GET("", ph.ReadAll())
		products.GET("/:id", ph.Read())

		products.PUT("/:id", ph.Update())

		products.PATCH("/:id", ph.UpdateNamePrice())

		products.DELETE("/:id", ph.Delete())
	}

	router.Run()
}
