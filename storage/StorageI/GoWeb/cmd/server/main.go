package main

import (
	"GoWeb/cmd/server/handler"
	"GoWeb/internals/transactions"
	"database/sql"
	"log"
	"os"

	"GoWeb/docs"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//@title MELI Bootcamp API Juan David Serna
//@version 1.0
//@description This API Handle MELI Transactions
//@termsOfService https://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones

//@contact.name API Support
//@contact.url https://developers.mercadolibre.com.ar/support
var StorageDB *sql.DB

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("error al intentar cargar el archivo .env")
	}

	//db := store.NewStore("transactions.json")
	dataSource := "root@tcp(localhost:3306)/storage"
	StorageDB, err := sql.Open("mysql", dataSource)

	if err != nil {
		panic(err)
	}
	if err = StorageDB.Ping(); err != nil {
		panic(err)
	}

	repo := transactions.NewRepository(StorageDB)
	service := transactions.NewService(repo)
	tran := handler.NewTransaction(service)

	router := gin.Default()
	docs.SwaggerInfo.Host = os.Getenv("HOST")
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	tr := router.Group("/transacciones")
	tr.Use(tran.TokenAuthMiddleware())
	tr.POST("/", tran.Store())
	tr.GET("/", tran.GetAll())
	tr.PUT("/:id", tran.Update())
	tr.DELETE("/:id", tran.Delete())
	tr.PATCH("/:id", tran.UpdateCode())
	tr.GET("/name/:name", tran.GetByName())
	router.Run()
}
