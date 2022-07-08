package main

import (
	"github.com/gin-gonic/gin"
	"goweb/cmd/server/controller"
	"goweb/internal/domain"
)

var user1 domain.User
var users []domain.User
var lastId int

func main() {

	router := gin.Default()

	router.GET("/index", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hola Pablo",
		})
	})
	usersGroup := router.Group("/users")
	{
		//GETs ------//
		usersGroup.GET("/", controller.GetUsers)
		usersGroup.GET("/:id", controller.GetUsersById)
		//POSTs ----//
		usersGroup.POST("/", controller.CreateUser(*user1))
	}
	router.Run()
}
