package main

import (
	"database/sql"
	"log"
	"os"
	"proyecto_meli/cmd/server/handler"
	"proyecto_meli/internal/products"
	"proyecto_meli/pkg/web"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func tokenMiddleware() gin.HandlerFunc {
	requiredToken := os.Getenv("TOKEN")

	if requiredToken == "" {
		log.Fatal("No se encontro el token en las variables de entorno")
	}
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token == "" {
			c.AbortWithStatusJSON(401, web.NewResponse(401, nil, "Falta el token en la cabecera"))
			return
		}
		if token != requiredToken {
			c.AbortWithStatusJSON(401, web.NewResponse(401, nil, "Token inválido"))
			return
		}
		c.Next()
	}
}

// @title MELI Bootcamp API
// @version 1.0
// @description This API Handle MELI Products.
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
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/storage")
	if err != nil {
		log.Fatal("ERROR AL INICIAR CONEXION MYSQL")
	}
	if err = db.Ping(); err != nil {
		log.Fatal("ERROR PING")
	}

	repo := products.NewRepository(db)
	service := products.NewService(repo)
	p := handler.NewProduct(service)

	r := gin.Default()

	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	pr := r.Group("/products")
	pr.Use(tokenMiddleware())
	pr.POST("/", p.Store())
	pr.GET("/", p.GetProductByName())
	// pr.GET("/", p.GetAll())
	// pr.GET("/:id", p.GetById())
	// pr.GET("/filter", p.FilterList())
	// pr.PUT("/:id", p.Update())
	// pr.DELETE("/:id", p.Delete())
	// pr.PATCH("/:id", p.Update_Name_Price())
	if errServer := r.Run(); errServer != nil {
		log.Fatal("error al intentar iniciar el servidor")
	}
}
