package main

import (
	"goweb/cmd/server/handler"
	"goweb/internal/users"
	"goweb/pkg/store"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("error al intentar cargar archivo .env")
	}

	db := store.NewStore("users.json")
	if err := db.Ping(); err != nil {
		log.Fatal("eroor al intentar cargar el archivo jsonn")
	}

	repository := users.NewRepository(db)
	service := users.NewService(repository)
	u := handler.NewUser(service)

	router := gin.Default()

	users := router.Group("/users")
	{
		users.GET("/", u.GetAll())
		users.GET("/:id", u.GetById())
		users.POST("/", u.Store())
		users.PUT("/:id", u.Update())
		users.DELETE("/:id", u.Delete())
		users.PATCH("/:id", u.Patch())
	}

	router.Run(":8080")
}
