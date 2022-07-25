package main

import (
	"GoWeb/cmd/server/handler"
	"GoWeb/internals/transactions"
	"GoWeb/pkg/store"
	"GoWeb/pkg/web"
	"log"
	"os"

	"GoWeb/docs"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//@title MELI Bootcamp API Juan David Serna
//@version 1.0
//@description This API Handle MELI Transactions
//@termsOfService https://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones

//@contact.name API Support
//@contact.url https://developers.mercadolibre.com.ar/support

func TokenAuthMiddleware() gin.HandlerFunc {
	requiredToken := os.Getenv("TOKEN")

	if requiredToken == "" {
		log.Fatal("no se encontro el token en la variable de entorno")
	}
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")

		if token == "" {
			ctx.AbortWithStatusJSON(401, web.NewResponse(401, nil, "falta token en cabecera"))
			return
		}

		if token != requiredToken {
			ctx.AbortWithStatusJSON(401, web.NewResponse(401, nil, "token incorrecto"))
		}
		ctx.Next()
	}
}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("error al intentar cargar el archivo .env")
	}

	db := store.NewStore("transactions.json")
	if err := db.Ping(); err != nil {
		log.Fatal("error al intentar cargar el archivo")
	}

	repo := transactions.NewRepository(db)
	service := transactions.NewService(repo)
	tran := handler.NewTransaction(service)

	router := gin.Default()
	docs.SwaggerInfo.Host = os.Getenv("HOST")
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	tr := router.Group("/transacciones")
	tr.Use(TokenAuthMiddleware())
	tr.POST("/", tran.Store())
	tr.GET("/", tran.GetAll())
	tr.PUT("/:id", tran.Update())
	tr.DELETE("/:id", tran.Delete())
	tr.PATCH("/:id", tran.UpdateCode())
	router.Run()
}
