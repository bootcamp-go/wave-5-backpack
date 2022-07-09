package main

import (
	"GoWeb/cmd/server/handler"
	"GoWeb/internals/transactions"
	"GoWeb/pkg/store"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("error al intentar cargar el archivo .env")
	}

	db := store.NewStore("transactions.json")
	if err := db.Ping(); err != nil {
		log.Fatal("error al intentar cargar el archivo")
	}

	repo := transactions.NewRepository(db)
	service := transactions.NewService(repo)
	tran := handler.NewTransaction(service)

	router := gin.Default()
	tr := router.Group("/transacciones")
	tr.POST("/", tran.Store())
	tr.GET("/", tran.GetAll())
	tr.PUT("/:id", tran.Update())
	tr.DELETE("/:id", tran.Delete())
	tr.PATCH("/:id", tran.UpdateCode())
	router.Run()
}
