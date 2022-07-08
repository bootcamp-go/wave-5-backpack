package main

import (
	"goweb/Clase3-2-WebServers/cmd/server/handler"
	"goweb/Clase3-2-WebServers/internal/usuarios"
	"goweb/Clase3-2-WebServers/pkg/store"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error al intentar cargar archivo .env")
	}
	db := store.NewStore("users.json")
	if err := db.Ping(); err != nil {
		log.Fatal("error al intentar cargar archivo")
	}

	repo := usuarios.NewRepository(db)
	service := usuarios.NewService(repo)
	user := handler.NewUser(service)

	r := gin.Default()
	ur := r.Group("/usuarios")
	ur.POST("/", user.Store())
	ur.GET("/", user.GetAll())
	ur.PUT("/:id", user.Update())
	ur.PATCH("/:id", user.UpdateSurnameAndAge())
	ur.DELETE("/:id", user.Delete())
	r.Run()
}
