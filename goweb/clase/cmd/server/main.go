package main

import (
	"github.com/bootcamp-go/go-web/cmd/server/handler"
	"github.com/bootcamp-go/go-web/internal/transactions"
	"github.com/bootcamp-go/go-web/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

/*Ejercicio 2 - Guardar información

Se debe implementar la funcionalidad para guardar la información de la petición en un archivo json,
para lograrlo se deben realizar los siguientes pasos:
	1. En lugar de guardar los valores de nuestra entidad en memoria, se debe crear un archivo;
	los valores que se vayan agregando se guardan en él.
*/

/*Ejercicio 3 - Leer información

Se debe implementar la funcionalidad para leer la información requerida en la petición del archivo
json generado al momento de guardar, para lograrlo se deben realizar los siguientes pasos:
	1. En lugar de leer los valores de nuestra entidad en memoria, se debe obtener del archivo
	generado en el punto anterior.
*/

func main() {
	loadEnv()

	db := store.NewFileStore(store.FileType, "./transactions.json")
	r := transactions.NewRepository(db)
	s := transactions.NewService(r)
	h := handler.NewTransaction(s)

	router := gin.Default()
	rTransaction := router.Group("transactions")
	rTransaction.GET("/", h.GetAll())
	rTransaction.POST("/", h.Store())
	rTransaction.PUT("/:id", h.Update())
	rTransaction.DELETE("/:id", h.Delete())
	rTransaction.PATCH("/:id", h.UpdateReceptorYMonto())

	router.Run(":8000")
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}
