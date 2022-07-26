package main

import (
	"os"

	"github.com/bootcamp-go/wave-5-backpack/tree/gonzalez_jose/goweb/cmd/server/handler"
	"github.com/bootcamp-go/wave-5-backpack/tree/gonzalez_jose/goweb/docs"
	"github.com/bootcamp-go/wave-5-backpack/tree/gonzalez_jose/goweb/internal/users"
	"github.com/bootcamp-go/wave-5-backpack/tree/gonzalez_jose/goweb/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Meli Bootcamp API Go
// @version 1.0
// @description This API Handle Wave5Go Endpoints
// @termsOfService

// @contact.name JofeGonzalezMeLi
// @contact.url

//@license.name
// @license.url
func main() {

	_ = godotenv.Load()
	db := store.NewStore("./usuarios.json")

	repository := users.NewRepository(db)
	service := users.NewService(repository)
	u := handler.NewUser(service)
	router := gin.Default()

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	ug := router.Group("/users")
	ug.POST("/", u.Store())
	ug.GET("/", u.GetAll())
	ug.PUT("/:id", u.Update())
	ug.DELETE("/:id", u.Delete())
	ug.PATCH("/:id", u.UpdateLastNameAndAge())
	router.Run()
}
