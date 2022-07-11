package main

import (
	"log"

	"github.com/bootcamp-go/wave-5-backpack/cmd/server/handler"
	"github.com/bootcamp-go/wave-5-backpack/internal/users"
	"github.com/bootcamp-go/wave-5-backpack/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("error archivo .env")
	}
	db := store.NewStore("users.json")
	if err := db.Ping(); err != nil {
		log.Fatal("error al intentar cargar archivo")
	}
	repo := users.NewRepositoy(db)
	service := users.NewService(repo)
	u := handler.NewUser(service)

	server := gin.Default()
	users := server.Group("/users")

	users.GET("/", u.GetAll())
	users.GET("/:id", u.GetById())
	users.POST("/", u.StoreUser())
	users.PUT("/:id", u.UpdateUser())

	server.Run(":8080")

}
