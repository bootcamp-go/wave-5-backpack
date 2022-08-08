package main

import (
	"log"
	"os"
	"products_project/cmd/server/handler"
	"products_project/docs"
	"products_project/internal/products"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var (
	StorageDB *sql.DB
)

// @title MELI Bootcamp API | Jessica Escobar
// @version 1.0
// @description This API implements the CRUD method for MELI Products
// @termsOfService https://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones

// @contact.name API Support
// @contact.url https://developers.mercadolibre.com.ar/support

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error al intentar cargar el archivo .env")
	}

	// db := store.NewStore("products.json")
	dataSource := "root@tcp(localhost:3306)/storage"
	StorageDB, err = sql.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	if err = StorageDB.Ping(); err != nil {
		panic(err)
	}
	repo := products.NewRepository(StorageDB)
	service := products.NewService(repo)
	p := handler.NewProduct(service)

	router := gin.Default()
	pr := router.Group("/products")

	// Swagger Documentation - Endpoint
	docs.SwaggerInfo.Host = os.Getenv("HOST")
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// <--------------------------------------------------------------->
	pr.Use(p.TokenAuthMiddleware())
	pr.POST("/", p.Store())
	pr.GET("/", p.GetAll())
	pr.GET("/:id", p.GetById())
	pr.GET("/nombre/:nombre", p.GetByName())
	pr.PUT("/:id", p.Update())
	pr.PATCH("/:id", p.UpdateFields())
	pr.DELETE("/:id", p.Delete())
	err = router.Run()
	if err != nil {
		panic(err)
	}
}
