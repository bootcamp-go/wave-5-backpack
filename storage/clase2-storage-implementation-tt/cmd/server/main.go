/*---------------------------------------------------------*

     Assignment:	C2 - TT | Practica #1
         Author:	Israel Fabela
	   Language:	go1.18.3 darwin/arm64
		  Topic:	Storage Implementation


	© Mercado Libre - IT Bootcamp 2022

----------------------------------------------------------*/

package main

import (
	"clase2-storage-implementation-tt/cmd/server/handler"
	"clase2-storage-implementation-tt/internal/transactions"
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
)

var (
	StorageDB *sql.DB
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error al intentar cargar archivo .env")
	}

	dataSource := "root@tcp(localhost:3306)/storage"
	// Open inicia un pool de conexiones. Sólo abrir una vez
	StorageDB, err = sql.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}

	if err := StorageDB.Ping(); err != nil {
		panic(err)
	}
	log.Println("database Configured")

	repo := transactions.NewRepository(StorageDB)
	service := transactions.NewService(repo)
	t := handler.NewTransaction(service)

	r := gin.Default()

	pr := r.Group("/transactions")
	{
		pr.GET("/", t.GetAll())
		pr.GET("/:id", t.GetOne())
		pr.GET("/code/:code", t.GetByName())
		pr.PUT("/:id", t.Update())
		pr.POST("/", t.Ecommerce())
		pr.PATCH("/:id", t.UpdateOne())
		pr.DELETE("/:id", t.Delete())
	}

	err = r.Run()
	if err != nil {
		panic(err)
	}
}
