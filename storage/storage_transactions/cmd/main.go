package main

import (
	"bootcamp/wave-5-backpack/storage/cmd/server/handler"
	cnn "bootcamp/wave-5-backpack/storage/db"
	"bootcamp/wave-5-backpack/storage/internal/products"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	loadEnv()

	db := cnn.MySQLConnection()
	repo := products.NewRepo(db)
	serv := products.NewService(repo)
	p := handler.NewProduct(serv)

	r := gin.Default()
	rg := r.Group(("/products"))
	rg.GET("/", p.GetAll())
	rg.GET("/:id", p.Get())
	rg.GET("/getFullData/:id", p.GetFullData())
	rg.GET("/withContext/:id", p.GetOneWithContext())
	rg.POST("/", p.Store())
	rg.DELETE("/:id", p.Delete())

	r.Run()
}

func loadEnv() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
}
