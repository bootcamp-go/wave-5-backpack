package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Usuario struct {
	Id       int
	Nombre   string
	Apellido string
	Email    string
	Altura   float64
	Activo   bool
	Fecha    string
}

var Usuarios []Usuario

func PaginaPrincipal(ctx *gin.Context) {
	ctx.String(200, "Hola Diego")
}

func GetAll(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": Usuarios,
	})
}

func main() {

	Usuarios = append(Usuarios, Usuario{1, "Diego", "Palacios", "d@example.com", 1.68, true, "21/06/2022"})
	Usuarios = append(Usuarios, Usuario{2, "Fernndo", "Palacios", "f@example.com", 1.70, true, "21/06/2022"})
	Usuarios = append(Usuarios, Usuario{3, "Cesar", "Parrado", "c@example.com", 1.80, true, "21/06/2022"})

	r := gin.Default()
	r.GET("", PaginaPrincipal)
	r.GET("/usuarios", GetAll)
	r.Run()

}
