package main

import (
	"desafio-go-web/cmd/server/handler"
	"desafio-go-web/internal/tickets"
	"desafio-go-web/pkg/store"

	"github.com/gin-gonic/gin"
)

func main() {

	file := store.NewStore("./data/tickets.csv")
	storeTickets, _ := file.ReadCSV()

	repository := tickets.NewRepository(storeTickets)
	service := tickets.NewService(repository)
	ticket := handler.NewTicket(service)

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) { c.String(200, "pong") })
	// Rutas a desarollar:
	routerTicket := r.Group("/ticket")

	// GET - “/”
	routerTicket.GET("/", ticket.GetAllTickets())

	// GET - “/ticket/getByCountry/:dest”
	routerTicket.GET("/getByCountry/:dest", ticket.GetTicketsByCountry())

	// GET - “/ticket/getAverage/:dest”
	routerTicket.GET("/getAverage/:dest", ticket.AverageDestination())

	if err := r.Run(); err != nil {
		panic(err)
	}

}
