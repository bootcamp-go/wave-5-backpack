package main

import (
	"C4-TT/cmd/server/handler"
	"C4-TT/internal/usuarios"
	"C4-TT/pkg/registro"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"C4-TT/docs"

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

	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	db := registro.NewFileStore(registro.FileType, "usuarios.json")

	repository := usuarios.NewRepository(db)
	service := usuarios.NewService(repository)
	handler := handler.NewUsuario(service)

	router := gin.Default()

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	rUsuario := router.Group("usuarios")
	rUsuario.GET("/", handler.GetAll())
	rUsuario.POST("/", handler.Registrar())
	rUsuario.PUT("/:id", handler.Modificar())
	rUsuario.DELETE("/:id", handler.Eliminar())
	rUsuario.PATCH("/:id", handler.ModificarAE())

	router.Run()
}
