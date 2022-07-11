package main

import (
	"log"
	"os"

	"github.com/bootcamp-go/wave-5-backpack/docs"
	"github.com/bootcamp-go/wave-5-backpack/goweb/cmd/server/handler"
	"github.com/bootcamp-go/wave-5-backpack/goweb/internal/users"
	"github.com/bootcamp-go/wave-5-backpack/goweb/pkg/store"
	"github.com/bootcamp-go/wave-5-backpack/goweb/pkg/web"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title MELI Bootcamp API
// @version 1.0
// @description This API Handle MELI Users.
// @termsOfService https://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones
// @contact.name API Support
// @contact.url https://developers.mercadolibre.com.ar/support
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	db := store.NewStore("../../users.json")
	repo := users.NewRepository(db)
	service := users.NewService(repo)
	userHandler := handler.NewUser(service)

	error := godotenv.Load("../../../.env")
	if error != nil {
		log.Fatal("error al intentar cargar el .env")
	}

	router := gin.Default()

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	usersGroup := router.Group("/users")
	usersGroup.Use(tokenAuthMiddleware())
	usersGroup.POST("/", userHandler.Store())
	usersGroup.GET("/", userHandler.GetAll())
	usersGroup.PUT("/:id", userHandler.Update())
	usersGroup.PATCH("/:id", userHandler.UpdateLastNameAndAge())
	usersGroup.DELETE("/:id", userHandler.Delete())
	router.Run()
}

func tokenAuthMiddleware() gin.HandlerFunc {
	requiredToken := os.Getenv("TOKEN")
	if requiredToken == "" {
		log.Fatal("no se encontro el token en la variable de entorno")
	}
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token == "" {
			ctx.AbortWithStatusJSON(401, web.NewResponse(401, nil, "falta token en cabecera"))
			return
		}
		if token != requiredToken {
			ctx.AbortWithStatusJSON(401, web.NewResponse(401, nil, "token incorrecto"))
			return
		}
		ctx.Next()
	}
}
