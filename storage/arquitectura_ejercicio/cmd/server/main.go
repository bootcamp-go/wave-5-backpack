package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/anesquivel/wave-5-backpack/storage/arquitectura_ejercicio/cmd/server/handler"
	"github.com/anesquivel/wave-5-backpack/storage/arquitectura_ejercicio/db"
	db2 "github.com/anesquivel/wave-5-backpack/storage/arquitectura_ejercicio/db"
	"github.com/anesquivel/wave-5-backpack/storage/arquitectura_ejercicio/docs"
	"github.com/anesquivel/wave-5-backpack/storage/arquitectura_ejercicio/internal/usuarios"
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
	gin.SetMode(gin.ReleaseMode)
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error al intentar cargar archivo .env")
	}

	//db := store.NewStore("usuarios.json")
	db2.Init()
	db := db.StorageDB

	/*if errR := db.Ping(); errR != nil {
		log.Fatal("error al intentar cargar archivo")
	}*/

	repo := usuarios.NewRepository(db)
	service := usuarios.NewService(repo)
	u := handler.NewUsuario(service)

	r := gin.Default()

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	users := r.Group("/usuarios")
	users.POST("/", u.Store())
	users.GET("/", u.GetAll())
	users.PUT("/:id", u.Update())
	users.PATCH("/:id", u.PatchLastNameAge())
	users.DELETE("/:id", u.Delete())

	errServer := r.Run()

	if errServer != nil {
		fmt.Println("error running server")
	}
}
