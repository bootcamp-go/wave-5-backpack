package main

import (
	"desafio-go-web/cmd/server/handler"
	"desafio-go-web/docs"
	"desafio-go-web/internal/tickets"
	"desafio-go-web/pkg/store"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Desafio Go-Web API MACN
// @version 1.0
// @description Api pade gestión de tickets
// @termsOfService https://go.dev/tos

// @contact.me SoporteMACN
// @contact.url https:&&go.dev/blog&

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error al intentar cargar el archivo de configuración")
	}

	db := store.NewStore("./tickets.csv")
	if err := db.Ping(); err != nil {
		log.Fatal("error al intentar leer el csv")
	}

	// Creamos el repositorio, el servicio y el handler
	repo := tickets.NewRepository(db)
	service := tickets.NewService(repo)
	hand := handler.NewService(service)

	// Definimos el router de la aplicación
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) { c.String(200, "pong") })

	// Se deshabilita la advertencia de que la api confia en todos los proxies
	router.SetTrustedProxies(nil)

	// Se define swagger
	docs.SwaggerInfo.Host = os.Getenv("HOST")
	router.GET("docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Definimos el middleware de seguridad
	router.Use(handler.TokenAuthMiddleware())

	// Definimos los endpoint de tickets
	pr := router.Group("/ticket")
	{
		pr.GET("/getByCountry/:dest", hand.GetTicketsByCountry())
		pr.GET("/getAverage/:dest", hand.AverageDestination())
	}

	if err := router.Run(); err != nil {
		panic(err)
	}
}
