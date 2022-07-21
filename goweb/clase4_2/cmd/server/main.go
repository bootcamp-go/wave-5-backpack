package main

import (
	"clase2_2/cmd/server/handler"
	"clase2_2/internal/users"
	"clase2_2/pkg/storage"

	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/ncostamagna/meli-bootcamp/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title MELI Bootcamp API
// @version 1.0
// @description This API Handle MELI Users.
// @termsOfService https://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones

// @contact.name API Support
// @contact.url https://developers.mercadolibre.com.ar/support

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	err := godotenv.Load()
	if err != nil {
		panic("error al obtener las variables de ambiente ")
	}
	db := storage.NewStore("usuarios.json")
	repository := users.NewRepository(db)
	service := users.NewService(repository)
	handler := handler.NewUser(service)
	router := gin.Default()

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//clase 2_2
	us := router.Group("users")
	us.GET("/", handler.GetAll())
	us.POST("/", handler.AddUser())
	//clase 3_1
	us.PUT("/:id", handler.UpdateUser())
	us.DELETE("/:id", handler.Delete())
	us.PATCH("/:id", handler.UpdateUserName())
	router.Run()
}
