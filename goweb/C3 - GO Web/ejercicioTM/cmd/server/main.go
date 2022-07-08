package main

import (
	"ejercicioTM/cmd/server/handler"
	"ejercicioTM/internal/users"

	"github.com/gin-gonic/gin"
)

func main() {
	repo := users.NewRepository()
	serv := users.NewService(repo)
	user := handler.NewUser(serv)

	r := gin.Default()
	us := r.Group("usuarios")
	us.GET("/", user.GetAll())
	us.POST("/", user.Store())
	us.PUT("/:id", user.Update())
	us.PATCH("/:id", user.UpdateLastAge())
	us.DELETE("/:id", user.Delete())

	r.Run()
}
