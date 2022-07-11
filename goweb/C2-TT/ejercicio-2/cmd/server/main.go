package main

import (
	"C2-TT/cmd/server/handler"
	"C2-TT/internal/usuarios"

	"github.com/gin-gonic/gin"
)

func main() {
	repository := usuarios.NewRepository()
	service := usuarios.NewService(repository)
	handler := handler.NewUsuario(service)

	router := gin.Default()
	rTransaction := router.Group("usuarios")
	rTransaction.GET("/", handler.GetAll())
	rTransaction.POST("/", handler.Registrar())

	router.Run()
}
