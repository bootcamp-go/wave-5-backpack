package main

import (
	"WebServer/internal/transactions"
	"WebServer/pkg/store"
	"WebServer/server/handler"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	// it will be visible for this scope.
	db := store.NewStore("Transactions.json")
	if err := db.Ping(); err != nil {
		log.Fatal("Cannot read file")
	}

	repo := transactions.NewRepository(db)
	service := transactions.NewService(repo)
	t := handler.NewTransaction(service)

	r := gin.Default()
	tr := r.Group("/transactions")
	tr.POST("/", t.Create())
	tr.GET("/", t.GetAll())
	tr.PUT("/:id", t.Update())
	tr.PATCH("/:id", t.UpdatePartial())
	tr.DELETE("/:id", t.Delete())
	r.Run()
}
