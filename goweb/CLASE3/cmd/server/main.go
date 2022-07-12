package main

import (
	"CLASE3/cmd/server/handler"
	"CLASE3/internal/users"
	"CLASE3/pkg/store"
	"CLASE3/pkg/web"
	"log"
	"os"

	"CLASE3/docs"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title MELI Bootcamp API
// @version 1.0
// @description This API Handle MELI users.
// @termsOfService https://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones

// @contact.name API Support
// @contact.url https://developers.mercadolibre.com.ar/support

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error al intentar cargar archivo .env")
	}

	db := store.NewStore("users.json")
	if err := db.Val(); err != nil {
		log.Fatal("Error al intentar cargar archivo")
	}

	repo := users.NewRepository(db)
	s := users.NewService(repo)
	p := handler.NewUsers(s)

	r := gin.Default()
	pr := r.Group("users")

	// Swagger Documentation - Endpoint
	docs.SwaggerInfo.Host = os.Getenv("HOST")
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// <--------------------------------------------------------------->
	pr.Use(TokenAuthMiddleware())
	pr.GET("/", p.GetAll())
	pr.POST("/", p.Store())
	pr.PUT("/:id", p.Update())
	pr.PATCH("/:id", p.UpdateApellidoAndEdad())
	pr.DELETE("/:id", p.Delete())

	r.Run()
}

func TokenAuthMiddleware() gin.HandlerFunc {
	requiredToken := os.Getenv("TOKEN")

	if requiredToken == "" {
		log.Fatal("No se encontro el Token")
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
