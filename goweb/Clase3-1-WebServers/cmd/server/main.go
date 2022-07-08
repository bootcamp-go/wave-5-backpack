package main

import (
	"goweb/Clase3-1-WebServers/cmd/server/handler"
	"goweb/Clase3-1-WebServers/internal/usuarios"

	"github.com/gin-gonic/gin"
)

func main() {
	repo := usuarios.NewRepository()
	service := usuarios.NewService(repo)
	user := handler.NewUser(service)

	r := gin.Default()
	ur := r.Group("/usuarios")
	ur.POST("/", user.Store())
	ur.GET("/", user.GetAll())
	ur.PUT("/:id", user.Update())
	ur.PATCH("/:id", user.UpdateSurnameAndAge())
	ur.DELETE("/:id", user.Delete())
	r.Run()
}
