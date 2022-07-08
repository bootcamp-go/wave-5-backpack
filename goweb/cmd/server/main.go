package main

import (
	"github.com/bootcamp-go/wave-5-backpack/goweb/cmd/server/handler"
	"github.com/bootcamp-go/wave-5-backpack/goweb/internal/usuarios"
	"github.com/gin-gonic/gin"
)

func main() {
	repo := usuarios.NewRepository()
	servi := usuarios.NewService(repo)
	u := handler.NewUsuario(servi)

	router := gin.Default()
	// router.GET("usuarios/:id", GetById)
	// router.GET("/usuarios", GetAll)
	// router.GET("/usuarios/filtroNombre", FilterByName)
	// router.GET("/usuarios/filtroApellido", FilterByLastName)
	// router.GET("/usuarios/filtroCorreo", filterByEmail)
	// router.GET("/usuarios/filtroEdad", filterByEdad)
	us := router.Group("/usuarios")
	us.PUT("/:id", u.Update())
	us.POST("/", u.Guardar())
	us.GET("/", u.GetAll())
	us.DELETE("/:id", u.Delete())
	us.PATCH("/:id", u.UpdateNameAndLastName())
	router.Run()
}
