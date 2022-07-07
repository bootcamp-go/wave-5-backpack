package main

import (
	"github.com/bootcamp-go/wave-5-backpack/tree/flood_patricio/goweb/internal/middleware"
	"github.com/bootcamp-go/wave-5-backpack/tree/flood_patricio/goweb/internal/user/usercontroller"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	userRouter := router.Group("/users")
	{
		userRouter.GET("/", usercontroller.GetAll)
		userRouter.POST("", middleware.Authorization, usercontroller.Create)
		userRouter.GET("/:Id", usercontroller.GetById)
	}
	router.Run()
}
