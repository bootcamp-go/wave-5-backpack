package main

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/nictes1/storage-implementation/cmd/server/handler"
	"github.com/nictes1/storage-implementation/internal/products"
)

func main() {

	dataSource := "root@tcp(localhost:3306)/storage"
	// Open inicia un pool de conexiones. Sólo abrir una vez
	var err error
	StorageDB, err := sql.Open("mysql", dataSource)
	if err != nil {
		log.Fatal(err)
	}
	if err = StorageDB.Ping(); err != nil {
		log.Fatal(err)
	}
	log.Println("database configured")

	repository := products.NewRepo(StorageDB)
	service := products.NewService(repository)
	p := handler.NewProduct(service)

	r := gin.Default()
	rg := r.Group(("/products"))
	rg.GET("/", p.GetAll())
	rg.GET("/:id", p.Get())
	rg.POST("/", p.Store())
	rg.PUT("/:id", p.Update())

	r.Run() //:8080

}
