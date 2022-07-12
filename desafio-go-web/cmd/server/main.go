package main

import (
	"desafio-go-web/cmd/server/handler"
	"desafio-go-web/docs"
	"desafio-go-web/internal/tickets"
	"desafio-go-web/pkg/store"
	"os"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title MELI Bootcamp API
// @version 1.0
// @description This API Handle Fly Tickets.
// @termsOfService https://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones

// @contact.name API Support
// @contact.url https://developers.mercadolibre.com.ar/support

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {

	// Cargo csv.
	store := store.NewStore("./tickets.csv")
	ticketsRecords, err := store.ReadTickets()
	repo := tickets.NewRepository(ticketsRecords)
	service := tickets.NewService(repo)
	handler := handler.NewService(service)

	if err != nil {
		panic("Couldn't load tickets")
	}

	docs.SwaggerInfo.Host = os.Getenv("HOST")

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) { c.String(200, "pong") })
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	ticketRoutes := r.Group("/ticket")
	{
		ticketRoutes.GET("/getByCountry/:dest", handler.GetTicketsByCountry())
		ticketRoutes.GET("/getPercentage/:dest", handler.PercentageByDestination())
	}
	if err := r.Run(); err != nil {
		panic(err)
	}

}
