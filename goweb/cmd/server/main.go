package main

import (
	"goweb/cmd/server/handler"
	"goweb/internal/products"
	"goweb/pkg/store"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Cargar las variables de entorno
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("error al intentar cargar archivo .env")
	}

	db := store.NewStore(os.Getenv("DB_PATH"))
	if err := db.Ping(); err != nil {
		log.Fatal("error al conectar con la base de datos")
	}

	//repository := products.NewRepositoryMemoria()
	//repository := products.NewRepositoryJsonDB()
	repository := products.NewRepositoryJsonCorrDB(db)
	service := products.NewService(repository)
	p := handler.NewProduct(service)

	router := gin.Default()
	productos := router.Group("/productos")
	{
		// Clase 1 Ejercicio 2 Parte 1
		productos.GET("/", p.GetAll())
		// Clase 1 Ejercicio 2 Parte 2
		productos.GET("/:id", p.GetById())
		// Clase 2 Ejercicio 1 Parte 1
		productos.POST("/", p.Store())
		// Clase 3 Ejercicio 1 Parte 1
		productos.PUT("/:id", p.Update())
		// Clase 3 Ejercicio 1 Parte 1
		productos.DELETE("/:id", p.Delete())
		// Clase 3 Ejercicio 1 Parte 1
		productos.PATCH("/:id", p.UpdateNombreYPrecio())
	}

	// Clase 1 Ejercicio 1 Parte 1
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "Hola " + "Juan Pablo Ortiz"})
	})

	router.Run(":8080")
}
