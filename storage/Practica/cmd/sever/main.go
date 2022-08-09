package main

import (
	"log"
	"os"

	"github.com/bootcamp-go/wave-5-backpack/tree/Ramos_Andres/goweb/practica/cmd/sever/routes"
	"github.com/bootcamp-go/wave-5-backpack/tree/Ramos_Andres/goweb/practica/pkg/db"
	"github.com/bootcamp-go/wave-5-backpack/tree/Ramos_Andres/goweb/practica/pkg/web"
	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
)

func main() {

	err := godotenv.Load("./resources/.env")
	if err != nil {
		log.Fatal("cant load .env file")
	}
	db := db.ConnectDB()

	r := gin.Default()
	router := routes.NewRouter(r, db)
	router.MapRoutes()
	r.Use(TokenAuthMiddleware())

	if err := r.Run(); err != nil {
		log.Fatal(err)
	}
}

func TokenAuthMiddleware() gin.HandlerFunc {
	requiredToken := os.Getenv("TOKEN")

	if requiredToken == "" {
		log.Fatal("no se encontro el token en variable de entorno")
	}

	return func(c *gin.Context) {
		token := c.GetHeader("token")

		if token == "" {
			c.AbortWithStatusJSON(401, web.NewRespose(401, nil, "falta token en cabecera"))
			return
		}

		if token != requiredToken {
			c.AbortWithStatusJSON(401, web.NewRespose(401, nil, "token incorrecto"))
			return
		}

		c.Next()
	}

}
