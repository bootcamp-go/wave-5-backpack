package main

import (
	"fmt"
	"goweb/cmd/server/handler"
	"goweb/internal/users"
	"goweb/pkg/store"

	"github.com/gin-gonic/gin"
)

func main() {

	db := store.NewStore("users.json")

	if err := db.Ping(); err != nil {
		fmt.Println(err.Error())
		return
	}

	repository := users.NewRepository(db)
	service := users.NewService(repository)
	users := handler.NewUser(service)

	router := gin.Default()

	useresRoute := router.Group("/users")
	useresRoute.GET("/", users.GetAll())
	useresRoute.GET("/:id", users.GetById())
	useresRoute.POST("/", users.Create())
	useresRoute.PUT("/:id", users.Update())
	useresRoute.DELETE("/:id", users.Delete())
	router.Run()
}
