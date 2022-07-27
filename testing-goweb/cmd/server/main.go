package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"log"
	"os"
	"testing-goweb/cmd/server/handler"
	"testing-goweb/docs"
	"testing-goweb/internal/products"
	"testing-goweb/pkg/store"
	"testing-goweb/pkg/web"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error al cargar archvo .env")
	}
	db := store.NewStore("products.json")
	if err != nil {
		log.Fatal("error al cargar archivo json")
	}

	repo := products.NewRepository(db)
	service := products.NewService(repo)
	products := handler.NewProduct(service)

	router := gin.Default()

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	pr := router.Group("/products")
	pr.Use(TokenAuthMiddleWare())
	pr.GET("/", products.GetAll())
	pr.POST("/", products.Store())
	pr.PUT("/:id", products.Update())
	pr.PATCH("/:id", products.UpdatePrecioStock())
	pr.DELETE("/:id", products.Delete())
	if err := router.Run(); err != nil {
		fmt.Println("Error al ejecutar el servidor")
	}

}

func TokenAuthMiddleWare() gin.HandlerFunc {
	requiredToken := os.Getenv("TOKEN")
	if requiredToken == "" {
		log.Fatal(" no se encontró variable de entorno")
	}
	return func(context *gin.Context) {
		token := context.GetHeader("token")
		if token == "" {
			context.AbortWithStatusJSON(401, web.NewResponse(401, nil, "missing token"))
			return
		}

		if token != requiredToken {
			context.AbortWithStatusJSON(401, web.NewResponse(401, nil, "token invalido"))
			return
		}
		context.Next()
	}
}
