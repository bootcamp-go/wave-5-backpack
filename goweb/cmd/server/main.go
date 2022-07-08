package main

import (
	"log"

	"github.com/bootcamp-go/wave-5-backpack/goweb/cmd/server/handler"
	"github.com/bootcamp-go/wave-5-backpack/goweb/internal/users"
	"github.com/bootcamp-go/wave-5-backpack/goweb/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	db := store.NewStore("../../users.json")
	repo := users.NewRepository(db)
	service := users.NewService(repo)
	userHandler := handler.NewUser(service)

	error := godotenv.Load("../../../.env")
	if error != nil {
		log.Fatal("error al intentar cargar el .env")
	}

	router := gin.Default()
	usersGroup := router.Group("/users")
	usersGroup.POST("/", userHandler.Store())
	usersGroup.GET("/", userHandler.GetAll())
	usersGroup.PUT("/:id", userHandler.Update())
	usersGroup.PATCH("/:id", userHandler.UpdateLastNameAndAge())
	usersGroup.DELETE("/:id", userHandler.Delete())
	router.Run()
}
