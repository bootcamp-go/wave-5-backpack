package main

import (
	"github.com/del_rio/web-server/cmd/server/controlador"
	"github.com/del_rio/web-server/internal/usuarios"
	"github.com/gin-gonic/gin"
)

func main() {
	repo := usuarios.NewRepository()
	servicio := usuarios.NewService(repo)
	controladorUsuarios := controlador.NewControlador(servicio)
	router := gin.Default()
	usuarioRoute := router.Group("/usuarios")
	usuarioRoute.GET("/", controladorUsuarios.VerUsuarios())
	usuarioRoute.POST("/", controladorUsuarios.AgregarUsuarios())
	router.Run()

}
