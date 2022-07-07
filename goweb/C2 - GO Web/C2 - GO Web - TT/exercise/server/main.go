package main

import (
	"exercise/internal/users"
	"exercise/server/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	repoUser := users.NewRepository()
	serviceUser := users.NewService(repoUser)
	usuario := handler.NewUser(serviceUser)

	r := gin.Default()
	us := r.Group("/usuarios")
	us.POST("/", usuario.Store())
	us.GET("/", usuario.GetAll())
	r.Run()
}
