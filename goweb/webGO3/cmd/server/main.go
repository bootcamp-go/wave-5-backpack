package main

import (
	"log"

	"github.com/del_rio/web-server/cmd/server/controlador"
	"github.com/del_rio/web-server/internal/usuarios"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error al cargar el archivo ", err)
	}
	repo := usuarios.NewRepository()

	servicio := usuarios.NewService(repo)
	controladorUsuarios := controlador.NewControlador(servicio)
	router := gin.Default()
	usuarioRoute := router.Group("/usuarios")
	usuarioRoute.GET("/", controladorUsuarios.VerUsuarios())
	usuarioRoute.POST("/", controladorUsuarios.AgregarUsuarios())
	usuarioRoute.PUT("/:id", controladorUsuarios.ActualizarUsuario())
	usuarioRoute.PATCH("/:id", controladorUsuarios.ActualizarAtribUsuario())
	usuarioRoute.DELETE("/:id", controladorUsuarios.BorrarUsuario())
	router.Run()

}
