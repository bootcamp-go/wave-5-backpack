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
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error al intentar cargar el archivo de configuración")
	}

	db := store.NewStore("./usuarios.json")
	if err := db.Ping(); err != nil {
		log.Fatal("error al intentar conectarse a la db")
	}
	repo := users.NewRepository(db)
	service := users.NewService(repo)
	u := handler.NewUser(service)

	router := gin.Default()
	pr := router.Group("/users")
	{
		pr.GET("/", u.GetAll())
		pr.GET("/:id", u.GetById())
		pr.POST("/", u.Store())
		pr.PUT("/:id", u.Update())
		pr.PATCH("/:id", u.UpdateApellidoEdad())
		pr.DELETE("/:id", u.Delete())
	}

	router.Run(":8080")
}
