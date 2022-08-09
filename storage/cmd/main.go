package main

import (
	"database/sql"
	"log"

	"github.com/bootcamp-go/wave-5-backpack/storage/cmd/server/handler"
	"github.com/bootcamp-go/wave-5-backpack/storage/internal/products"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dataSource := "root@tcp(localhost:3306)/storage"

	storageDB, err := sql.Open("mysql", dataSource)
	if err != nil {
		log.Fatal("ERROR")
	}

	if err = storageDB.Ping(); err != nil {
		log.Fatal("ERROR PING")
	}

	repository := products.NewRepository(storageDB)
	service := products.NewService(repository)
	product := handler.NewProduct(service)

	router := gin.Default()

	group := router.Group("products")

	group.GET("/", product.GetProductByName())
	//group.GET("/", product.GetAll())
	group.POST("/", product.Store())
	//group.PUT("/:id", product.UpdateAll())
	//group.PATCH("/:id", product.Update())
	//group.DELETE("/:id", product.Delete())

	//group.GET("/", GetFilter)

	if err := router.Run(); err != nil {
		panic(err)
	}
}
