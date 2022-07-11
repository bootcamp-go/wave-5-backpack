package main

import (
	"goweb/cmd/server/handler"
	"goweb/internal/users"

	"github.com/gin-gonic/gin"
)

func main() {
	repository := users.NewRepository()
	service := users.NewService(repository)
	users := handler.NewUser(service)

	router := gin.Default()

	useresRoute := router.Group("/users")
	useresRoute.GET("/", users.GetAll())
	router.Run()
}
