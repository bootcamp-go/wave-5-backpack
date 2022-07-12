package main

import (
	"desafio-go-web/cmd/server/handler"
	"desafio-go-web/docs"
	"desafio-go-web/internal/tickets"
	"desafio-go-web/pkg/store"
	"desafio-go-web/pkg/web"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title MELI Bootcamp Desafío
// @version 1.0
// @description Documentación API Desafío GoWeb
// @termsOfService https://developers.mercadolibre.com.ar/es_ar/terminos_y_condiciones
// @contact.name Cristobal Monsalve
// @contact.url https://developers.mercadolibre.com.ar/support
// @licence.name Apache 2.0
// @licence.url http://www.apache.org/licences/LICENCE-2.0.html
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("could not load .env")
	}

	db := store.NewStore("tickets.csv")
	if _, err := db.LoadTicketsFromFile(); err != nil {
		log.Fatal("error loading tickets")
	}
	repository := tickets.NewRepository(db)
	service := tickets.NewService(repository)
	tickets := handler.NewTicketHandler(service)

	r := gin.Default()

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/ping", func(c *gin.Context) { c.String(200, "pong") })

	ticket := r.Group("/ticket")
	ticket.Use(TokenAuth())
	ticket.GET("/getByCountry/:dest", tickets.GetTicketsByCountry())
	ticket.GET("/getAverage/:dest", tickets.AverageDestination())
	if err := r.Run(); err != nil {
		panic(err)
	}
}

func TokenAuth() gin.HandlerFunc {
	requiredToken := os.Getenv("TOKEN")
	if requiredToken == "" {
		log.Fatal("could not load required token for api")
	}

	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token == "" || token != requiredToken {
			c.AbortWithStatusJSON(http.StatusUnauthorized, web.NewResponse(nil, "no autorizado"))
			return
		}
		c.Next()
	}
}
