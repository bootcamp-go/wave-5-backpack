package main

import (
	"log"
	"os"

	"github.com/del_rio/web-server/cmd/server/controlador"
	"github.com/del_rio/web-server/docs"
	"github.com/del_rio/web-server/internal/usuarios"
	"github.com/del_rio/web-server/pkg/store"
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
		log.Fatal("error al cargar el archivo ", err)
	}
	fileStore := store.NewStore("usuarios.json")
	repo := usuarios.NewRepository(fileStore)

	servicio := usuarios.NewService(repo)
	controladorUsuarios := controlador.NewControlador(servicio)
	router := gin.Default()
	router.Use(controlador.TokenAuthMiddleware())

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	usuarioRoute := router.Group("/usuarios")
	router.Use(controlador.TokenAuthMiddleware())
	usuarioRoute.GET("/", controladorUsuarios.VerUsuarios())
	usuarioRoute.POST("/", controladorUsuarios.AgregarUsuarios())
	// usuarioRoute.PUT("/:id", controladorUsuarios.ActualizarUsuario())
	usuarioRoute.PATCH("/:id", controladorUsuarios.ActualizarAtribUsuario())
	usuarioRoute.DELETE("/:id", controladorUsuarios.BorrarUsuario())
	if err := router.Run(); err != nil {
		panic("no se pudo correr la aplicacion")
	}
}
