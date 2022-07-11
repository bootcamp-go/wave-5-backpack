package main

import (
	"goweb/cmd/handler"
	"goweb/docs"
	"goweb/internals/transactions"
	"goweb/pkg/store"
	"goweb/pkg/web"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//@title MELI Bootcamp API Test
//@version 1.0
//@description This API is an excercise

//@contact.name API Support
//@contact.url https://www.twitter.com/imakheri

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ocurrió un error al leer el archivo .env")
	}

	db := store.NewStore("transacciones.json")
	if err := db.Ping(); err != nil {
		log.Fatal("error al intentar cargar archivo")
	}

	repo := transactions.NewRepository(db)
	service := transactions.NewService(repo)
	h := handler.NewTransaction(service)

	router := gin.Default()

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	transacciones := router.Group("transacciones")
	transacciones.GET("/", h.GetAll())
	transacciones.GET("/filtros", h.GetByQuery())
	transacciones.GET("/:id", h.GetByID())
	transacciones.POST("/", TokenAuthMiddleware(), h.Store())
	transacciones.PUT("/", TokenAuthMiddleware(), h.Update())
	transacciones.DELETE("/", TokenAuthMiddleware(), h.Delete())

	router.Run(":8080")
}

func TokenAuthMiddleware() gin.HandlerFunc {
	requiredToken := os.Getenv("TOKEN")

	if requiredToken == "" {
		log.Fatal("No se encontró el token en la variable de entorno")
	}
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token == "" {
			c.AbortWithStatusJSON(401, web.NewResponse(401, nil, "Falta token en la cabecera"))
			return
		}
		if token != requiredToken {
			c.AbortWithStatusJSON(401, web.NewResponse(401, nil, "El token es incorrecto"))
			return
		}
		c.Next()
	}
}
