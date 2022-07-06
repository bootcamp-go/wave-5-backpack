package main

import (
	"goweb/services"

	"github.com/gin-gonic/gin"
)

func GetAll(c *gin.Context) {
	users, err := services.Read()
	if err != nil {
		c.JSON(500, gin.H{"Error:": err.Error()})
		return
	}
	c.JSON(200, users)
}

func UserById(ctx *gin.Context) {
	users, err := services.Read()
	id := ctx.Param("id")

	if err != nil {
		ctx.JSON(500, gin.H{"Error:": err.Error()})
		return
	}

	for _, user := range users {
		if user.Id == id {
			ctx.JSON(200, user)
			return
		}
	}
	ctx.JSON(404, gin.H{"Error:": "User not found"})
}

func main() {
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hola" + " " + c.Query("nombre"),
		})
	})

	user := router.Group("/users")
	{
		user.GET("/", GetAll)
		user.GET("/:id", UserById)
	}

	//puerto 8080
	router.Run()
}
