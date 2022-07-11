/*---------------------------------------------------------*

     Assignment:	Practica #2
         Author:	Israel Fabela
	   Language:	go1.18.3 darwin/arm64
		  Topic:	Go Web

	Description:
		‣	Handling of generic responses

	© Mercado Libre - IT Bootcamp 2022

----------------------------------------------------------*/

package main

import (
	"goweb/clase4-go-web-tt/cmd/server/handler"
	"goweb/clase4-go-web-tt/docs"
	"goweb/clase4-go-web-tt/internal/transactions"
	"goweb/clase4-go-web-tt/pkg/bank"
	"goweb/clase4-go-web-tt/pkg/bank/web"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title MELI Bootcamp API
// @version 1.0
// @description This API Handle MELI Transactions.
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

	db := bank.NewBank("transacciones.json")
	if err := db.Ping(); err != nil {
		log.Fatal("error al intentar cargar archivo")
	}

	repo := transactions.NewRepository(db)
	service := transactions.NewService(repo)
	t := handler.NewTransaction(service)

	r := gin.Default()

	r.GET("/", handler.PaginaPrincipal)
	r.GET("/hola", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hola Israel ! 👋",
		})
	})

	//	SWAGGER  DOCUMENTATION - EndPoint
	docs.SwaggerInfo.Host = os.Getenv("HOST")
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	pr := r.Group("/transactions")
	pr.Use(tokenAuthMiddleware())
	pr.GET("/", t.GetAll())
	pr.GET("/:id", t.GetOne())
	pr.PUT("/:id", t.Update())
	pr.POST("/", t.Ecommerce())
	pr.PATCH("/:id", t.UpdateOne())
	pr.DELETE("/:id", t.Delete())

	r.Run()
}

func tokenAuthMiddleware() gin.HandlerFunc {
	requiredToken := os.Getenv("TOKEN")

	if requiredToken == "" {
		log.Fatal("No se encontro el token en variable de entorno 😞")
	}

	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token == "" {
			ctx.AbortWithStatusJSON(401, web.NewResponse(401, nil, "Falta token en Cabecera 🔐"))
		}

		if token != requiredToken {
			ctx.AbortWithStatusJSON(401, web.NewResponse(401, nil, "Token incorrecto 🔐"))
		}

		ctx.Next()
	}
}
