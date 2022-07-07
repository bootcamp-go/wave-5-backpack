package main

import (
	"goweb/cmd/server/handler"
	"goweb/internal/users"

	"github.com/gin-gonic/gin"
)

func main() {
	repository := users.NewRepository()
	service := users.NewService(repository)
	p := handler.NewProduct(service)

	router := gin.Default()

	users := router.Group("/users")
	{
		users.GET("/", p.GetAll())
		users.GET("/:id", p.GetById())
		users.POST("/", p.Store())
	}

	router.Run(":8080")
}
