package main

import (
	"goweb/Clase2-2-WebServers/cmd/server/handler"
	"goweb/Clase2-2-WebServers/internal/usuarios"

	"github.com/gin-gonic/gin"
)

func main() {
	repo := usuarios.NewRepository()
	service := usuarios.NewService(repo)
	p := handler.NewUser(service)

	r := gin.Default()
	pr := r.Group("/usuarios")
	pr.POST("/", p.Store())
	pr.GET("/", p.GetAll())
	r.Run()
}
