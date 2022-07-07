package main

import (
	"github.com/gin-gonic/gin"
	"goweb/user"
)

var user1 *user.User
var users []user.User
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
		usersGroup.GET("/", user.GetUsers)
		usersGroup.GET("/:id", user.GetUsersById)
		//POSTs ----//
		usersGroup.POST("/", user.CreateUser(*user1))
	}
	router.Run()
}
