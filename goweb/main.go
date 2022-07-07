package main

import (
	"github.com/gin-gonic/gin"
	"goweb/user"
)

var u = user.User{}
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
		usersGroup.POST("/", user.CreateUser)
	}
	router.Run()
}
