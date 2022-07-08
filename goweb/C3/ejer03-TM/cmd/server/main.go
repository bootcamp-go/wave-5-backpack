package main

import (
	"ejer02-TT/cmd/server/handler"
	"ejer02-TT/internal/transactions"
	"ejer02-TT/pkg/store"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("error al cargar archivo .env")
	}

	db := store.NewStore("transacciones.json")

	repository := transactions.NewRepository(db)
	service := transactions.NewService(repository)
	transactions := handler.NewTransaction(service)

	router := gin.Default()
	tr := router.Group("transactions")
	tr.GET("/", transactions.GetAll())
	tr.POST("/", transactions.Store())
	tr.PUT("/:id", transactions.Update())
	tr.PATCH("/:id", transactions.UpdateCodeAndAmount())
	tr.DELETE("/:id", transactions.Delete())
	router.Run()

}
