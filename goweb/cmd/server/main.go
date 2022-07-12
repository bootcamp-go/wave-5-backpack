package main

import (
	"log"
	"os"
	"web-server/cmd/server/handler"
	"web-server/docs"
	"web-server/internal/products"
	"web-server/pkg/store"
	"web-server/pkg/web"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Meli Bootcamp API
// @version 1.0
// @description This API Handle MELI Products
// @termOfService https://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones

// @contact.name API support
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
	service := products.NewService(repo)
	p := handler.NewProduct(service)

	//Creo un servidor web con 2 middlewares por defecto: logger and recovery middleware
	router := gin.Default()
	docs.SwaggerInfo.Host = os.Getenv("HOST")
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	pr := router.Group("products")
	pr.POST("/", TokenAuthMiddleware(), p.Store())
	pr.GET("/", TokenAuthMiddleware(), p.GetAll())
	pr.PUT("/:id", TokenAuthMiddleware(), p.Update())
	pr.PATCH("/name/:id", TokenAuthMiddleware(), p.UpdateName())
	pr.PATCH("/price/:id", TokenAuthMiddleware(), p.UpdatePrice())
	pr.DELETE("/:id", TokenAuthMiddleware(), p.Delete())

	//Creo un handler utilizando la funcion router.GET("endpoint",Handler) donde endpoint es la ruta relativa y handler es la funcion que toma *gin.Context como argumento.
	router.Run()
}

func TokenAuthMiddleware() gin.HandlerFunc {
	requiredToken := os.Getenv("TOKEN")

	if requiredToken == "" {
		log.Fatal("no se encontro el token en variable de entorno")
	}

	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")

		if token == "" {
			ctx.AbortWithStatusJSON(401, web.NewResponse(401, nil, "falta token en cabecera"))
			return
		}

		if token != requiredToken {
			ctx.AbortWithStatusJSON(401, web.NewResponse(401, nil, "token incorrecto"))
			return
		}

		ctx.Next()
	}
}
