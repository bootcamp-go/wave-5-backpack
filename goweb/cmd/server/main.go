package main

import (
	"os"

	"github.com/bootcamp-go/wave-5-backpack/tree/flood_patricio/goweb/cmd/server/handler"
	"github.com/bootcamp-go/wave-5-backpack/tree/flood_patricio/goweb/cmd/server/middleware"
	"github.com/bootcamp-go/wave-5-backpack/tree/flood_patricio/goweb/internal/users"
	"github.com/bootcamp-go/wave-5-backpack/tree/flood_patricio/goweb/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	store := store.NewStore(os.Getenv("DB_FILENAME"))
	repo := users.NewRepository(store)
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
