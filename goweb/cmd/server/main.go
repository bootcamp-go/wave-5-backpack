package main

import (
	"goweb/cmd/server/handler"
	"goweb/internal/products"
	"goweb/pkg/store"
	"log"
	"os"

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
// @license.url https://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error al intentar cargar el archivo .env")
	}

	db := store.InitStore("products.json")
	if err := db.Ping(); err != nil {
		log.Fatal("Error al intentar cargar el archivo")
	}
	repo := products.InitRepository(db)
	service := products.InitService(repo)
	p := handler.InitProduct(service)

	r := gin.Default()
	pr := r.Group("/products")

	//middleware de autenticacion con token
	pr.Use(TokenAuthMiddleware())

	pr.POST("/create", p.CreateProduct())
	pr.GET("/", p.GetAll())
	pr.PUT("/:id", p.Update())
	pr.DELETE("/:id", p.Delete())
	pr.PATCH("/:id", p.UpdateOne())
	r.Run()
}

func TokenAuthMiddleware() gin.HandlerFunc {
	requiredToken := os.Getenv("TOKEN")

	if requiredToken == "" {
		log.Fatal("No se encontro el token")
	}

	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")

		if token == "" {
			ctx.AbortWithStatusJSON(401, gin.H{
				"response": "Falta token en la cabecera",
			})
			return
		}
		if token != requiredToken {
			ctx.AbortWithStatusJSON(401, gin.H{
				"response": "Token incorrecto",
			})
			return
		}

		ctx.Next()
	}

}
