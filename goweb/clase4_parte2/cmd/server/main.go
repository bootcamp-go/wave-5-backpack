package main
// Se importan e inyectan el repositorio, servicio y handler
import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"goweb/clase4_parte2/cmd/server/handler"
	"goweb/clase4_parte2/internal/users"
	"goweb/clase4_parte2/pkg/store"
	"goweb/clase4_parte2/pkg/web"
	"goweb/clase4_parte2/docs"
	"log"
	"os"
)

// @title GO WEB API
// @version 1.0
// @description This API Handle MELI Users.
// @termsOfService https://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones

// @contact.name API Support
// @contact.url https://developers.mercadolibre.com.ar/support

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error al cargar archivo .env")
	}

	db := store.NewStore("../../users.json")
	if err := db.Ping(); err != nil {
		log.Fatal("Error al intentar cargar archivo json")
	}
	
	r := gin.Default()
	// Agrega un endpoint mediante GET para visualizar la documentación generada.
	docs.SwaggerInfo.Host = os.Getenv("HOST")
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	repo := users.NewRepository(db)
	service := users.NewService(repo)
	u := handler.NewUser(service)

	ur := r.Group("/users")
	ur.Use(TokenAuthMiddleware())
	{
		ur.POST("/", u.Store())
		ur.GET("/", u.GetAll())
		ur.PUT("/:id", u.Update())
		ur.PATCH("/:id", u.UpdateLastNameAndAge())
		ur.DELETE("/:id", u.Delete())
	}
	r.Run()
}

func TokenAuthMiddleware() gin.HandlerFunc {
	requiredToken := os.Getenv("TOKEN")

	if requiredToken == "" {
		log.Fatal("No se encontró la variable de entorno del token")
	}

	return func(c *gin.Context) {
		token := c.GetHeader("token")

		if token == "" {
			c.AbortWithStatusJSON(401, web.NewResponse(401, nil, "No existe token en cabecera"))
			return
		}

		if token != requiredToken {
			c.AbortWithStatusJSON(401, web.NewResponse(401, nil, "El token no es válido"))
			return
		}

		c.Next()
	}
}