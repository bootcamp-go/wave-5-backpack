package main

import (
	"goweb/4/tm/pkg/web"
	"goweb/4/tt/cmd/server/handler"
	"goweb/4/tt/docs"
	"goweb/4/tt/internal/repository"
	"goweb/4/tt/internal/service"
	"goweb/4/tt/pkg/store"
	"log"
	"net/http"
	"os"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func TokenAuthMiddleware() gin.HandlerFunc {
	requiredToken := os.Getenv("TOKEN")

	if requiredToken == "" {
		log.Fatal("No se encontro TOKEN en variables de entorno")
	}

	return func(c *gin.Context) {
		token := c.GetHeader("token")

		if token == "" {
			code := http.StatusUnauthorized
			c.AbortWithStatusJSON(code, web.NewResponse(code, nil, "API token requerido"))
			return
		}

		if token != requiredToken {
			code := http.StatusUnauthorized
			c.AbortWithStatusJSON(code, web.NewResponse(code, nil, "API token invalido"))
			return
		}

		c.Next()
	}
}

// @title MELI Bootcamp API
// @version 1.0
// @description This API Handle MELI Products.
// @thermsOfService https://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones

// @contact.name API Support
// @contact.url https://developers.mercadolibre.com.ar/support

// @licence.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error al intentar cargar .env", err)
	}

	fs := store.NewStore(os.Getenv("FILE_PATH"))
	pr := repository.NewRepository(fs)
	ps := service.NewService(pr)
	ph := handler.NewHandler(ps)

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	products := router.Group("/products")
	products.Use(TokenAuthMiddleware())
	{
		products.POST("", ph.Create())

		products.GET("", ph.ReadAll())
		products.GET("/:id", ph.Read())

		products.PUT("/:id", ph.Update())

		products.PATCH("/:id", ph.UpdateNamePrice())

		products.DELETE("/:id", ph.Delete())
	}

	router.Run()
}
