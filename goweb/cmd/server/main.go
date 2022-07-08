package main

import (
	"github.com/bootcamp-go/wave-5-backpack/goweb/cmd/server/handler"
	"github.com/bootcamp-go/wave-5-backpack/goweb/internal/users"
	"github.com/gin-gonic/gin"
)

func main() {
	repo := users.NewRepository()
	service := users.NewService(repo)
	u := handler.NewUser(service)

	router := gin.Default()
	pr := router.Group("/users")
	{
		pr.POST("/", u.Store())
		pr.GET("/", u.GetAll())
	}

	router.Run(":8080")
}
