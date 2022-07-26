package main

import (
	"log"
	"os"

	"github.com/bootcamp-go/wave-5-backpack/tree/ospina_christian/testing/c2_testingTM/cmd/server/handler"
	"github.com/bootcamp-go/wave-5-backpack/tree/ospina_christian/testing/c2_testingTM/docs"
	"github.com/bootcamp-go/wave-5-backpack/tree/ospina_christian/testing/c2_testingTM/internal/productos"
	"github.com/bootcamp-go/wave-5-backpack/tree/ospina_christian/testing/c2_testingTM/pkg/store"
	"github.com/bootcamp-go/wave-5-backpack/tree/ospina_christian/testing/c2_testingTM/pkg/web"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//@title MELI Bootcamp API Christian
//@version 1.0
//@description This is API of meli products
//@termsOfService https://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones

//@contact.name API Support
//@contact.url https://developers.mercadolibre.com.ar/support

//@license.name Apache 2.0
//@license.url http://www.apache.org/licenses/LICENSE-2.0.html

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error al intentar cargar el archivo .env")
	}

	db := store.NewStore("productos.json")
	if err := db.Ping(); err != nil {
		log.Fatal("error al intentar cargar el archivo")
	}

	repo := productos.NewRepository(db)
	s := productos.NewService(repo)
	p := handler.NewProduct(s)
	router := gin.Default()

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	router.GET("docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	rProductos := router.Group("productos")
	rProductos.Use(TokenAuthMiddleware())
	rProductos.GET("/", p.GetAll())
	rProductos.GET("/:id", p.GetForId())
	rProductos.POST("/", p.Store())
	rProductos.PUT("/:id", p.Update())
	rProductos.PATCH("/:id", p.UpdatePrecio())
	rProductos.DELETE("/:id", p.Delete())

	if err := router.Run(); err != nil {
		log.Panic(err)
	}

}

func TokenAuthMiddleware() gin.HandlerFunc {
	requiredToken := os.Getenv("TOKEN")
	if requiredToken == "" {
		log.Fatal("No se encontro el token en la variable de entorno")
	}
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token == "" {
			ctx.AbortWithStatusJSON(401, web.NewResponse(401, nil, "Falta el token en la cabecera"))
			return
		}
		if token != requiredToken {
			ctx.AbortWithStatusJSON(401, web.NewResponse(401, nil, "Token Incorrecto"))
			return
		}
		ctx.Next()
	}
}
