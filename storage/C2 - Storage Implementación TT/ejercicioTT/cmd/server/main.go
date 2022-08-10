package main

import (
	"database/sql"
	"ejercicioTT/cmd/server/handler"
	"ejercicioTT/docs"
	users "ejercicioTT/internal/users"
	"ejercicioTT/pkg/web"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var (
	StorageDB *sql.DB
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
		log.Fatal("error al intentar cargar archivo .env")
	}

	dataSource := "root@tcp(localhost:3306)/storage?parseTime=true"

	// Open inicia un pool de conexiones. Sólo abrir una vez

	StorageDB, err = sql.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	if err = StorageDB.Ping(); err != nil {
		panic(err)
	}
	log.Println("database Configured")

	repo := users.NewRepository(StorageDB)
	serv := users.NewService(repo)
	user := handler.NewUser(serv)

	r := gin.Default()
	us := r.Group("usuarios")

	// Swagger Documentation - Endpoint
	docs.SwaggerInfo.Host = os.Getenv("HOST")
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	us.Use(TokenAuthMiddleware())
	us.POST("/", user.Store())
	us.PUT("/", user.Update())
	us.GET("/", user.GetAll())
	us.GET("/:id", user.GetOne())
	us.GET("/userware/:id", user.GetFullData())
	us.GET("/byName/:nombre", user.GetByName())
	us.DELETE("/:id", user.Delete())
	//us.PATCH("/:id", user.UpdateLastAge())

	r.Run()
}

func TokenAuthMiddleware() gin.HandlerFunc {
	requiredToken := os.Getenv("TOKEN")

	if requiredToken == "" {
		log.Fatal("no se encontró el token en variable de entorno")
	}

	return func(c *gin.Context) {
		token := c.GetHeader("token")

		if token == "" {
			c.AbortWithStatusJSON(401, web.NewResponse(401, nil, "falta token en cabecera"))
			return
		}

		if token != requiredToken {
			c.AbortWithStatusJSON(401, web.NewResponse(401, nil, "token incorrecto"))
			return
		}

		c.Next()
	}

}
