package main

import (
	"goweb/cmd/server/handler"
	"goweb/db"
	"goweb/docs"
	"goweb/internal/users"
	"goweb/pkg/web"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//@title MELI Bootcamp API usuarios
//@version 1.0
//@description esta API maneja la informacion de usuarios

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("error al intentar cargar archivo .env")
	}

	db := db.Connection()

	repository := users.NewRepository(db)
	service := users.NewService(repository)
	u := handler.NewUser(service)

	router := gin.Default()

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Use(validateToken())

	users := router.Group("/users")
	{
		users.GET("/", u.GetAll())
		users.GET("/name/:nombre", u.GetByName())
		users.GET("/:id", u.GetById())

		users.POST("/", u.Store())
		users.PUT("/:id", u.Update())
		users.DELETE("/:id", u.Delete())
		users.PATCH("/:id", u.Patch())

	}

	if err := router.Run(); err != nil {
		return
	}
}

func validateToken() gin.HandlerFunc {
	tokenEnv := os.Getenv("TOKEN")

	if tokenEnv == "" {
		log.Fatal("No se encontró el token en la variable token")
	}

	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")

		if token == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, "El token es obligatorio"))
			return
		}

		if token != tokenEnv {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, "NO tiene permisos para realizar la petición solicitada"))
			return
		}

		ctx.Next()
	}
}
