package main

import (
	"ejercicioTT/cmd/server/handler"
	users "ejercicioTT/internal/users"
	"ejercicioTT/pkg/store"
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
		log.Fatal("error al intentar cargar archivo json con usuarios")
	}

	repo := users.NewRepository(db)
	serv := users.NewService(repo)
	user := handler.NewUser(serv)

	r := gin.Default()
	us := r.Group("usuarios")
	us.GET("/", user.GetAll())
	us.POST("/", user.Store())
	us.PUT("/:id", user.Update())
	us.PATCH("/:id", user.UpdateLastAge())
	us.DELETE("/:id", user.Delete())

	r.Run()
}
