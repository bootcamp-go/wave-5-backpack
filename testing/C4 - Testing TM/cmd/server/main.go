package main

import (
	"C3ejercicioTT/cmd/server/handler"
	users "C3ejercicioTT/internal/users"
	"C3ejercicioTT/pkg/store"
	"C3ejercicioTT/pkg/web"
	"log"
	"os"

	"C3ejercicioTT/docs"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	//"github.com/swaggo/swag/example/basic/docs"
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
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error al intentar cargar archivo .env")
	}

	db := store.New(store.FileType, "users.json")

	repo := users.NewRepository(db)
	serv := users.NewService(repo)
	user := handler.NewUser(serv)

	r := gin.Default()
	us := r.Group("usuarios")

	// Swagger Documentation - Endpoint
	docs.SwaggerInfo.Host = os.Getenv("HOST")
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	us.Use(TokenAuthMiddleware())
	us.GET("/", user.GetAll())
	us.POST("/", user.Store())
	us.PUT("/:id", user.Update())
	us.PATCH("/:id", user.UpdateLastAge())
	us.DELETE("/:id", user.Delete())

	r.Run()
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
