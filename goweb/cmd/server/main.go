package main

import (
	"os"

	"github.com/bootcamp-go/wave-5-backpack/tree/flood_patricio/goweb/cmd/server/handler"
	"github.com/bootcamp-go/wave-5-backpack/tree/flood_patricio/goweb/cmd/server/middleware"
	"github.com/bootcamp-go/wave-5-backpack/tree/flood_patricio/goweb/docs"
	"github.com/bootcamp-go/wave-5-backpack/tree/flood_patricio/goweb/internal/users"
	"github.com/bootcamp-go/wave-5-backpack/tree/flood_patricio/goweb/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title MELI Bootcamp API
// @version 1.0
// @description This API Handle MELI Users.
// @contact.name API Support
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	_ = godotenv.Load()
	store := store.NewStore(os.Getenv("DB_FILENAME"))
	repo := users.NewRepository(store)
	service := users.NewService(repo)
	u := handler.NewUser(service)

	router := gin.Default()

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	userRouter := router.Group("/users")
	{
		userRouter.GET("/", u.GetAll)
		userRouter.GET("/:Id", u.GetById)

		userRouter.Use(middleware.Authorization)
		userRouter.POST("/", u.Store)
		userRouter.PUT("/:Id", u.Update)
		userRouter.PATCH("/:Id", u.UpdateAgeLastName)
		userRouter.DELETE("/:Id", u.Delete)
	}
	router.Run()
}
