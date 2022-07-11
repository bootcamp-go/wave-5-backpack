package main

import (
	"log"

	"github.com/bootcamp-go/wave-5-backpack/goweb/cmd/handler"
	"github.com/bootcamp-go/wave-5-backpack/goweb/internal/products"
	"github.com/bootcamp-go/wave-5-backpack/goweb/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("error al intentar cargar archivo .env")
	}

	db := store.NewStore("products.json")
	if err := db.Ping(); err != nil {
		log.Fatal("error al intentar cargar el archivo .json")
	}

	repository := products.NewRepository(db)
	service := products.NewService(repository)
	product := handler.NewProduct(service)

	router := gin.Default()

	group := router.Group("products")

	group.GET("/:id", product.GetProduct())
	group.GET("/", product.GetAll())
	group.POST("/", product.Store())
	group.PUT("/:id", product.UpdateAll())
	group.PATCH("/:id", product.Update())
	group.DELETE("/:id", product.Delete())

	//group.GET("/", GetFilter)

	router.Run()

}
