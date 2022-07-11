package main

import (
	"github.com/gin-gonic/gin"
	"goweb/internal/users"
	"goweb/cmd/server/handler"
	"github.com/joho/godotenv"
	"goweb/pkg/store"
	"log"
)

func main(){

	err := godotenv.Load()
	if err != nil {
		log.Fatal("error al intentar cargar archivo .env")
	}

	db := store.NewStore("users.json")
	if err := db.Ping(); err != nil {
		log.Fatal("error al intentar cargar archivo")
	}
	
	repo := users.NewRepository(db)
	service := users.NewService(repo)
	u := handler.NewUser(service)


	router := gin.Default()
	userGroup := router.Group("/users")
	userGroup.GET("/", u.GetAllUsers())
	userGroup.GET("/:id", u.GetUserById())
	userGroup.POST("/", u.StoreUser())
	userGroup.PUT("/:id", u.UpdateTotal())
	userGroup.PATCH("/:id", u.UpdatePartial())
	userGroup.DELETE("/:id", u.Delete())

	router.Run() // si no especifico ningún puerto, por ej ":3001" toma por defecto el 8080 

}

