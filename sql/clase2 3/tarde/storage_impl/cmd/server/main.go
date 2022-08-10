package main

import (
	"github.com/bootcamp-go/storage/cmd/server/handler"
	cnn "github.com/bootcamp-go/storage/db"
	"github.com/bootcamp-go/storage/internal/products"
	"github.com/bootcamp-go/storage/internal/users"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

const (
	tableUsers = "Users"
)

func main() {
	loadEnv()

	db := cnn.MySQLConnection()
	repo := products.NewRepository(db)
	serv := products.NewService(repo)
	p := handler.NewProduct(serv)
	r := gin.Default()
	// Products
	pr := r.Group("/api/v1/products")
	pr.POST("/", p.Store())
	pr.GET("/:id", p.GetAll())
	pr.GET("/", p.GetByName())
	pr.PATCH("/:id", p.Update())

	dynamodb := cnn.InitDynamo()
	repoU := users.NewDynamoRepository(dynamodb, tableUsers)
	servU := users.NewService(repoU)
	u := handler.NewUser(servU)

	// Users (Dynamo Implement)
	ur := r.Group("/api/v1/users")
	ur.POST("/", u.Store())
	ur.GET("/:id", u.GetOne())
	ur.DELETE("/:id", u.Delete())
	ur.PATCH("/:id", u.Update())
	r.Run()
}

func loadEnv() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
}
