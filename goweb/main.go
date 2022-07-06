package main

import (
	"github.com/bootcamp-go/wave-5-backpack/tree/flood_patricio/goweb/internal/users/usercontroller"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	userRouter := router.Group("/users")
	{
		userRouter.GET("/", usercontroller.GetAll)
		userRouter.GET("/:Id", usercontroller.GetById)
	}
	router.Run()
}
