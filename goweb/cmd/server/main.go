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
		pr.GET("/", u.GetAll())
		pr.GET("/:id", u.GetById())
		pr.POST("/", u.Store())
		pr.PUT("/:id", u.Update())
		pr.PATCH("/:id", u.UpdateApellidoEdad())
		pr.DELETE("/:id", u.Delete())
	}

	router.Run(":8080")
}
