package main

import (
	"clase2_2/cmd/server/handler"
	"clase2_2/internal/users"

	"github.com/gin-gonic/gin"
)

func main() {
	repository := users.NewRepository()
	service := users.NewService(repository)
	handler := handler.NewUser(service)
	router := gin.Default()
	//clase 2_2
	rTransaction := router.Group("users")
	rTransaction.GET("/", handler.GetAll())
	rTransaction.POST("/", handler.AddUser())
	//clase 3_1
	rTransaction.PUT("/:id", handler.UpdateUser())
	rTransaction.DELETE("/:id", handler.Delete())
	rTransaction.PATCH("/:id", handler.UpdateUserName())
	router.Run()
}
