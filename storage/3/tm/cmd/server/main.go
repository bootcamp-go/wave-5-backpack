package main

import (
	"log"
	"net/http"
	"os"
	"storage/3/tm/cmd/server/handler"
	"storage/3/tm/internal/product"
	"storage/3/tm/pkg/web"
	"storage/3/tm/util"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func TokenAuthMiddleware() gin.HandlerFunc {
	requiredToken := os.Getenv("TOKEN")

	if requiredToken == "" {
		log.Fatal("No se encontro TOKEN en variables de entorno")
	}

	return func(c *gin.Context) {
		token := c.GetHeader("token")

		if token == "" {
			code := http.StatusUnauthorized
			c.AbortWithStatusJSON(code, web.NewResponse(code, nil, "API token requerido"))
			return
		}

		if token != requiredToken {
			code := http.StatusUnauthorized
			c.AbortWithStatusJSON(code, web.NewResponse(code, nil, "API token invalido"))
			return
		}

		c.Next()
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error al intentar cargar .env", err)
	}

	db, err := util.InitDynamo()
	if err != nil {
		log.Fatal(err)
	}

	pr := product.NewDynamoRepository(db, "products")
	ps := product.NewService(pr)
	ph := handler.NewProduct(ps)

	router := gin.Default()

	products := router.Group("/products")
	products.Use(TokenAuthMiddleware())
	{
		products.POST("", ph.Create())

		products.GET("", ph.ReadAll())
		products.GET("/:id", ph.Read())

		products.PUT("/:id", ph.Update())

		products.DELETE("/:id", ph.Delete())
	}

	router.Run()
}
