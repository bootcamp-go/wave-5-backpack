package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rodrigoeshard/goweb/Practica2.2/cmd/handler"
	"github.com/rodrigoeshard/goweb/Practica2.2/internal/repositorio"
	"github.com/rodrigoeshard/goweb/Practica2.2/internal/servicio"
	"github.com/rodrigoeshard/goweb/Practica2.2/pkg/store"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("error al cargar el archivo .env")
	}
	db := store.NewStore(".json", "../users.json")
	repo := repositorio.NewRepository(db)
	service := servicio.NewService(repo)
	u := handler.NewUser(service)

	r := gin.Default()
	pr := r.Group("/users")
	pr.POST("/", u.CreateUser())
	pr.GET("/", u.GetAll())
	pr.PUT("/:id", u.UpdateUser())
	pr.PATCH("/:id", u.UpdateLastNameAge())
	pr.DELETE("/:id", u.Delete())
	r.Run()

}
