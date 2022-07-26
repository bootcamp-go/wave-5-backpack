package main

import (
	"log"
	"os"

	"github.com/bootcamp-go/wave-5-backpack/cmd/server/handler"
	"github.com/bootcamp-go/wave-5-backpack/docs"
	"github.com/bootcamp-go/wave-5-backpack/internal/users"
	"github.com/bootcamp-go/wave-5-backpack/pkg/store"
	"github.com/bootcamp-go/wave-5-backpack/pkg/web"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
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

	if err := godotenv.Load(); err != nil {
		log.Fatal("error archivo .env")
	}
	db := store.NewStore("users.json")
	if err := db.Ping(); err != nil {
		log.Fatal("error al intentar cargar archivo")
	}
	repo := users.NewRepository(db)
	service := users.NewService(repo)
	u := handler.NewUser(service)

	server := gin.Default()

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	server.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	users := server.Group("/users")

	users.Use(TokenAuthMiddleware())
	users.GET("/", u.GetAll())
	users.GET("/:id", u.GetById())
	users.POST("/", u.StoreUser())
	users.PATCH("/:id", u.UpdateLastnameAndAge())
	users.PUT("/:id", u.UpdateUser())
	users.DELETE("/:id", u.DeleteUser())

	server.Run(":8080")
}

func TokenAuthMiddleware() gin.HandlerFunc {
	requiredToken := os.Getenv("TOKEN")
	if requiredToken == "" {
		log.Fatal("no se encontró el token en variable de entorno")
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
