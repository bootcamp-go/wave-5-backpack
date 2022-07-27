package main

import (
	"github.com/bootcamp-go/wave-5-backpack/goweb/cmd/server/handler"
	"github.com/bootcamp-go/wave-5-backpack/goweb/internal/products"
	"github.com/bootcamp-go/wave-5-backpack/goweb/pkg/storage"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var variableNoUsada string

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("error al intentar leer el archivo .env")
	}

	file, err := os.Open("products.json")
	defer file.Close()

	storage := storage.NewStorage("products.json")
	repo := products.NewRepository(storage)
	service := products.NewService(repo)
	p := handler.NewProduct(service)

	router := gin.Default()

	rt := router.Group("products")
	{
		rt.GET("", p.GetAll)
		rt.GET("/:id", p.GetByID)

		rt.PUT("/:id", p.Update)
		rt.PATCH("/:id", p.UpdatePrecioStock)

		rt.POST("", p.CreateProduct)

		rt.DELETE("/:id", p.Delete)
	}

	router.Run()

}
