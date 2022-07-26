package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/swaggo/swag/example/basic/docs"
	"log"
	"net/http"
	"os"
	"proyectoFinal/cmd/server/handler"
	"proyectoFinal/internal/users"
	"proyectoFinal/pkg/store"
	"proyectoFinal/pkg/web"
)

func AuthMiddleware() gin.HandlerFunc {
	token := os.Getenv("TOKEN")
	if token == "" {
		log.Fatal("error: Token was not found as enviroment variable")
	}

	return func(ctx *gin.Context) {
		tokenRequest := ctx.GetHeader("Authorization")
		if token != tokenRequest {
			ctx.AbortWithStatusJSON(web.NewResponse(http.StatusUnauthorized, nil, "error: Unauthorized"))
		}

		ctx.Next()
	}
}

func main() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("error: ", err.Error())
	}

	db := store.NewStore("users.json")

	if err := db.Ping(); err != nil {
		fmt.Println(err.Error())
		return
	}

	repository := users.NewRepository(db)
	service := users.NewService(repository)
	users := handler.NewUser(service)

	docs.SwaggerInfo.Host = os.Getenv("HOST")

	router := gin.Default()

	useresRoute := router.Group("/users")
	useresRoute.Use()
	useresRoute.GET("/", users.GetAll())
	useresRoute.GET("/:id", users.GetById())
	useresRoute.POST("/", users.Create())
	useresRoute.PUT("/:id", users.Update())
	useresRoute.DELETE("/:id", users.Delete())
	router.Run()
}
