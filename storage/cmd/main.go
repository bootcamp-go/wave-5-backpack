package main

import (
	"log"

	"github.com/bootcamp-go/wave-5-backpack/storage/cmd/server/handler"
	"github.com/bootcamp-go/wave-5-backpack/storage/internal/products"
	"github.com/bootcamp-go/wave-5-backpack/storage/pkg/store"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("error al intentar cargar archivo .env")
	}

	repository := products.NewRepository(store.DBConnection())
	service := products.NewService(repository)
	product := handler.NewProduct(service)

	router := gin.Default()

	group := router.Group("products")

	group.GET("/", product.GetProductByName())
	group.GET("/p&w", product.GetProductAndWareHouse())
	group.GET("/all", product.GetAll())
	group.POST("/", product.Store())
	group.PUT("/:id", product.UpdateAll())
	//group.PATCH("/:id", product.Update())
	//group.DELETE("/:id", product.Delete())

	//group.GET("/", GetFilter)

	if err := router.Run(); err != nil {
		panic(err)
	}
}
