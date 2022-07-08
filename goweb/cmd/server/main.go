package main

import (
	"github.com/bootcamp-go/wave-5-backpack/tree/flood_patricio/goweb/cmd/server/handler"
	"github.com/bootcamp-go/wave-5-backpack/tree/flood_patricio/goweb/cmd/server/middleware"
	"github.com/bootcamp-go/wave-5-backpack/tree/flood_patricio/goweb/internal/users"
	"github.com/gin-gonic/gin"
)

func main() {
	repo := users.NewRepository()
	service := users.NewService(repo)
	u := handler.NewUser(service)

	router := gin.Default()
	userRouter := router.Group("/users")
	{
		userRouter.GET("/", u.GetAll)
		userRouter.POST("/", middleware.Authorization, u.Store)
		userRouter.GET("/:Id", u.GetById)
		userRouter.PUT("/:Id", middleware.Authorization, u.Update)
		userRouter.PATCH("/:Id", middleware.Authorization, u.UpdateAgeLastName)
		userRouter.DELETE("/:Id", middleware.Authorization, u.Delete)
	}
	router.Run()
}
