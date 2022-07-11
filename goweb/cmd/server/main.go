package main

import (
	"goweb/cmd/server/handler"
	"goweb/internal/users"
	"goweb/pkg/dataStore"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error al intentar cargar archivo .env")
	}

	db := dataStore.NewStore("users.json")
	if err := db.Ping(); err != nil {
		log.Fatal("error al intentar cargar archivo")
	}
	repo := users.NewRepository(db)
	service := users.NewService(repo)
	u := handler.NewUsers(service)

	r := gin.Default()
	us := r.Group("/products")
	us.POST("/", u.NewUser())
	us.GET("/", u.GetAll())
	us.PUT("/:id", u.Update())
	us.PATCH("/:id", u.UpdateName())
	us.DELETE("/:id", u.Delete())
	r.Run()
}
