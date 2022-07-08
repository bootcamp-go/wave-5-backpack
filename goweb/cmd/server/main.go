package main

import (
	"fmt"
	"goweb/cmd/server/handler"
	"goweb/internal/transactions"
	"goweb/pkg/store"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

func setupValidation() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
	}
}

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Println(err.Error())
		return
	}
	db := store.NewStore("transactions.json")
	if err := db.Ping(); err != nil {
		fmt.Println(err.Error())
		return
	}

	repository := transactions.NewRepository(db)
	service := transactions.NewService(repository)
	transactions := handler.NewTransaction(service)
	setupValidation()

	router := gin.Default()

	transactionsRoute := router.Group("/transactions")
	transactionsRoute.GET("/", transactions.GetAll())
	transactionsRoute.GET("/search", transactions.Search())
	transactionsRoute.GET("/:id", transactions.GetById())
	transactionsRoute.POST("/", transactions.CreateTransaction())
	transactionsRoute.PUT("/:id", transactions.Update())
	transactionsRoute.PATCH("/:id", transactions.UpdateCurrencyAndAmount())
	transactionsRoute.DELETE("/:id", transactions.Delete())
	router.Run()
}
