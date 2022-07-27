package main

import (
	"log"
	"os"

	"github.com/bootcamp-go/wave-5-backpack/goweb/cmd/server/handler"
	"github.com/bootcamp-go/wave-5-backpack/goweb/docs"
	"github.com/bootcamp-go/wave-5-backpack/goweb/internal/users"
	"github.com/bootcamp-go/wave-5-backpack/goweb/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwaggwer "github.com/swaggo/gin-swagger"
)

// @title API Usuarios Bootcamp MACN
// @version 1.0
// @description Api para gestionar usuarios
// @termsOfService https://go.dev/tos

// @contact.name SoporteMACN
// @contact.url https://go.dev/blog/

// @license.name Apache 2.0
// @license.url http://www.apache.orh/licenses/LICENSE-2.0.html
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error al intentar cargar el archivo de configuración")
	}

	db := store.NewStore("./usuarios.json")
	if err := db.Ping(); err != nil {
		log.Fatal("error al intentar conectarse a la db")
	}

	// Creamos el repositorio, el servicio y el handler
	repo := users.NewRepository(db)
	service := users.NewService(repo)
	hand := handler.NewUser(service)

	// Definimos el router de la aplicación
	router := gin.Default()

	// Se deshabilita la advertencia de que la api confia en todos los proxies
	errProxies := router.SetTrustedProxies(nil)
	if errProxies != nil {
		panic(err)
	}

	// Definimos Swagger
	docs.SwaggerInfo.Host = os.Getenv("HOST")
	router.GET("/docs/*any", ginSwaggwer.WrapHandler((swaggerFiles.Handler)))

	// Definimos el middleware de seguridad
	router.Use(handler.TokenAuthMiddleware())

	// Definimos los endpoint de Usuarios
	pr := router.Group("/users")
	{
		pr.GET("/", hand.GetAll())
		pr.GET("/:id", hand.GetById())
		pr.GET("/search", hand.SearchUser())
		pr.POST("/", hand.Store())
		pr.PUT("/:id", hand.Update())
		pr.PATCH("/:id", hand.UpdateApellidoEdad())
		pr.DELETE("/:id", hand.Delete())
	}

	errRun := router.Run(":8080")
	if errRun != nil {
		panic(errRun)
	}
}
