package main

import (
	"database/sql"
	"fmt"
	"goweb/cmd/server/handler"
	"goweb/internal/products"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error al intentar cargar archivo .env")
	}

	dataSource := "root:@tcp(localhost:3306)/storage"
	StorageDB, err := sql.Open("mysql", dataSource)
	if err != nil {
		fmt.Println("estoy aqui 0")
		log.Fatal(err)
	}
	if err = StorageDB.Ping(); err != nil {
		log.Fatal(err)
	}
	log.Println("database configured")

	repository := products.NewRepository(StorageDB)
	products := handler.NewProduct(repository)

	router := gin.Default()

	r := router.Group("/products")
	r.GET("/:name", products.GetOneProductByName())
	r.POST("/", products.Create())

	if err := router.Run("localhost:8080"); err != nil {
		panic("err")
	}
}
