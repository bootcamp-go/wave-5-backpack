package main

import (
	"github.com/bootcamp-go/storage/cmd/server/handler"
	cnn "github.com/bootcamp-go/storage/db"
	"github.com/bootcamp-go/storage/internal/products"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

/* Ejercicio 1 - Refactor de Repository
Aplicar las buenas prácticas y recomendaciones para refactorizar el código de la capa repository.
1. Almacenar las queries como constantes.
2. No utilizar ”select *” en las queries.
*/
func main() {
	loadEnv()

	db := cnn.MySQLConnection()
	repo := products.NewRepository(db)
	serv := products.NewService(repo)
	p := handler.NewProduct(serv)

	r := gin.Default()
	pr := r.Group("/api/v1/products")

	pr.POST("/", p.Store())
	pr.GET("/:id", p.GetAll())
	pr.GET("/", p.GetByName())
	pr.PATCH("/:id", p.Update())

	r.Run()
}

func loadEnv() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
}
