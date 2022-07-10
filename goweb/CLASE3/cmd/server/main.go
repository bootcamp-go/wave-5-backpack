package main

import (
	"clase2_parte2/cmd/server/handler"
	"clase2_parte2/internal/users"
	"clase2_parte2/pkg/store"
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
	if err := db.Val(); err != nil {
		log.Fatal("error al intentar cargar archivo")
	}

	repo := users.NewRepository(db)
	s := users.NewService(repo)
	p := handler.NewUsers(s)

	r := gin.Default()
	pr := r.Group("users")
	pr.GET("/", p.GetAll())
	pr.POST("/", p.Store())
	pr.PUT("/:id", p.Update())
	pr.PATCH("/:id", p.UpdateApellidoAndEdad())
	pr.DELETE("/:id", p.Delete())

	r.Run()
}
