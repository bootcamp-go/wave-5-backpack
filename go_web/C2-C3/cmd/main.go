package main

import (
	"log"
	"os"

	"C2-C3/cmd/handler"

	"C2-C3/internal/repositorio"
	"C2-C3/internal/servicio"
	"C2-C3/pkg/store"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	// gin-swagger middleware

	"C2-C3/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title MELI Bootcamp API
// @version 1.0
// @description This API Handle MELI Products.
// @termsOfService https://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones
// @contact.name API Support
// @contact.url https://developers.mercadolibre.com.ar/support
// @contact.email rodrigo.vazquez@mercadolibre.com.mx
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host      localhost:8080
// @BasePath  /api/v1

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("error al cargar el archivo .env")
	}
	db := store.NewStore(".json", "../users.json")
	repo := repositorio.NewRepository(db)
	service := servicio.NewService(repo)
	u := handler.NewUser(service)
	r := gin.Default()

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	pr := r.Group("/users")
	pr.POST("/", u.CreateUser())
	pr.GET("/", u.GetAll())
	pr.PUT("/:id", u.UpdateUser())
	pr.PATCH("/:id", u.UpdateLastNameAge())
	pr.DELETE("/:id", u.Delete())
	r.Run()
}
