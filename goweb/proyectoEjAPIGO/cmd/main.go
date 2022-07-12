package main

import (
	"goweb/cmd/server/handler"
	"goweb/internal/transacciones"
	"goweb/pkg/store"
	"goweb/pkg/web"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	// server.GET("/", HandlerRaiz)

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error al intentar cargar el archivo.env")
	}

	db := store.NewStore("/Users/joaalvarez/Desktop/wave-5-backpack/goweb/proyectoEjAPIGO/transacciones.json")
	if err := db.Ping(); err != nil {
		log.Fatal("Error al intentar cargar el archivo")
	}

	repo := transacciones.NewRepository(db)
	s := transacciones.NewService(repo)
	t := handler.NewProduct(s)

	server := gin.Default()

	transacciones := server.Group("/transacciones")

	transacciones.Use(TokenAuthMiddleware())
	{

		transacciones.GET("/", t.GetAll())
		transacciones.POST("/", t.Store())
		transacciones.PUT("/:id", t.Update())
		// transacciones.POST("/", AgregarTransacciones())
	}

	server.Run(":8006")
}

func TokenAuthMiddleware() gin.HandlerFunc {
	requiredToken := os.Getenv("TOKEN")

	if requiredToken == "" {
		log.Fatal("no se encontro el token en variable de entorno")
	}

	return func(c *gin.Context) {
		token := c.GetHeader("token")

		if token == "" {
			c.AbortWithStatusJSON(401, web.NewResponse(401, nil, "falta token en cabecera"))
			return
		}

		if token != requiredToken {
			c.AbortWithStatusJSON(401, web.NewResponse(401, nil, "token incorrecto"))
			return
		}

		c.Next()
	}

}
