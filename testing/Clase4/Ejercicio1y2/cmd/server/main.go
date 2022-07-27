package main

import (
	"clase2_2/cmd/server/handler"
	"clase2_2/internal/users"
	"clase2_2/pkg/storage"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("error al obtener las variables de ambiente ")
	}
	db := storage.NewStore("usuarios.json")
	repository := users.NewRepository(db)
	service := users.NewService(repository)
	handler := handler.NewUser(service)
	router := gin.Default()
	//clase 2_2
	us := router.Group("users")
	us.GET("/", handler.GetAll())
	us.POST("/", handler.AddUser())
	//clase 3_1
	us.PUT("/:id", handler.UpdateUser())
	us.DELETE("/:id", handler.Delete())
	router.Run()
}
