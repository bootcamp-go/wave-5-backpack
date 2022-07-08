package main

import (
	"goweb/cmd/server/handler"
	"goweb/internal/transactions"
	"goweb/pkg/store"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("error al intentar cargar archivo .env")
	}
	db := store.NewStore("transactions.json")
	repo := transactions.NewRepository(db)
	service := transactions.NewService(repo)
	t := handler.NewTransaction(service)

	router := gin.Default()
	router.GET("/transactions", t.GetAll)

	gt := router.Group("/transaction")
	gt.POST("/", t.Create)
	gt.GET("/:id", t.GetOne)
	gt.PUT("/:id", t.Update)
	gt.DELETE("/:id", t.Delete)
	gt.PATCH("/:id", t.Update2)
	router.Run()
}
