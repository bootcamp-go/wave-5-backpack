package main

import (
	"github.com/bootcamp-go/wave-5-backpack/cmd/server/handler"
	"github.com/bootcamp-go/wave-5-backpack/internal/users"
	"github.com/gin-gonic/gin"
)

func main() {

	repo := users.NewRepositoy()
	service := users.NewService(repo)
	u := handler.NewUser(service)

	server := gin.Default()
	users := server.Group("/users")

	users.GET("/", u.GetAll())
	users.POST("/", u.StoreUser())

	server.Run(":8080")

}
