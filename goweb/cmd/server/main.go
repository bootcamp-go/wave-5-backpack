package main

import (
	"github.com/gin-gonic/gin"
	"goweb/cmd/server/controller"
	"goweb/internal/user"
)

var user1 *user.UserModel

//var users []user.UserModel
//var lastId int

func main() {

	router := gin.Default()

	router.GET("/index", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hola Pablo",
		})
	})
	usersGroup := router.Group("/users")
	{
		//GETs ------//
		usersGroup.GET("/", controller.GetUsers)
		usersGroup.GET("/:id", controller.GetUsersById)
		//POSTs ----//
		usersGroup.POST("/", controller.CreateUser())
	}
	err := router.Run()
	if err != nil {
		return
	}
}
