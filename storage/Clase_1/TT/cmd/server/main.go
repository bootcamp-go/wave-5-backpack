package main

import (
	"database/sql"
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
		log.Fatal(err)
	}
	if err = StorageDB.Ping(); err != nil {
		log.Fatal(err)
	}
	log.Println("database configured")

	repository := products.NewRepository(StorageDB)
	service := products.NewService(repository)
	products := handler.NewProduct(service)

	router := gin.Default()

	r := router.Group("/products")
	r.GET("/byName/:name", products.GetOneProductByName())
	r.GET("/:id", products.GetById())
	r.PATCH("/:id", products.Update())
	r.GET("/", products.GetAll())
	r.POST("/", products.Create())
	r.DELETE("/:id", products.Delete())

	if err := router.Run("localhost:8080"); err != nil {
		panic("err")
	}
}
