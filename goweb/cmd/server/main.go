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

	r := gin.Default()
	pr := r.Group("/users")
	pr.POST("/", u.Store())
	pr.GET("/", u.GetAll())
	r.Run()
}
