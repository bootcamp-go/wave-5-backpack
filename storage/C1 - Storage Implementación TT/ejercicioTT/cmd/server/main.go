package db

import (
	"database/sql"
	"ejercicioTT/cmd/server/handler"
	"ejercicioTT/docs"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
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
	dataSource := "root@tcp(localhost:3306)/storage"

	// Open inicia un pool de conexiones. Sólo abrir una vez
	var err error
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
	us.GET("/", user.GetAll())
	us.POST("/", user.Store())
	us.PUT("/:id", user.Update())
	us.PATCH("/:id", user.UpdateLastAge())
	us.DELETE("/:id", user.Delete())

	r.Run()
}
