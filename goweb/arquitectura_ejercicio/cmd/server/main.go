package main

import (
	"log"

	"github.com/anesquivel/wave-5-backpack/goweb/arquitectura_ejercicio/cmd/server/handler"
	"github.com/anesquivel/wave-5-backpack/goweb/arquitectura_ejercicio/internal/usuarios"
	"github.com/anesquivel/wave-5-backpack/goweb/arquitectura_ejercicio/pkg/store"
	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	err := godotenv.Load()
	if err != nil {
		log.Fatal("error al intentar cargar archivo .env")
	}

	db := store.NewStore("usuarios.json")

	if errR := db.Ping(); errR != nil {
		log.Fatal("error al intentar cargar archivo")
	}

	repo := usuarios.NewRepository(db)
	service := usuarios.NewService(repo)
	u := handler.NewUsuario(service)

	r := gin.Default()
	users := r.Group("/usuarios")
	users.POST("/", u.Store())
	users.GET("/", u.GetAll())
	users.PUT("/:id", u.Update())
	users.PATCH("/:id", u.PatchLastNameAge())
	users.DELETE("/:id", u.Delete())

	r.Run()
}
