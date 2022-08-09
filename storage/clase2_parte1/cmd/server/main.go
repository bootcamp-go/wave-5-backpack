package main

import (
	"database/sql"
	"log"

	"clase2_parte1/cmd/server/handler"
	"clase2_parte1/internal/products"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	dataSource := "root@tcp(localhost:3306)/storage"
	var err error
	StorageDB, err := sql.Open("mysql", dataSource)
	if err != nil {
		log.Fatal(err)
	}
	if err = StorageDB.Ping(); err != nil {
		log.Fatal(err)
	}

	repository := products.NewRepo(StorageDB)
	service := products.NewService(repository)
	p := handler.NewProduct(service)

	r := gin.Default()
	rg := r.Group(("/products"))
	rg.GET("/", p.GetAll())
	rg.GET("/:id", p.Get())
	rg.GET("/getFullData/:id", p.GetFullData())
	rg.GET("/withContext/:id", p.GetOneWithcontext())
	rg.POST("/", p.Store())
	rg.DELETE("/:id", p.Delete())

	r.Run()

}
