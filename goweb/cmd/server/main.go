package main

import (
	"goweb/cmd/server/handler"
	"goweb/docs"
	"goweb/internal/transactions"
	"goweb/pkg/store"
	"goweb/pkg/web"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func TokenAuthMiddleWare() gin.HandlerFunc {
	requiredToken := os.Getenv("TOKEN")

	if requiredToken == "" {
		log.Fatal("Token don't found in enviroment variables")
	}

	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token == "" {
			ctx.AbortWithStatusJSON(401, web.NewResponse(401, nil, "token don't found in header"))
			return
		}
		if token != requiredToken {
			ctx.AbortWithStatusJSON(401, web.NewResponse(401, nil, "incorrect token"))
			return
		}
		ctx.Next()
	}
}

// @title MELI Bootcamp API
// @version 1.0
// @description This API Handle MELI Transactions
// @tremsofService https://developers.mercadolibre.com.co/es_co/terminos-y-condiciones

// @contact.name API Support
// @contact.url https://developers.mercadolibre.com.co/support

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE/2.0.html
func main() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("error al intentar cargar archivo .env")
	}
	db := store.NewStore("transactions.json")
	repo := transactions.NewRepository(db)
	service := transactions.NewService(repo)
	t := handler.NewTransaction(service)

	router := gin.Default()

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/transactions", TokenAuthMiddleWare(), t.GetAll)

	gt := router.Group("/transaction")

	gt.Use(TokenAuthMiddleWare())
	gt.POST("/", t.Create)
	gt.GET("/:id", t.GetOne)
	gt.PUT("/:id", t.Update)
	gt.DELETE("/:id", t.Delete)
	gt.PATCH("/:id", t.Update2)
	router.Run()
}
