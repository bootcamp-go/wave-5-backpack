package main

import (
	"github.com/gin-gonic/gin"
	"goweb/internal/users"
	"goweb/cmd/server/handler"
)

func main(){
	
	repo := users.NewRepository()
	service := users.NewService(repo)
	u := handler.NewUser(service)


	router := gin.Default()
	userGroup := router.Group("/users")
	userGroup.GET("/", u.GetAllUsers())
	userGroup.GET("/:id", u.GetUserById())
	userGroup.POST("/", u.StoreUser()) 

	router.Run(":8080")

}

