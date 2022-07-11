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
	users.GET("/:id", u.GetById())
	users.POST("/", u.StoreUser())
	users.PUT("/:id", u.UpdateUser())

	server.Run(":8080")

}
