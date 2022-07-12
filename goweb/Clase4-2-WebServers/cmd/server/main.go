package main

import (
	"goweb/Clase4-2-WebServers/cmd/server/handler"
	"goweb/Clase4-2-WebServers/docs"
	"goweb/Clase4-2-WebServers/internal/usuarios"
	"goweb/Clase4-2-WebServers/pkg/store"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

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
	db := store.NewStore("users.json")
	if err := db.Ping(); err != nil {
		log.Fatal("error al intentar cargar archivo")
	}

	repo := usuarios.NewRepository(db)
	service := usuarios.NewService(repo)
	user := handler.NewUser(service)

	r := gin.Default()
	docs.SwaggerInfo.Host = os.Getenv("HOST")
	r.GET("docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	ur := r.Group("/usuarios")
	ur.POST("/", user.Store())
	ur.GET("/", user.GetAll())
	ur.PUT("/:id", user.Update())
	ur.PATCH("/:id", user.UpdateSurnameAndAge())
	ur.DELETE("/:id", user.Delete())
	r.Run()
}
