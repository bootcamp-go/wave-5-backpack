package main

import (
	"github.com/bootcamp-go/wave-5-backpack/goweb/cmd/server/handler"
	"github.com/bootcamp-go/wave-5-backpack/goweb/internal/users"
	"github.com/gin-gonic/gin"
)

func main() {
	repo := users.NewRepository()
	service := users.NewService(repo)
	userHandler := handler.NewUser(service)

	router := gin.Default()
	usersGroup := router.Group("/users")
	usersGroup.POST("/", userHandler.Store())
	usersGroup.GET("/", userHandler.GetAll())
	usersGroup.PUT("/:id", userHandler.Update())
	usersGroup.PATCH("/:id", userHandler.UpdateLastNameAndAge())
	usersGroup.DELETE("/:id", userHandler.Delete())
	router.Run()
}
