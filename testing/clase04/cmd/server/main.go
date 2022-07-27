package main

import (
	"goweb/cmd/server/handler"
	"goweb/internal/users"
	"goweb/pkg/store"
	"log"
	"os"

	"goweb/docs"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Franco Visintini API
// @version 1.0
// @description This API Handle MELI Users. You can get info about the users of the platform, store new users in the database, edit existing users and more!
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

	repo := users.NewRepository(db)
	service := users.NewService(repo)
	u := handler.NewUser(service)

	router := gin.Default()

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	userGroup := router.Group("/users")
	userGroup.GET("/", u.GetAllUsers())
	userGroup.GET("/:id", u.GetUserById())
	userGroup.POST("/", u.StoreUser())
	userGroup.PUT("/:id", u.UpdateTotal())
	userGroup.PATCH("/:id", u.UpdatePartial())
	userGroup.DELETE("/:id", u.Delete())

	if err := router.Run(); err != nil {// si no especifico ningún puerto, por ej ":3001" toma por defecto el 8080
		panic(err)
	}

}
