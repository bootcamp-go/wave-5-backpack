package main

import (
	"clase4_parte2/cmd/server/handler"
	"clase4_parte2/docs"
	"clase4_parte2/internal/products"
	"clase4_parte2/pkg/store"
	"clase4_parte2/pkg/store/web"
	"log"
	"os"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// @title MELI Bootcamp API
// @version 1.0
// @description This API Handle MELI Products.
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
	db := store.NewStore("products.json")
	if err := db.Ping(); err != nil {
		log.Fatal("error al intentar cargar archivo")
	}

	repo := products.NewRepository(db)
	s := products.NewService(repo)
	p := handler.NewProduct(s)

	r := gin.Default()

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	pr := r.Group("products")
	pr.Use(TokenAuthMiddleware())
	pr.GET("/", p.GetAll())
	pr.POST("/", p.Store())
	pr.PUT("/:id", p.Update())
	pr.PATCH("/:id", p.UpdateName())
	pr.DELETE("/:id", p.Delete())

	r.Run()
}

func TokenAuthMiddleware() gin.HandlerFunc {
	requiredToken := os.Getenv("TOKEN")

	if requiredToken == "" {
		log.Fatal("No se encontro el token en variable de entorno")
	}

	return func(c *gin.Context) {
		token := c.GetHeader("token")

		if token == "" {
			c.AbortWithStatusJSON(401, web.NewResponse(401, nil, "Falta token en cabecera"))
			return
		}

		if token != requiredToken {
			c.AbortWithStatusJSON(401, web.NewResponse(401, nil, "Token incorrecto"))
			return
		}
		c.Next()
	}

}
