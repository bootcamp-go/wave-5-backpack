package main

import (
	"log"
	"net/http"
	"os"

	"github.com/bootcamp-go/wave-5-backpack/goweb/cmd/handler"
	"github.com/bootcamp-go/wave-5-backpack/goweb/docs"
	"github.com/bootcamp-go/wave-5-backpack/goweb/internal/products"
	"github.com/bootcamp-go/wave-5-backpack/goweb/pkg/store"
	"github.com/bootcamp-go/wave-5-backpack/goweb/pkg/web"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title MELI Bootcamp API
// @version 1.0
// @description This Api Handle MELI Products.
// @termsOfService https://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones

// @contact.name API Support
// @contact.url https://developers.mercadolibre.com.ar/support

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /products
func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("error al intentar cargar archivo .env")
	}

	db := store.NewStore("products.json")
	if err := db.Ping(); err != nil {
		log.Fatal("error al intentar cargar el archivo .json")
	}

	repository := products.NewRepository(db)
	service := products.NewService(repository)
	product := handler.NewProduct(service)

	router := gin.Default()

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	group := router.Group("products")

	group.Use(TokenValidator())

	group.GET("/:id", product.GetProduct())
	group.GET("/", product.GetAll())
	group.POST("/", product.Store())
	group.PUT("/:id", product.UpdateAll())
	group.PATCH("/:id", product.Update())
	group.DELETE("/:id", product.Delete())

	//group.GET("/", GetFilter)

	if err := router.Run(); err != nil {
		panic(err)
	}

}

func TokenValidator() gin.HandlerFunc {
	getToken := os.Getenv("TOKEN")
	if getToken == "" {
		log.Fatal("No se encontro el token en variable de entorno")
	}

	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")

		if token == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, web.NewResponse(401, nil, "Falta token en cabecera"))
			return
		}

		if token != getToken {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, web.NewResponse(401, nil, "Token incorrecto"))
			return
		}
		ctx.Next()
	}
}
