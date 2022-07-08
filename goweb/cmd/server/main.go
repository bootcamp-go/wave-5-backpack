package main

import (
	"goweb/cmd/server/handler"
	"goweb/internal/transactions"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
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
	repository := transactions.NewRepository()
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
