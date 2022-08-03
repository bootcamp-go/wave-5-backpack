package main

import (
	"goweb/go-web-II/cmd/handler"
	"goweb/go-web-II/docs"
	"goweb/go-web-II/internal/products"
	"goweb/go-web-II/pkg/store"
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
	_ = godotenv.Load()
	db := store.NewStore("/Users/namonserrat/Desktop/wave-5-backpack/goweb/go-web-II/usuarios.json")
	router := gin.Default()
	repository := products.NewRepository(db)
	service := products.NewService(repository)
	handler := handler.NewUser(service)

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	rUser := router.Group("users")
	/*
		rUser.Use(tokenValidator())
	*/
	rUser.GET("/users", handler.GetAll())
	rUser.POST("/", handler.Store())
	rUser.PUT("/:id", handler.Update())
	rUser.DELETE("/id", handler.Delete())
	_ = router.Run()

}

/*
1.- Lo recibe el handler, le pide a service que haga su trabajo
2.- Service implementa la lógica necesaria para el pedido
3.- Sí es que necesitamos consultar la BD, Service se lo va a delegar
a repository.

4.- Service busca que interceptar si es que hay sucede un error en esa
capa o en repository, para que no llegue al Server
*/
