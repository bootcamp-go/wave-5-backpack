package main

import (
	"ejercicio-2-3/cmd/server/handler"
	"ejercicio-2-3/internal/usuarios"
	"ejercicio-2-3/pkg/registro"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	db := registro.NewFileStore(registro.FileType, "usuarios.json")

	repository := usuarios.NewRepository(db)
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
