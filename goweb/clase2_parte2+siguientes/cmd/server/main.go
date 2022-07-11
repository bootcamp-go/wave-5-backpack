package main

import (
	"log"

	"github.com/bootcamp-go/wave-5-backpack/tree/olivera_sebastian/goweb/clase2_parte2+siguientes/cmd/server/handler"
	"github.com/bootcamp-go/wave-5-backpack/tree/olivera_sebastian/goweb/clase2_parte2+siguientes/internal/transactions"
	"github.com/bootcamp-go/wave-5-backpack/tree/olivera_sebastian/goweb/clase2_parte2+siguientes/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	loadEnv()

	//db := store.NewFileStore(store.FileType, "transactions.json")
	db := store.NewStore("products.json")
	if err := db.Ping(); err != nil {
		log.Fatal("error al intentar cargar archivo")
	}
	repository := transactions.NewRepository(db)
	service := transactions.NewService(repository)
	handler := handler.NewTransaction(service)

	router := gin.Default()
	rTransaction := router.Group("transactions")
	rTransaction.GET("/", handler.GetAll())
	rTransaction.POST("/", handler.Store())
	rTransaction.PUT("/:id", handler.Update())
	rTransaction.DELETE("/:id", handler.Delete())
	rTransaction.PATCH("/:id", handler.UpdateReceptorYMonto())

	router.Run(":8080")
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}
