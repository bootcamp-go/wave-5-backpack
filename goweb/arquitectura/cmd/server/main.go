package main

import (
	"arquitectura/cmd/server/handler"
	"arquitectura/docs"
	"arquitectura/internal/transactions"
	"arquitectura/pkg/store"
	"arquitectura/pkg/web"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title MELI Bootcamp API
// @version 1.0
// @description Documentación API Bootcamp Go
// @termsOfService https://developers.mercadolibre.com.ar/es_ar/terminos_y_condiciones
// @contact.name Cristobal Monsalve
// @contact.url https://developers.mercadolibre.com.ar/support
// @licence.name Apache 2.0
// @licence.url http://www.apache.org/licences/LICENCE-2.0.html
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error al intentar cargar archivo .env")
	}
	db := store.NewStore("transactions.json")
	if err := db.Ping(); err != nil {
		log.Fatal("error al cargar archivo")
	}
	repository := transactions.NewRepository(db)
	service := transactions.NewService(repository)
	transactions := handler.NewTransaction(service)

	router := gin.Default()

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	tr := router.Group("/transactions")
	tr.Use(TokenAuthMiddleWare())
	tr.POST("/", transactions.Store())
	tr.GET("/", transactions.GetAll())
	tr.PUT("/:id", transactions.Update())
	tr.DELETE("/:id", transactions.Delete())
	tr.PATCH("/:id", transactions.UpdateFields())
	router.Run()
}

func TokenAuthMiddleWare() gin.HandlerFunc {
	requiredToken := os.Getenv("TOKEN")
	if requiredToken == "" {
		log.Fatal("no se encontro el token en variable de entorno")
	}

	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token == "" {
			c.AbortWithStatusJSON(401, web.NewResponse(401, nil, "missing token"))
			return
		}

		if token != requiredToken {
			c.AbortWithStatusJSON(401, web.NewResponse(401, nil, "invalid token"))
			return
		}
		c.Next()
	}
}
