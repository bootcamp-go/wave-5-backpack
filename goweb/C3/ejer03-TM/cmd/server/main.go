package main

import (
	"ejer02-TT/cmd/server/handler"

	"ejer02-TT/internal/transactions"
	"ejer02-TT/pkg/store"
	"ejer02-TT/pkg/web"
	"log"
	"os"

	"ejer02-TT/docs"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//@title MELI BOOTCAMP API
//@version 1.0
//@description This API Handle MELI Transactions.
//@TermsOfService https://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones

//@contact.name API Support
//@contact.url https://developers.mercadolibre.com.ar/support

//@license.name Apache 2.0
//@license.url http://www.apache.org/licenses/LICENSE-2.0.html

func main() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("error al cargar archivo .env")
	}

	db := store.NewStore("transacciones.json")
	if err := db.Ping(); err != nil {
		log.Fatal("error al intentar cargar archivo")
	}

	router := gin.Default()

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	repository := transactions.NewRepository(db)
	service := transactions.NewService(repository)
	transactions := handler.NewTransaction(service)

	tr := router.Group("transactions")
	tr.Use(TokenAuthMiddleware())
	tr.GET("/", transactions.GetAll())
	tr.POST("/", transactions.Store())
	tr.PUT("/:id", transactions.Update())
	tr.PATCH("/:id", transactions.UpdateCodeAndAmount())
	tr.DELETE("/:id", transactions.Delete())
	router.Run()

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
