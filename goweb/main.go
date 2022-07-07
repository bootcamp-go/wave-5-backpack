package main

import (
	"github.com/gin-gonic/gin"
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
		usersGroup.GET("/", UsersHandler)
		usersGroup.GET("/:id", GetUsersByIdHandler)
	}
	router.Run()
}
