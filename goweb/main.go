package main

import (
	"github.com/gin-gonic/gin"
	"goweb/user"
)

func main() {

	router := gin.Default()

	router.GET("/index", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hola Pablo",
		})
	})
	usersGroup := router.Group("/users")
	{
		usersGroup.GET("/", user.GetUsersHandler)
		usersGroup.GET("/:id", user.GetUsersByIdHandler)
		// POST ----
		usersGroup.POST("/", user.CreateUser)
	}
	router.Run()
}
