package main

import (
	"C2-TT/cmd/server/handler"
	"C2-TT/internal/usuarios"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	repository := usuarios.NewRepository()
	service := usuarios.NewService(repository)
	handler := handler.NewUsuario(service)

	router := gin.Default()
	rUsuario := router.Group("usuarios")
	rUsuario.GET("/", handler.GetAll())
	rUsuario.POST("/", handler.Registrar())
	rUsuario.PUT("/:id", handler.Modificar())
	rUsuario.DELETE("/:id", handler.Eliminar())
	rUsuario.PATCH("/:id", handler.ModificarAE())

	router.Run()
}
