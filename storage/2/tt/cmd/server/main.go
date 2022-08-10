package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"storage/2/tt/cmd/server/handler"
	"storage/2/tt/internal/product"
	"storage/2/tt/pkg/web"

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

	db, err := sql.Open("mysql", os.Getenv("DB_CONN"))
	if err != nil {
		log.Fatal(err)
	}

	pr := product.NewRepository(db)
	ps := product.NewService(pr)
	ph := handler.NewProduct(ps)

	router := gin.Default()

	products := router.Group("/products")
	products.Use(TokenAuthMiddleware())
	{
		products.POST("", ph.Create())

		products.GET("", ph.ReadAll())
		products.GET("/:id", ph.Read())
		products.GET("/name/:name", ph.ReadByName())

		products.PUT("/:id", ph.Update())

		products.DELETE("/:id", ph.Delete())
	}

	router.Run()
}
