package main

import (
	"fmt"
	"goweb/cmd/server/handler"
	"goweb/docs"
	"goweb/internal/transactions"
	"goweb/pkg/store"
	"goweb/pkg/web"
	"log"
	"net/http"
	"os"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

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

func AuthMiddleware() gin.HandlerFunc {
	token := os.Getenv("TOKEN")
	if token == "" {
		log.Fatal("Token not found: Please set enviroment variable TOKEN")
	}
	return func(ctx *gin.Context) {
		tokenRequest := ctx.GetHeader("Authorization")
		if token != tokenRequest {
			ctx.AbortWithStatusJSON(web.NewResponse(http.StatusUnauthorized, nil, "Access Denied: Token Unauthorized"))
			return
		}
		ctx.Next()
	}

}

// @title MELI Bootcamp API
// @version 1.0
// @description This API Handle MELI Transactions.
// @termsOfService https://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones

// @contact.name API Support
// @contact.url https://developers.mercadolibre.com.ar/support

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
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

	docs.SwaggerInfo.Host = os.Getenv("HOST")

	router := gin.Default()
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	transactionsRoute := router.Group("/transactions")
	transactionsRoute.Use(AuthMiddleware())
	transactionsRoute.GET("/", transactions.GetAll())
	transactionsRoute.GET("/search", transactions.Search())
	transactionsRoute.GET("/:id", transactions.GetById())
	transactionsRoute.POST("/", transactions.CreateTransaction())
	transactionsRoute.PUT("/:id", transactions.Update())
	transactionsRoute.PATCH("/:id", transactions.UpdateCurrencyAndAmount())
	transactionsRoute.DELETE("/:id", transactions.Delete())
	if err := router.Run(); err != nil {
		fmt.Println(err.Error())
	}

}
