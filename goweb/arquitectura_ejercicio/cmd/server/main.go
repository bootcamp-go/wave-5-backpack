package main

import (
	"github.com/anesquivel/wave-5-backpack/goweb/arquitectura_ejercicio/cmd/server/handler"
	"github.com/anesquivel/wave-5-backpack/goweb/arquitectura_ejercicio/internal/usuarios"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	repo := usuarios.NewRepository()
	repo.SetFromFile()
	service := usuarios.NewService(repo)
	u := handler.NewUsuario(service)

	r := gin.Default()
	users := r.Group("/usuarios")
	users.POST("/", u.Store())
	users.GET("/", u.GetAll())
	users.PUT("/:id", u.Update())
	users.PATCH("/:id", u.PatchLastNameAge())

	r.Run()
}
